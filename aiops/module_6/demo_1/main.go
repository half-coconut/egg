package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	// kubectl get pods
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	// kubectl get deployment
	deplyments, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, deployment := range deplyments.Items {
		fmt.Println(deployment.Name)
	}
}

//func LocalConfig(t *testing.T) {
//	// 集群外的访问方式，通过 config 来交互
//	kubeconfig := flag.String("kubeconfig", "/Users/chenchen/.kube/config", "Path to a kubeconfig. Only requirement...")
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		fmt.Println(err)
//	}
//	//var config *rest.Config
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// kubectl get pods
//	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
//	for _, pod := range pods.Items {
//		fmt.Println(pod.Name)
//	}
//	// kubectl get deployment
//	deplyments, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
//	for _, deployment := range deplyments.Items {
//		fmt.Println(deployment.Name)
//	}
//}
