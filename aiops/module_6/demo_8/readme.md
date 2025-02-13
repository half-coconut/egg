

```shell
# 先创建好 resource
kubectl apply -f crd.yaml
kubectl apply -f myresource.yaml

chenchen@chenchendeMacBook-Pro-2 demo_8 % go run main.go get myresource
Name: my-resource-instance, Namespace: default, UID: 1ad0f10e-0072-40d1-baee-9f16ed7d9aad

chenchen@chenchendeMacBook-Pro-2 demo_8 % kubectl get myresource
NAME                   AGE
my-resource-instance   12m

```