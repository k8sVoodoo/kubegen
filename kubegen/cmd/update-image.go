package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	updateImageType    string
	updateImageName    string
	updateImageVersion string
)

var UpdateImageCmd = &cobra.Command{
	Use:   "update-image",
	Short: "Update the container image in a Kubernetes resource",
	Long: `Update the container image in a specified Kubernetes resource.

Supported resource types are 'deployment', 'daemonset', and 'statefulset'. 

Examples:
  kubegen update-image --type deployment --name my-deployment --version my-image:new-tag
  kubegen update-image -t daemonset -n my-daemonset -v my-image:latest
`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateImageType == "" || updateImageName == "" || updateImageVersion == "" {
			fmt.Println("Error: --type, --name, and --version flags are required")
			os.Exit(1)
		}

		var kubectlCmd *exec.Cmd
		switch updateImageType {
		case "deployment":
			kubectlCmd = exec.Command("kubectl", "set", "image", fmt.Sprintf("deployment/%s", updateImageName), fmt.Sprintf("*=%s", updateImageVersion))
		case "daemonset":
			kubectlCmd = exec.Command("kubectl", "set", "image", fmt.Sprintf("daemonset/%s", updateImageName), fmt.Sprintf("*=%s", updateImageVersion))
		case "statefulset":
			kubectlCmd = exec.Command("kubectl", "set", "image", fmt.Sprintf("statefulset/%s", updateImageName), fmt.Sprintf("*=%s", updateImageVersion))
		default:
			fmt.Println("Error: unsupported resource type. Use 'deployment', 'daemonset', or 'statefulset'.")
			os.Exit(1)
		}

		output, err := kubectlCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error updating image: %s\n", err)
			fmt.Println(string(output))
			os.Exit(1)
		}

		fmt.Printf("Successfully updated image to '%s' in %s '%s'\n", updateImageVersion, updateImageType, updateImageName)
		fmt.Println(string(output))
	},
}

func init() {
	UpdateImageCmd.Flags().StringVarP(&updateImageType, "type", "t", "", "Type of Kubernetes resource (e.g., deployment, daemonset, statefulset)")
	UpdateImageCmd.Flags().StringVarP(&updateImageName, "name", "n", "", "Name of the Kubernetes resource")
	UpdateImageCmd.Flags().StringVarP(&updateImageVersion, "version", "v", "", "New container image version (e.g., my-image:latest)")
}
