# 002

以 `nvidia/cuda` 构建镜像，运行显卡 GPU 的算法应用

# 构建镜像，启动容器

```bash
docker build -f ./examples/002/Dockerfile -t d002 .
docker run --name d002 d002
```

# 删除容器，删除镜像

```bash
docker rm d002
docker rmi d002
```
