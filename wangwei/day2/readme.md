### 将容器镜像部署到k8s

#### 安装 Kubernetes

- 本地安装 K8s 集群的办法：Kind
- 安装 Kind
- https://kind.sigs.k8s.io/docs/user/quick-start/#installation

`config.yaml`

```shell
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP

```

#### 创建集群

执行 kind create 命令，创建 K8s 集群：

```shell
kind create cluster --config config.yaml
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.23.4) 🖼
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind
```

#### Manifest

用来向 K8s 描述“期望最终状态”的文件，就叫做 K8s Manifest，也可以称之为清单文件。

- Manifest 就好比餐厅的菜单，你只管点菜，做菜的过程我不管
- 需要聚焦四个字段，分别是 Kind、Containers、Image、Ports

#### 部署容器镜像到 K8s

`flask-pod.yaml`

```shell
apiVersion: v1
kind: Pod
metadata:
  name: hello-world-flask
spec:
  containers:
    - name: flask
      image: lyzhang1999/hello-world-flask:latest
      ports:
        - containerPort: 5000

```

通过 kubectl apply 应用这段 Manifest：

```shell
kubectl apply -f flask-pod.yaml
pod/hello-world-flask created
```

#### 访问 Pod

```shell
kubectl get pods
NAME                READY   STATUS    RESTARTS   AGE
hello-world-flask   1/1     Running   0          1m
```

使用 kubectl port-forward 命令进行端口转发操作，打通容器和本地网络
注意：本地和集群的网络是隔离的，自然不能在集群外部，也就是本地电脑访问 Pod

```shell
kubectl port-forward pod/hello-world-flask 8000:5000
Forwarding from 127.0.0.1:8000 -> 5000
Forwarding from [::1]:8000 -> 5000
```

进入 Pod 容器

```shell
kubectl exec -it hello-world-flask -- bash
root@hello-world-flask:/app#
```

删除 Pod

```shell
kubectl delete pod hello-world-flask
pod "hello-world-flask" deleted
```