### k8s自动扩容和自愈

#### kubectl create 命令

```shell
kubectl create deployment hello-world-flask --image=lyzhang1999/hello-world-flask:latest --replicas=2 
deployment.apps/hello-world-flask created
```

hello-world-flask 代表工作负载的名称，–image 代表镜像，day1 部署的镜像，–replicas 指的是 Pod 副本数，可以把它类比为弹性伸缩组的
VM 数量。
注意：这条命令会生成 Deployment Manifest，然后自动执行 kubectl apply 将 Manifest 应用到集群内
省略了我们手动编写 Manifest 的过程。还可以为上面的命令增加 --dry-run 和 -o 参数，单纯输出 Manifest 内容：

```shell
kubectl create deployment hello-world-flask --image lyzhang1999/hello-world-flask:latest --replicas=2 --dry-run=client -o yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: hello-world-flask
  name: hello-world-flask
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello-world-flask
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: hello-world-flask
    spec:
      containers:
      - image: lyzhang1999/hello-world-flask:latest
        name: hello-world-flask
        resources: {}
status: {}
```

使用 kubectl create service 命令创建 Service：

```shell
kubectl create service clusterip hello-world-flask --tcp=5000:5000
service/hello-world-flask created
```

使用 kubectl create ingress 命令创建 Ingress：

```shell
kubectl create ingress hello-world-flask --rule="/=hello-world-flask:5000"
ingress.networking.k8s.io/hello-world-flask created
```

部署 Ingress-nginx：

```shell
kubectl create -f https://ghproxy.com/https://raw.githubusercontent.com/lyzhang1999/resource/main/ingress-nginx/ingress-nginx.yaml
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
serviceaccount/ingress-nginx-admission created
......
```

- Pod 会被 Deployment 工作负载管理起来，例如创建和销毁等；
- Service 相当于弹性伸缩组的负载均衡器，它能以**加权轮训**的方式将流量转发到多个 Pod 副本上；
- Ingress 相当于集群的外网访问入口。

#### K8s 自愈实验

```shell
kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
hello-world-flask-56fbff68c8-2xz7w   1/1     Running   0          3m38s
hello-world-flask-56fbff68c8-4f9qz   1/1     Running   0          3m38s
```

有了 Ingress，我们访问 Pod 就不再需要进行端口转发了，我们可以直接访问 127.0.0.1。
下面的命令会每隔 1 秒钟发送一次请求，并打印出时间和返回内容：

```shell
while true; do sleep 1; curl http://127.0.0.1; echo -e '\n'$(date);done
Hello, my first docker images! hello-world-flask-56fbff68c8-4f9qz
2022年 9月 7日 星期三 19时21分03秒 CST
Hello, my first docker images! hello-world-flask-56fbff68c8-2xz7w
2022年 9月 7日 星期三 19时21分04秒 CST
```

以下操作，模拟进程意外中止导致宕机的情况

```shell
kubectl exec -it hello-world-flask-56fbff68c8-2xz7w -- bash -c "killall python3"
```

所有的请求流量都被转发到了没有故障的 Pod，也就是说，故障成功地被转移了

```shell
Hello, my first docker images! hello-world-flask-56fbff68c8-4f9qz
2022年 9月 7日 星期三 19时27分44秒 CST
Hello, my first docker images! hello-world-flask-56fbff68c8-4f9qz
2022年 9月 7日 星期三 19时27分45秒 CST
Hello, my first docker images! hello-world-flask-56fbff68c8-4f9qz
```

#### K8s 的自动扩容

自动扩容依赖于 K8s Metric Server 提供的监控指标，首先我们需要安装它：

```shell
kubectl apply -f https://ghproxy.com/https://raw.githubusercontent.com/lyzhang1999/resource/main/metrics/metrics.yaml
serviceaccount/metrics-server created
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader created
clusterrole.rbac.authorization.k8s.io/system:metrics-server created
......
```

安装完成后，等待 Metric 工作负载就绪：

```shell
kubectl wait deployment -n kube-system metrics-server --for condition=Available=True --timeout=90s
deployment.apps/metrics-server condition met
```

Metric Server 就绪后，通过 kubectl autoscale 命令来为 Deployment 创建自动扩容策略：

```shell
kubectl autoscale deployment hello-world-flask --cpu-percent=50 --min=2 --max=10
```

其中，–cpu-percent 表示 CPU 使用率阈值，当 CPU 超过 50% 时将进行自动扩容，–min 代表最小的 Pod 副本数，–max
代表最大扩容的副本数。也就是说，自动扩容会根据 CPU 的使用率在 2 个副本和 10 个副本之间进行扩缩容。

要使自动扩容生效，还需要为我们刚才部署的 hello-world-flask Deployment 设置资源配额。你可以通过下面的命令来配置：

```shell
$ kubectl patch deployment hello-world-flask --type='json' -p='[{"op": "add", "path": "/spec/template/spec/containers/0/resources", "value": {"requests": {"memory": "100Mi", "cpu": "100m"}}}]'
deployment.apps/hello-world-flask patched
```

现在，Deployment 将会重新创建两个新的 Pod，你可以使用下面的命令筛选出新的 Pod：

```shell
kubectl get pod --field-selector=status.phase==Running
NAME                                 READY   STATUS    RESTARTS   AGE
hello-world-flask-64dd645c57-4clbp   1/1     Running   0          117s
hello-world-flask-64dd645c57-cc6g6   1/1     Running   0          117s
```

选择一个 Pod 并使用 kubectl exec 进入到容器内：

```shell
kubectl exec -it hello-world-flask-64dd645c57-4clbp -- bash
root@hello-world-flask-64dd645c57-4clbp:/app#
```

我们模拟业务高峰期场景，使用 ab 命令来创建并发请求：

```shell
root@hello-world-flask-64dd645c57-4clbp:/app# ab -c 50 -n 10000 http://127.0.0.1:5000/
```

打开一个新的命令行窗口，使用下面的命令来持续监控 Pod 的状态：

```shell
$ kubectl get pods --watch
NAME                                 READY   STATUS    RESTARTS   AGE
hello-world-flask-64dd645c57-9x869   1/1     Running   0          4m6s
hello-world-flask-64dd645c57-vw8nc   0/1     Pending   0          0s
hello-world-flask-64dd645c57-46b6s   0/1     ContainerCreating   0          0s
hello-world-flask-64dd645c57-vw8nc   1/1     Running             0          18s
```

--watch 表示持续监听 Pod 状态变化。在 ab 压力测试的过程中，会不断创建新的 Pod 副本，这说明 K8s 已经感知到了 Pod
的业务压力，并且正在自动进行横向扩容。