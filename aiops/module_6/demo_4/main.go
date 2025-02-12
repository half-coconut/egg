package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
)

//go:embed deploy.yaml
var deployYaml string

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/chenchen/.kube/config", "Path to a kubeconfig. Only requirement...")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("1 err %s\n", err.Error())
	}
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	//}
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err)
	//}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("2 err %s\n", err.Error())
	}

	// 解析 yaml 转成 Unstructured
	deployObj := &unstructured.Unstructured{}
	if err = yaml.Unmarshal([]byte(deployYaml), deployObj); err != nil {
		fmt.Printf("3 err %s\n", err.Error())
	}
	// 从 deployObj 中获取 GVK
	apiVersion, found, err := unstructured.NestedString(deployObj.Object, "apiVersion")
	if err != nil || !found {
		fmt.Printf("apiVersion not found: %s", err.Error())
	}

	kind, found, err := unstructured.NestedString(deployObj.Object, "kind")
	if err != nil || !found {
		log.Printf("kind not found: %s", err.Error())
	}

	// GVK -> GVR，制作 GVR
	gvr := schema.GroupVersionResource{}
	versionParts := strings.Split(apiVersion, "/")
	if len(versionParts) == 2 {
		gvr.Group = versionParts[0]
		gvr.Version = versionParts[1]
	} else {
		gvr.Version = versionParts[0]
	}

	switch kind {
	case "Deployment":
		gvr.Resource = "deployments"
	case "Pod":
		gvr.Resource = "pods"
	default:
		fmt.Printf("Unsupported kind %s", kind)
	}
	// 使用 dynamicClient 创建资源
	_, err = dynamicClient.Resource(gvr).Namespace("default").Create(context.TODO(), deployObj, v1.CreateOptions{})
	if err != nil {
		fmt.Printf("4 err: %s\n", err.Error())
	}
	log.Println("Create resource successfully")
}
