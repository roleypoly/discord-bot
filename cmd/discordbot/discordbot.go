package main

import (
	"github.com/roleypoly/common/version"
	"k8s.io/klog"
)

func main() {
	klog.Info(version.StartupInfo("discord-bot"))
}
