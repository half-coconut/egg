package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/chenchen/.kube/config", "Path to a kubeconfig. Only requirement...")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 指定 GVR
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	// 获取 Pod 列表
	unStructPodList, err := dynamicClient.Resource(gvr).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("=====: ", unStructPodList)

	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructPodList.UnstructuredContent(), podList)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, pod := range podList.Items {
		fmt.Println(pod.Name, pod.Namespace, pod.Status.Phase)
	}
}
