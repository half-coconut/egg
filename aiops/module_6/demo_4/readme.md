删除 deployment 通过 name

```shell
chenchen@chenchendeMacBook-Pro-2 demo_1 % kubectl delete deployment "nginx-deployment" 
deployment.apps "nginx-deployment" deleted

```

### 嵌入 yaml

在 Go 语言中，//go:embed 是一个指令，

- 用于将文件嵌入到 Go 程序的二进制文件中。
- 这个特性是从 Go 1.16 版本开始引入的。
- 允许开发者将静态文件（如 YAML、HTML、CSS、JavaScript 等）直接嵌入到 Go 代码中，而无需在运行时单独加载这些文件。

`//go:embed deploy.yaml` 的意思是：

    deploy.yaml 是要被嵌入的文件名。
    使用 //go:embed 指令，Go 编译器会将 deploy.yaml 文件的内容嵌入到 Go 程序中。

### GVK (Group-Version-Kind)

GVK 是 Group-Version-Kind 的简称。它是 Kubernetes 中一种标识资源类型的机制。GVK的结构包括三个部分：

- Group：资源所属的 API 组。Kubernetes API 将资源分为多个组（如 apps, v1, extensions 等）。
- Version：API 版本，表示该资源定义的版本。例如，v1、v1beta1 等。
- Kind：资源的具体类型，例如 Pod、Deployment、Service 等。

例如，一个 Deployment 的 GVK 表示为：

```shell
Group: apps
Version: v1
Kind: Deployment
```

由此，可以唯一地识别 Kubernetes 集群中的特定资源类型。

### GCR (Google Container Registry)

GCR 是 Google Container Registry 的缩写，是 Google Cloud 提供的一项服务，用于存储和管理 Docker 容器镜像。GCR
是一个安全、易于访问的私有和公共容器注册表服务，允许开发者在 Google Cloud 平台上构建、存储和部署容器化应用程序。

GCR 的主要特点包括：

- 安全性：提供身份验证和授权，以确保只有授权用户可以访问和推送镜像。
- 集成：与 Google Cloud Platform 的其他服务（如 Google Kubernetes Engine）无缝集成。
- 高可用性和可靠性：自动备份和高可用性设计。
- 镜像版本控制：支持镜像的标签和版本管理。

总结

    GVK 用于标识 Kubernetes 中的各种资源类型，帮助 Kubernetes 用户理解资源的类别和版本。
    GCR 是一个用于存储和管理容器镜像的注册表服务，属于 Google Cloud 平台的一部分。
