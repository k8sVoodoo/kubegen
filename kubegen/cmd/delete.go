package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [resource] [name]",
	Short: "Delete a Kubernetes resource",
	Long:  `Delete a specified Kubernetes resource by its name, such as deployment or service.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		resource, name := args[0], args[1]

		// Initialize Kubernetes client
		config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			fmt.Printf("Error loading kubeconfig: %v\n", err)
			os.Exit(1)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("Error creating Kubernetes client: %v\n", err)
			os.Exit(1)
		}

		// Call delete logic based on resource type
		err = deleteResource(clientset, resource, name)
		if err != nil {
			fmt.Printf("Error deleting resource: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Resource %s/%s deleted successfully\n", resource, name)
	},
}

func deleteResource(clientset *kubernetes.Clientset, resource, name string) error {
	switch resource {
	case "deployment":
		return clientset.AppsV1().Deployments("default").Delete(context.TODO(), name, metav1.DeleteOptions{})
	case "service":
		return clientset.CoreV1().Services("default").Delete(context.TODO(), name, metav1.DeleteOptions{})
	// Add cases for other resources as needed
	default:
		return fmt.Errorf("unsupported resource type: %s", resource)
	}
}

func init() {
	rootCmd.AddCommand(DeleteCmd)
}