// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildlet

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/build/buildenv"
	"golang.org/x/build/dashboard"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

// GCEGate optionally specifies a function to run before any GCE API call.
// It's intended to be used to bound QPS rate to GCE.
var GCEGate func()

func apiGate() {
	if GCEGate != nil {
		GCEGate()
	}
}

// StartNewVM boots a new VM on GCE and returns a buildlet client
// configured to speak to it.
func StartNewVM(creds *google.Credentials, buildEnv *buildenv.Environment, instName, hostType string, opts VMOpts) (Client, error) {
	ctx := context.TODO()
	computeService, _ := compute.New(oauth2.NewClient(ctx, creds.TokenSource))

	if opts.Description == "" {
		opts.Description = fmt.Sprintf("Go Builder for %s", hostType)
	}
	if opts.ProjectID == "" {
		opts.ProjectID = buildEnv.ProjectName
	}
	if opts.Zone == "" {
		opts.Zone = buildEnv.RandomVMZone()
	}
	zone := opts.Zone
	if opts.DeleteIn == 0 {
		opts.DeleteIn = 30 * time.Minute
	}

	hconf, ok := dashboard.Hosts[hostType]
	if !ok {
		return nil, fmt.Errorf("invalid host type %q", hostType)
	}
	if !hconf.IsVM() && !hconf.IsContainer() {
		return nil, fmt.Errorf("host %q is type %q; want either a VM or container type", hostType, hconf.PoolName())
	}

	projectID := opts.ProjectID
	if projectID == "" {
		return nil, errors.New("buildlet: missing required ProjectID option")
	}

	prefix := "https://www.googleapis.com/compute/v1/projects/" + projectID
	machType := prefix + "/zones/" + zone + "/machineTypes/" + hconf.MachineType()
	diskType := "https://www.googleapis.com/compute/v1/projects/" + projectID + "/zones/" + zone + "/diskTypes/pd-ssd"
	if hconf.RegularDisk {
		diskType = "" // a spinning disk
	}

	srcImage := "https://www.googleapis.com/compute/v1/projects/" + projectID + "/global/images/" + hconf.VMImage
	minCPU := hconf.MinCPUPlatform
	if hconf.IsContainer() {
		if hconf.NestedVirt {
			minCPU = "Intel Cascade Lake" // n2 vms (which support NestedVirtualization) are either Ice Lake or Cascade Lake.
		}
		if vm := hconf.ContainerVMImage(); vm != "" {
			srcImage = "https://www.googleapis.com/compute/v1/projects/" + projectID + "/global/images/" + vm
		} else {
			var err error
			srcImage, err = cosImage(ctx, computeService)
			if err != nil {
				return nil, fmt.Errorf("error find Container-Optimized OS image: %v", err)
			}
		}
	}

	instance := &compute.Instance{
		Name:           instName,
		Description:    opts.Description,
		MachineType:    machType,
		MinCpuPlatform: minCPU,
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    instName,
					SourceImage: srcImage,
					DiskType:    diskType,
				},
			},
		},
		Tags: &compute.Tags{
			// Warning: do NOT list "http-server" or "allow-ssh" (our
			// project's custom tag to allow ssh access) here; the
			// buildlet provides full remote code execution.
			// The https-server is authenticated, though.
			Items: []string{"https-server"},
		},
		Metadata: &compute.Metadata{},
		NetworkInterfaces: []*compute.NetworkInterface{{
			Network: prefix + "/global/networks/default-vpc",
		}},

		// Prior to git rev 1b1e086fd, we used preemptible
		// instances, as we were helping test the feature. It was
		// removed after git rev a23395d because we hadn't been
		// using it for some time. Our VMs are so short-lived that
		// the feature doesn't really help anyway. But if we ever
		// find we want it again, this comment is here to point to
		// code that might be useful to partially resurrect.
		Scheduling: &compute.Scheduling{Preemptible: false},
	}

	// Container builders use the COS image, which defaults to logging to Cloud Logging.
	// Permission is granted to this service account.
	if hconf.IsContainer() && buildEnv.COSServiceAccount != "" {
		instance.ServiceAccounts = []*compute.ServiceAccount{
			{
				Email:  buildEnv.COSServiceAccount,
				Scopes: []string{compute.CloudPlatformScope},
			},
		}
	}

	addMeta := func(key, value string) {
		instance.Metadata.Items = append(instance.Metadata.Items, &compute.MetadataItems{
			Key:   key,
			Value: &value,
		})
	}
	// The buildlet-binary-url is the URL of the buildlet binary
	// which the VMs are configured to download at boot and run.
	// This lets us/ update the buildlet more easily than
	// rebuilding the whole VM image.
	addMeta("buildlet-binary-url", hconf.BuildletBinaryURL(buildenv.ByProjectID(opts.ProjectID)))
	addMeta("buildlet-host-type", hostType)
	if !opts.TLS.IsZero() {
		addMeta("tls-cert", opts.TLS.CertPEM)
		addMeta("tls-key", opts.TLS.KeyPEM)
		addMeta("password", opts.TLS.Password())
	}
	if hconf.IsContainer() {
		addMeta("gce-container-declaration", fmt.Sprintf(`spec:
  containers:
    - name: buildlet
      image: 'gcr.io/%s/%s'
      volumeMounts:
        - name: tmpfs-0
          mountPath: /workdir
      securityContext:
        privileged: true
      stdin: false
      tty: false
  restartPolicy: Always
  volumes:
    - name: tmpfs-0
      emptyDir:
        medium: Memory
`, opts.ProjectID, hconf.ContainerImage))
	}

	if opts.DeleteIn > 0 {
		// In case the VM gets away from us (generally: if the
		// coordinator dies while a build is running), then we
		// set this attribute of when it should be killed so
		// we can kill it later when the coordinator is
		// restarted. The cleanUpOldVMs goroutine loop handles
		// that killing.
		addMeta("delete-at", fmt.Sprint(time.Now().Add(opts.DeleteIn).Unix()))
	}

	for k, v := range opts.Meta {
		addMeta(k, v)
	}

	apiGate()
	op, err := computeService.Instances.Insert(projectID, zone, instance).Do()
	if err != nil {
		return nil, fmt.Errorf("Failed to create instance: %v", err)
	}
	condRun(opts.OnInstanceRequested)
	createOp := op.Name

	// Wait for instance create operation to succeed.
