# 001

编译一个 go 应用，然后在 ubuntu 上运行

> 应用运行前需要 source 加载一个`env.sh`脚本，用于声明环境变量

# 构建镜像，启动容器

```bash
docker build -f ./examples/001/Dockerfile -t d001 .
docker run --name d001 d001
```

# 删除容器，删除镜像

```bash
docker rm d001
docker rmi d001
```

# 错误解决记录

## 1. 构建时错误提示: /bin/sh: 1: source: not found

```Dockerfile
FROM ubuntu:16.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
    && chmod +x /app/env.sh \
    && source /app/env.sh
```

构建失败

```bash
Step 8/10 : RUN chmod +x /app/entrypoint.sh && chmod +x /app/env.sh && source /app/env.sh
---> Running in ab6ab60bf660
/bin/sh: 1: source: not found
The command '/bin/sh -c chmod +x /app/entrypoint.sh && chmod +x /app/env.sh && source /app/env.sh' returned a non-zero code: 127
```

原因

```bash
# https://blog.csdn.net/liurizhou/article/details/86670097

# ubuntu 默认的 Shell 是: /bin/sh
# 通过执行 /bin/sh -l 可以知道指向的是 Dash Shell
# 而 source 是 Bash Shell 的指令
```

解决，指定通过 `/bin/bash -c "命令"` 来指定 Shell 来执行

```Dockerfile
FROM ubuntu:16.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
 && chmod +x /app/env.sh \
 && /bin/bash -c "source /app/env.sh"
```

构建没报错，但运行容器后，发现:source 操作没有生效,环境变量空的....

解决，网上查说可以放到 ~/.bashrc 中, 两种方式:

1. 执行该命令时指定 /bin/bash 环境

```Dockerfile
RUN echo "source /app/env.sh" >> ~/.bashrc \
    && /bin/bash -c "source ~/.bashrc"
```

2. 指定全局 shell 环境

```Dockerfile
RUN echo "source /app/env.sh" >> ~/.bashrc
SHELL ["/bin/bash", "-c"]
RUN && source ~/.bashrc
```

选择第一种进行解决试试

```Dockerfile
FROM ubuntu:16.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
    && chmod +x /app/env.sh
RUN echo "source /app/env.sh" >> ~/.bashrc \
    && /bin/bash -c "source ~/.bashrc"
```

构建成功，运行容器后发现还是没生效，这次进入容器看看是否真的是没生效:

```bash
# 进入容器
docker exec -it d001 /bin/bash

# 查看 ~/.bashrc 末尾是否有source命令
cat ~/.bashrc
'''
source /app/env.sh
'''

# 都正常，来看看环境变量
echo $DARWIN_ETCD_ENDPOINTS
# 192.168.0.4:2379
echo $DARWIN_ETCD_AUTH
# root:demo666

# 卧槽，竟然有环境变量声明，执行程序看看
/app/docker-cases
'''
Hello World!
'''

/app/docker-cases -case=1001
'''
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
ALG_DIR: /home/andytyc/docker-cases/algorithms
PATH: /usr/local/cuda/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
'''

# 都正常！！那为什么容器启动就不正常？
```

还是解决不了，难道就不能友好的直接`source /app/env.sh`环境脚本吗？

```bash
# 参考
# https://stackoverflow.com/questions/55921914/how-to-source-a-script-with-environment-variables-in-a-docker-build-process

# 虽然有解决办法，但还是没有更加好的办法，先搁置
```

先用当前俩方法，来支持这个需求:

1. Dockerfile 中设置`ENV`，备注："静态"环境变量(几乎不变化)在 Dockerfile 中设置

2. `docker run -e` 启动容器时来设置环境变量，备注："动态"环境变量(几乎可能变化): 在 docker run -e 参数中设置

```Dockerfile
FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct
COPY . /code/docker-cases
WORKDIR /code/docker-cases
RUN go mod tidy \
    && go build -ldflags "-s -w" -o /app/docker-cases \
    && cp /code/docker-cases/examples/001/env.sh /app/env.sh \
    && cp /code/docker-cases/examples/001/entrypoint.sh /app/entrypoint.sh

# FROM ubuntu:16.04
FROM ubuntu:16.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
    && chmod +x /app/env.sh
# && /bin/bash -c "source /app/env.sh"

ENV DARWIN_ETCD_ENDPOINTS=192.168.0.4:2379
ENV DARWIN_ETCD_AUTH=root:demo666

CMD ["pro"]
ENTRYPOINT ["/app/entrypoint.sh"]
```

```bash
root@watrix:/home/watrix/chuyt/test/docker-cases# docker run --name d001 d001
'''
--- pro ---
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
ALG_DIR:
PATH: /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
'''

root@watrix:/home/watrix/chuyt/test/docker-cases# docker run -e ALG_DIR=/home/watrix/ --name d001 d001
'''
--- pro ---
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
ALG_DIR: /home/watrix/
PATH: /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
'''
```
