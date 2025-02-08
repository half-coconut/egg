### Client-go 简介和配置

#### In Cluster Configuration

 ```shell
 kubectl exec -it <pod-address> /bin/sh
 cd /var/run/secrets/kubernetes.io/serviceaccount
# ls
ca.crt  namespace  token
 ```

```shell
docker build -t halfcoconut/hello-egg:eggV1 . --push
docker build -t halfcoconut/hello-egg:in-cluster-v1  . --push

kubectl delete -f deployment.yaml 

```
Role 权限问题： User "system:serviceaccount:default:default" 没有权限
```shell
chenchen@chenchendeMacBook-Pro-2 demo_1 % kubectl logs demo-1-6d9fbb798-wmqln
deployments.apps is forbidden: User "system:serviceaccount:default:default" cannot list resource "deployments" in API group "apps" in the namespace "default"

```

赋予权限
```shell
kubectl create role demo --resource pods,deployment --verb list

kubectl create rolebinding demo --role demo --serviceaccount default:default
```
有了权限后，就可以正常打印了：
```shell
chenchen@chenchendeMacBook-Pro-2 demo_1 % kubectl logs demo-1-6d9fbb798-khm5p
db-74574d66dd-whj5x
demo-1-6d9fbb798-khm5p
redis-6c5fb9c4b7-8qk97
result-5f99548f7c-fwwxl
vote-5d74dcd7c7-g99zc
worker-6f5f6cdd56-gxjsn
db
demo-1
redis
result
vote
worker
chenchen@chenchendeMacBook-Pro-2 demo_1 % 


```