package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a Kubernetes configuration",
	Long: `Apply a Kubernetes configuration file to the cluster. 

The configuration file should be a valid Kubernetes YAML or JSON file. 

Examples:
  kubegen apply path/to/generated-deployment-my-deployment.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: no file specified. Provide a path to the configuration file.")
			os.Exit(1)
		}

		filePath := args[0]

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Error: file '%s' does not exist\n", filePath)
			os.Exit(1)
		}

		kubectlCmd := exec.Command("kubectl", "apply", "-f", filePath)
		output, err := kubectlCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error applying configuration: %s\n", err)
			fmt.Println(string(output))
			os.Exit(1)
		}

		fmt.Printf("Successfully applied configuration from '%s'\n", filePath)
		fmt.Println(string(output))
	},
}
