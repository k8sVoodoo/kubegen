package main

import (
	"fmt"
	"os"

	"kubegen/cmd" // Adjust the import path as necessary

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "kubegen",
		Short: "kubegen is a CLI tool for generating, applying, and updating Kubernetes configurations",
		Long: `kubegen is a CLI tool that helps with generating Kubernetes configurations,
applying them to a cluster, and updating container images in deployments, daemonsets, and statefulsets.

You can use the following commands:
  generate      Generate a Kubernetes configuration
  apply         Apply a Kubernetes configuration
  update-image  Update the container image in a Kubernetes resource

For more details on each command, use the --help flag.`,
	}

	rootCmd.AddCommand(cmd.GenerateCmd)
	rootCmd.AddCommand(cmd.ApplyCmd)
	rootCmd.AddCommand(cmd.UpdateImageCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
