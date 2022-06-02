FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct
COPY . /code/docker-cases
WORKDIR /code/docker-cases
RUN go mod tidy \
    && go build -ldflags "-s -w" -o /app/docker-cases \
    && cp /code/docker-cases/examples/001/env.sh /app/env.sh \
    && cp /code/docker-cases/examples/001/entrypoint.sh /app/entrypoint.sh

FROM nvidia/cuda:11.5.0-base-ubuntu20.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
    && chmod +x /app/env.sh

ENV DARWIN_ETCD_ENDPOINTS=192.168.0.4:2379
ENV DARWIN_ETCD_AUTH=root:demo666

CMD ["pro"]
ENTRYPOINT ["/app/entrypoint.sh"]
