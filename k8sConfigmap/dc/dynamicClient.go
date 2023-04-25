package main

import (
	"fmt"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
)

const (
	namespace = "l7cplane"
	CM_NAME2  = "policy-version-configmap-v000"
)

func connectK8sClusterDynamic() (dynamic.Interface, error) {

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

	clientSet, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		klog.Error("connectk8scluster(): failed to get clientset ", err)
		return nil, err
	}
	return clientSet, err
}

var gvr = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "configmaps",
}

func main() {
	// Load the kubeconfig file
	clientset, err := connectK8sClusterDynamic()
	if err != nil {
		panic(err)
	}

	// Get the ConfigMap using the dynamic client
	configMap, err := clientset.Resource(gvr).Namespace(namespace).Get(context.TODO(), CM_NAME2, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	obj := &unstructured.Unstructured{}
	// Decode YAML to unstructured object.
	var decUnstructured = yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	////Watching configmap
	//var (
	//	mutex *sync.Mutex
	//)
	//
	//mutex = &sync.Mutex{}
	//go watchForChanges2(clientset, namespace, mutex)
	//
	//time.Sleep(6000 * time.Second)
}

//func watchForChanges2(clientset *kubernetes.Clientset, namespace string, mutex *sync.Mutex) {
//	for {
//		watcher, err := clientset.CoreV1().ConfigMaps(namespace).Watch(
//			context.TODO(),
//			metav1.SingleObject(metav1.ObjectMeta{
//				Name: CM_NAME, Namespace: namespace}))
//		if err != nil {
//			panic("Unable to create watcher")
//		}
//		printNew2(watcher.ResultChan(), mutex)
//	}
//}
//
//func printNew2(eventChannel <-chan watch.Event, mutex *sync.Mutex) {
//	for {
//		event, open := <-eventChannel
//		if open {
//			switch event.Type {
//			case watch.Added:
//				fallthrough
//			case watch.Modified:
//				mutex.Lock()
//				// Update our endpoint
//				if updatedMap, ok := event.Object.(*corev1.ConfigMap); ok {
//					if policyVersion, ok := updatedMap.Data["version"]; ok {
//						fmt.Println("new policyversion ", policyVersion)
//					}
//				}
//				mutex.Unlock()
//			case watch.Deleted:
//				mutex.Lock()
//				// Fall back to the default value
//				fmt.Println("Configmap deleted")
//				mutex.Unlock()
//			default:
//				fmt.Println("In the default case")
//			}
//		}
//	}
//}
