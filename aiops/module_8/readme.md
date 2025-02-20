## Operator


安装 kubebuilder

https://book.kubebuilder.io/quick-start#installation
```shell
# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/

```

步骤

- mkdir application && cd application
- go mod init github.com/lyzhang1999/application
- kubebuilder init --domain=aiops.com
- kubebuilder create api --group application --version v1 --kind Application
  - make manifests
  - make install
  - make run // 本地运行，kubebuilder 可以处理权限的问题
  - > kubectl apply -f application_v1_application.yaml
  - > kubectl get deployment <deployment-name> -oyaml
  - > kubectl get deployment application-sample -oyaml
  - 修改 samples 里的内容，比如镜像版本号等，重现部署后，获取 yaml 文件，可以看到已经变更了
    - ```shell
      chenchen@chenchendeMacBook-Pro-2 samples % kubectl apply -f application_v1_application.yaml
      application.application.aiops.com/application-sample created
      chenchen@chenchendeMacBook-Pro-2 samples % kubectl get deployment
      NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
      application-sample   1/1     1            1           23s
      chenchen@chenchendeMacBook-Pro-2 samples % kubectl get service   
      NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
      application-sample   ClusterIP   10.96.166.106   <none>        80/TCP           52s
      chenchen@chenchendeMacBook-Pro-2 samples % kubectl get ingress
      NAME                 CLASS    HOSTS                   ADDRESS     PORTS   AGE
      application-sample   nginx    application.aiops.com   localhost   80      68s
      ```

  - > kubectl delete -f application_v1_application.yaml
- 完善 cronhpa/api/v1/application_types.go
  - 修改 ApplicationSpec
- 生成 CRD：make manifests，查看 config/crd/bases/application.aiops.com_applications.yaml 文件 
- 编写 internal/controller/application_controller.go Reconcile 业务逻辑 
- 将 CRD 安装到集群：make install 
- 运行 Operator：make run 
- 部署示例：kubectl apply -f config/samples/application_v1_application.yaml
