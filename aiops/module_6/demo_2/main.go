package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/chenchen/.kube/config", "Path to a kubeconfig. Only requirement...")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 初始化 restclient
	restclient, err := rest.RESTClientFor(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 获取 pod 列表
	podList := corev1.PodList{}
	err = restclient.Get().Namespace("kube-system").Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).Do(context.Background()).Into(&podList)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, pod := range podList.Items {
		fmt.Println(pod.Name, pod.Namespace, pod.Status.Phase)
	}
}
