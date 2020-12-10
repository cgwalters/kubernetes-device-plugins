package main

import (
	"flag"

	"github.com/cgwalters/kubernetes-device-plugins/pkg/kvm"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	flag.Parse()

	manager := dpm.NewManager(kvm.KVMLister{})
	manager.Run()
}