OpLoop:
	for {
		time.Sleep(2 * time.Second)
		apiGate()
		op, err := computeService.ZoneOperations.Get(projectID, zone, createOp).Do()
		if err != nil {
			return nil, fmt.Errorf("Failed to get op %s: %v", createOp, err)
		}
		switch op.Status {
		case "PENDING", "RUNNING":
			continue
		case "DONE":
			if op.Error != nil {
				for _, operr := range op.Error.Errors {
					log.Printf("failed to create instance %s in zone %s: %v", instName, zone, operr.Code)
					// TODO: catch Code=="QUOTA_EXCEEDED" and "Message" and return
					// a known error value/type.
					return nil, fmt.Errorf("Error creating instance: %+v", operr)
				}
				return nil, errors.New("Failed to start.")
			}
			break OpLoop
		default:
			return nil, fmt.Errorf("Unknown create status %q: %+v", op.Status, op)
		}
	}
	condRun(opts.OnInstanceCreated)

	apiGate()
	inst, err := computeService.Instances.Get(projectID, zone, instName).Do()
	if err != nil {
		return nil, fmt.Errorf("Error getting instance %s details after creation: %v", instName, err)
	}

	// Finds its internal and/or external IP addresses.
	intIP, extIP := instanceIPs(inst)

	// Wait for it to boot and its buildlet to come up.
	var buildletURL string
	var ipPort string
	if !opts.TLS.IsZero() {
		if extIP == "" {
			return nil, errors.New("didn't find its external IP address")
		}
		buildletURL = "https://" + extIP
		ipPort = extIP + ":443"
	} else {
		if intIP == "" {
			return nil, errors.New("didn't find its internal IP address")
		}
		buildletURL = "http://" + intIP
		ipPort = intIP + ":80"
	}
	if opts.OnGotInstanceInfo != nil {
		opts.OnGotInstanceInfo(inst)
	}
	var closeFuncs []func()
	if opts.UseIAPTunnel {
		localPort, closeFunc, err := createIAPTunnel(ctx, inst)
		if err != nil {
			return nil, fmt.Errorf("creating IAP tunnel: %v", err)
		}
		buildletURL = "http://localhost:" + localPort
		ipPort = "127.0.0.1:" + localPort
		closeFuncs = append(closeFuncs, closeFunc)
	}
	client, err := buildletClient(ctx, buildletURL, ipPort, &opts)
	if err != nil {
		return nil, err
	}
	for _, cf := range closeFuncs {
		client.AddCloseFunc(cf)
	}
	return client, nil
}

