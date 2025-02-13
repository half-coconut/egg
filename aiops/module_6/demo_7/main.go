package main

import (
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"path/filepath"
	"time"
)

type Controller struct {
	queue    workqueue.TypedRateLimitingInterface[string]
	informer cache.Controller
	indexer  cache.Indexer
}

func NewController(queue workqueue.TypedRateLimitingInterface[string], informer cache.Controller, indexer cache.Indexer) *Controller {
	return &Controller{
		queue:    queue,
		informer: informer,
		indexer:  indexer,
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.syncToStdout(key)
	c.handleErr(err, key)
	return true
}

func (c *Controller) syncToStdout(key string) error {
	// 通过 key 直接从 indexer 中获取对象
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Printf("fetching object with key %s from store failed with %v", key, err)
		return err
	}
	if !exists {
		fmt.Printf("deployment %s does not exist anymore\n", key)
	} else {
		deployment := obj.(*appsv1.Deployment)
		fmt.Printf("sync object %s\n", deployment.GetName())
		if deployment.Name == "nginx-deployment" {
			return fmt.Errorf("nginx-deployment is not allowed\n")
		}
	}
	return nil
}

func (c *Controller) handleErr(err error, key string) {
	if err == nil {
		c.queue.Forget(key)
		return
	}
	if c.queue.NumRequeues(key) < 5 {
		fmt.Printf("Retry %d for key %s\n", c.queue.NumRequeues(key), key)
		//fmt.Printf("error syncing %q: %v", key, err)
		// 重新加入队列，并且进行速率限制，这会让他多一段时间才会被处理，避免过度重试
		c.queue.AddRateLimited(key)
		return
	}
	c.queue.Forget(key)
	fmt.Printf("Dropping deployment %q out of the queue: %v\n", key, err)
}

func main() {
	var (
		err        error
		kubeconfig *string
		config     *rest.Config
	)

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("1 err %s\n", err.Error())
	}
	informerFactory := informers.NewSharedInformerFactory(clientSet, time.Hour*60)

	// ratelimitqueue
	queue := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[string]())

	// 对 deployment 监听
	deployInformer := informerFactory.Apps().V1().Deployments()
	informer := deployInformer.Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) { onAddDeployment(obj, queue) },
		UpdateFunc: func(oldObj, newObj interface{}) { onUpdateDeployment(newObj, queue) },
		DeleteFunc: func(obj interface{}) { onDeleteDeployment(obj, queue) },
	})
	controller := NewController(queue, informer, deployInformer.Informer().GetIndexer())
	stopper := make(chan struct{})
	defer close(stopper)

	// 启动 Informer
	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)

	// 处理队列的事件
	go func() {
		for {
			if !controller.processNextItem() {
				break
			}
		}
	}()
	<-stopper
}

func onAddDeployment(obj interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err == nil {
		queue.Add(key)
	}
}

func onUpdateDeployment(new interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		queue.Add(key)
	}
}

func onDeleteDeployment(obj interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err == nil {
		queue.Add(key)
	}
}
