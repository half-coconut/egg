package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// 实现 client-go watch 方法，生产不推荐直接使用，有超时机制。
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
	timeOut := int64(60)
	watcher, err := clientSet.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{TimeoutSeconds: &timeOut})
	if err != nil {
		panic(err)
	}
	for event := range watcher.ResultChan() {
		item := event.Object.(*corev1.Pod)

		switch event.Type {
		case watch.Added:
			processPod(item.GetName())
			//case watch.Modified:
			//case watch.Deleted:
		}
	}
}

func processPod(name string) {
	fmt.Println("Processing pod", name)
}
