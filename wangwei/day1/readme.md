### 将业务代码构建容器镜像

#### docker 命令

```shell
docker pull lyzhang1999/hello-world-flask:latest

docker images
REPOSITORY                      TAG        IMAGE ID       CREATED          SIZE
lyzhang1999/hello-world-flask               latest     e2b1a18ed1c1   1 minutes ago   116MB

docker run -d -p 8000:5000 lyzhang1999/hello-world-flask:latest
c370825640b6b3669cae20f14e2684ec82b20e4980b329c02b47e47771c931fd

docker ps
CONTAINER ID   IMAGE                                  COMMAND                  CREATED        STATUS        PORTS                       NAMES
c370825640b6   lyzhang1999/hello-world-flask:latest   "python3 -m flask ru…"   1 hours ago   Up 1 hours   0.0.0.0:8000->5000/tcp      xenodochial_black

docker exec -it c370825640b6 bash
root@c370825640b6:/app#


docker stop c370825640b6
```

#### 制作镜像

```shell
docker build -t hello-world-flask .
```
-t 标签，. context 上下文

#### Docker Hub

```shell
docker tag hello-world-flask my_dockerhub_name/hello-world-flask
docker push my_dockerhub_name/hello-world-flask

```

Dockerfile 文件（略）