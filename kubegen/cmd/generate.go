package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	resourceType string
	resourceName string
	namespace    string
	image        string
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Kubernetes configuration",
	Long: `Generate a Kubernetes configuration file based on the specified type and name. 

You can specify the type of Kubernetes resource (e.g., deployment, statefulset, etc.), the resource name, and optionally a namespace and container image. If the namespace is not provided, it defaults to 'default'. If the image is not specified, it defaults to 'my-image:latest'.

Examples:
  kubegen generate --type deployment --name my-deployment --namespace my-namespace --image my-image:new-tag
  kubegen generate -t service -n my-service -N my-namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		if resourceType == "" {
			fmt.Println("Error: resource type is required. Use --type flag.")
			os.Exit(1)
		}

		if resourceName == "" {
			fmt.Println("Error: resource name is required. Use --name flag.")
			os.Exit(1)
		}

		if namespace == "" {
			fmt.Println("Warning: namespace is not specified. Defaulting to 'default'.")
			namespace = "default"
		}

		if image == "" {
			fmt.Println("Warning: image is not specified. Defaulting to 'my-image:latest'.")
			image = "my-image:latest"
		}

		templateFile := fmt.Sprintf("templates/%s.yaml", resourceType)
		if _, err := os.Stat(templateFile); os.IsNotExist(err) {
			fmt.Printf("Error: template for %s not found\n", resourceType)
			os.Exit(1)
		}

		tmpl, err := template.ParseFiles(templateFile)
		if err != nil {
			fmt.Println("Error loading template:", err)
			os.Exit(1)
		}

		fileName := fmt.Sprintf("generated-%s-%s.yaml", resourceType, resourceName)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()

		data := map[string]interface{}{
			"Name":      resourceName,
			"Namespace": namespace,
			"Image":     image,
			"Replicas":  3,
			"Host":      "example.com",
			"Key":       "base64encodedkey",
			"Value":     "configValue",
		}
		err = tmpl.Execute(file, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
			os.Exit(1)
		}

		fmt.Printf("Generated %s configuration in '%s'\n", resourceType, fileName)
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&resourceType, "type", "t", "", "Type of Kubernetes resource to generate (e.g., deployment, statefulset, daemonset, service, httpproxy, ingress, secrets, configmap)")
	GenerateCmd.Flags().StringVarP(&resourceName, "name", "n", "", "Name of the Kubernetes resource")
	GenerateCmd.Flags().StringVarP(&namespace, "namespace", "N", "", "Namespace for the Kubernetes resource")
	GenerateCmd.Flags().StringVarP(&image, "image", "i", "", "Container image to use (e.g., my-image:latest)") // Added flag for image
}
