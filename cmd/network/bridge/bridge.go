package main

import (
	"flag"
	"os"

	"github.com/cgwalters/kvm-device-plugin/pkg/network/bridge"
	"github.com/golang/glog"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	flag.Parse()

	_, bridgesListDefined := os.LookupEnv(bridge.BridgesListEnvironmentVariable)
	if !bridgesListDefined {
		glog.Exit("BRIDGES environment variable must be set in format BRIDGE[,BRIDGE[...]]")
	}

	manager := dpm.NewManager(bridge.BridgeLister{})
	manager.Run()
}
