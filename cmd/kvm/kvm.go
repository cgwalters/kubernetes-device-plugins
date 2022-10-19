package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/cgwalters/kvm-device-plugin/pkg/kvm"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	flag.Parse()
	glog.V(3).Infof("Starting kvm-device-plugin (new)")

	manager := dpm.NewManager(kvm.KVMLister{})
	manager.Run()
}
