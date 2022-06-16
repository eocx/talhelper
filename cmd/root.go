package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var version string

var rootLongHelp = strings.TrimSpace(`
talhelper is a tool to help you create a Talos cluster.

Currently there is only one usage, which is creating a cluster config YAML file.

Workflow:
  Create talconfig.yaml file defining your nodes information like so:
  ----------------------------------------
  clustername: mycluster
  talosVersion: v1.0
  endpoint: https://192.168.200.10:6443
  nodes:
    - hostname: master1
      ipAddress: 192.168.200.11
      installDisk: /dev/sdb
      controlPlane: true
    - hostname: worker1
      ipAddress: 192.168.200.21
      installDisk: /dev/nvme1
      controlPlane: false
  ----------------------------------------

  Then run these commands:
  > talhelper gensecret --patch-configfile > talenv.sops.yaml
  > sops -e -i talenv.sops.yaml
  > talhelper genconfig"

  The generated yaml files will be in ./clusterconfig directory
`)

var rootCmd = &cobra.Command{
	Use:           "talhelper",
	Short:         "A tool to help with creating Talos cluster",
	Long:          rootLongHelp,
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version,
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
