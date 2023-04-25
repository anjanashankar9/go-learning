package main

import (
	"fmt"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	namespace = "l7cplane"
	CM_NAME   = "anjana-test-configmap"
)

func connectK8sCluster() (*kubernetes.Clientset, error) {

	// Connect to cluster
	//config, err := rest.InClusterConfig()

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}

	// Create a Kubernetes client
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}
	return clientSet, err
}

func main() {
	// Load the kubeconfig file
	clientset, err := connectK8sCluster()
	if err != nil {
		panic(err)
	}

	//Reading configmap
	cm, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), CM_NAME, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Print(cm)

	//Watching configmap
	var (
		mutex *sync.Mutex
	)

	mutex = &sync.Mutex{}
	go watchForChanges(clientset, namespace, mutex)

	time.Sleep(6000 * time.Second)
}
func watchForChanges(clientset *kubernetes.Clientset, namespace string, mutex *sync.Mutex) {
	for {
		watcher, err := clientset.CoreV1().ConfigMaps(namespace).Watch(
			context.TODO(),
			metav1.SingleObject(metav1.ObjectMeta{
				Name: CM_NAME, Namespace: namespace}))
		if err != nil {
			panic("Unable to create watcher")
		}
		printNew(watcher.ResultChan(), mutex)
	}
}

func printNew(eventChannel <-chan watch.Event, mutex *sync.Mutex) {
	for {
		event, open := <-eventChannel
		if open {
			switch event.Type {
			case watch.Added:
				fallthrough
			case watch.Modified:
				mutex.Lock()
				// Update our endpoint
				if updatedMap, ok := event.Object.(*corev1.ConfigMap); ok {
					if policyVersion, ok := updatedMap.Data["version"]; ok {
						fmt.Println("new policyversion ", policyVersion)
					}
				}
				mutex.Unlock()
			case watch.Deleted:
				mutex.Lock()
				// Fall back to the default value
				fmt.Println("Configmap deleted")
				mutex.Unlock()
			default:
				fmt.Println("In the default case")
			}
		}
	}
}
