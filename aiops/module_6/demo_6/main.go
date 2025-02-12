package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"
)

// 使用 Informer 替代 watch
func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("1 err %s\n", err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(clientSet, time.Hour*12)

	// Deployment
	deploymentInformer := informerFactory.Apps().V1().Deployments()
	informer := deploymentInformer.Informer()
	deployLister := deploymentInformer.Lister()
	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Deployment added")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Deployment updated")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Deployment deleted")
		},
	})
	if err != nil {
		panic(err)
	}
	// Service
	serviceInformer := informerFactory.Core().V1().Services()
	informer = serviceInformer.Informer()
	//serviceLister := serviceInformer.Lister()
	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Service added")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Service updated")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Service deleted")
		},
	})
	if err != nil {
		panic(err)
	}
	stopper := make(chan struct{})
	defer close(stopper)

	// 启动 Informer，List & Watch，可以使用缓存，性能更好，解决一些难点的问题
	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)
	deployments, err := deployLister.Deployments("default").List(labels.Everything())
	if err != nil {
		fmt.Printf("err %s\n", err.Error())
	}
	for idx, deploy := range deployments {
		fmt.Println(idx, deploy.Name)
	}
	<-stopper
}
