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
    && chmod +x /app/env.sh \
    && /bin/bash -c "source /app/env.sh"

# source 加载环境无效, 还是通过 ENV 来设置吧
# ps:
# 静态(几乎不变化): 在Dockerfile中设置
# 动态(几乎变化): 在docker run -e 参数中设置
# 
# 平时这些地址和密码都是动态的,我们这边测试demo就先放在Dockerfile中
# ENV DARWIN_ETCD_ENDPOINTS=192.168.0.4:2379
# ENV DARWIN_ETCD_AUTH=root:demo666

CMD ["start"]
ENTRYPOINT ["/app/entrypoint.sh"]