func createIAPTunnel(ctx context.Context, inst *compute.Instance) (string, func(), error) {
	// Allocate a local listening port.
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", nil, err
	}
	localAddr := ln.Addr().(*net.TCPAddr)
	ln.Close()
	// Start the gcloud command. For some reason, when gcloud is run with a
	// pipe for stdout, it doesn't log the success message, so we can only
	// check for success empirically.
	tunnelCmd := exec.CommandContext(ctx,
		"gcloud", "compute", "start-iap-tunnel", "--iap-tunnel-disable-connection-check",
		"--zone", inst.Zone, inst.Name, "80", "--local-host-port", localAddr.String())
	tunnelCmd.Stderr = os.Stderr
	tunnelCmd.Stdout = os.Stdout
	if err := tunnelCmd.Start(); err != nil {
		return "", nil, err
	}
	// Start the process. Either it's going to fail to start after a bit, or
	// it'll start listening on its port. Because we told it not to check the
	// connection above, the connections won't be functional, but we can dial.
	errc := make(chan error, 1)
	go func() { errc <- tunnelCmd.Wait() }()
	for start := time.Now(); time.Since(start) < 60*time.Second; time.Sleep(5 * time.Second) {
		// Check if the server crashed.
		select {
		case err := <-errc:
			return "", nil, err
		default:
		}
		// Check if it's healthy.
		conn, err := net.DialTCP("tcp", nil, localAddr)
		if err == nil {
			conn.Close()
			kill := func() { tunnelCmd.Process.Kill() }
			return fmt.Sprint(localAddr.Port), kill, nil
		}
	}
	return "", nil, fmt.Errorf("iap tunnel startup timed out")
}

// DestroyVM sends a request to delete a VM. Actual VM description is
// currently (2015-01-19) very slow for no good reason. This function
// returns once it's been requested, not when it's done.
func DestroyVM(ts oauth2.TokenSource, proj, zone, instance string) error {
	computeService, _ := compute.New(oauth2.NewClient(context.TODO(), ts))
	apiGate()
	_, err := computeService.Instances.Delete(proj, zone, instance).Do()
	return err
}

type VM struct {
	// Name is the name of the GCE VM instance.
	// For example, it's of the form "mote-bradfitz-plan9-386-foo",
	// and not "plan9-386-foo".
	Name   string
	IPPort string
	TLS    KeyPair
	Type   string // buildlet type
}

// ListVMs lists all VMs.
func ListVMs(ts oauth2.TokenSource, proj, zone string) ([]VM, error) {
	var vms []VM
	computeService, _ := compute.New(oauth2.NewClient(context.TODO(), ts))

	// TODO(bradfitz): paging over results if more than 500
	apiGate()
	list, err := computeService.Instances.List(proj, zone).Do()
	if err != nil {
		return nil, err
	}
	for _, inst := range list.Items {
		if inst.Metadata == nil {
			// Defensive. Not seen in practice.
			continue
		}
		meta := map[string]string{}
		for _, it := range inst.Metadata.Items {
			if it.Value != nil {
				meta[it.Key] = *it.Value
			}
		}
		hostType := meta["buildlet-host-type"]
		if hostType == "" {
			continue
		}
		vm := VM{
			Name: inst.Name,
			Type: hostType,
			TLS: KeyPair{
				CertPEM: meta["tls-cert"],
				KeyPEM:  meta["tls-key"],
			},
		}
		_, extIP := instanceIPs(inst)
		if extIP == "" || vm.TLS.IsZero() {
			continue
		}
		vm.IPPort = extIP + ":443"
		vms = append(vms, vm)
	}
	return vms, nil
}

func instanceIPs(inst *compute.Instance) (intIP, extIP string) {
	for _, iface := range inst.NetworkInterfaces {
		if strings.HasPrefix(iface.NetworkIP, "10.") {
			intIP = iface.NetworkIP
		}
		for _, accessConfig := range iface.AccessConfigs {
			if accessConfig.Type == "ONE_TO_ONE_NAT" {
				extIP = accessConfig.NatIP
			}
		}
	}
	return
}

var (
	cosListMu      sync.Mutex
	cosCachedTime  time.Time
	cosCachedImage string
)

// cosImage returns the GCP VM image name of the latest stable
// Container-Optimized OS image. It caches results for 15 minutes.
func cosImage(ctx context.Context, svc *compute.Service) (string, error) {
	const cacheDuration = 15 * time.Minute
	cosListMu.Lock()
	defer cosListMu.Unlock()
	if cosCachedImage != "" && cosCachedTime.After(time.Now().Add(-cacheDuration)) {
		return cosCachedImage, nil
	}

	imList, err := svc.Images.List("cos-cloud").Filter(`(family eq "cos-stable")`).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	if imList.NextPageToken != "" {
		return "", fmt.Errorf("too many images; pagination not supported")
	}
	ims := imList.Items
	if len(ims) == 0 {
		return "", errors.New("no image found")
	}
	sort.Slice(ims, func(i, j int) bool {
		if ims[i].Deprecated == nil && ims[j].Deprecated != nil {
			return true
		}
		return ims[i].CreationTimestamp > ims[j].CreationTimestamp
	})

	im := ims[0].SelfLink
	cosCachedImage = im
	cosCachedTime = time.Now()
	return im, nil
}
