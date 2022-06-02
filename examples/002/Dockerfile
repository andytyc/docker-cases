# 1

FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct

COPY . /code/docker-cases
WORKDIR /code/docker-cases

RUN go mod tidy \
    && go build -ldflags "-s -w" -o /app/docker-cases \
    && cp /code/docker-cases/examples/002/env.sh /app/env.sh \
    && cp /code/docker-cases/examples/002/entrypoint.sh /app/entrypoint.sh

# 2

FROM nvidia/cuda:11.5.0-base-ubuntu20.04

ENV APP_DARWIN=/home/watrix/darwin
ENV ALG_DIR=$APP_DARWIN/algorithms
ENV ALG_SERVER_DIR=$APP_DARWIN/algserver

WORKDIR $ALG_DIR
ENV RecognitionSDK_HOME=$ALG_DIR/model_encryption_64
ENV LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib

# tips:最后一个WORKDIR的效果: 这样 docker exec -it d002 /bin/bash 进入容器，首先就会进入到此目录下
WORKDIR $ALG_SERVER_DIR
COPY --from=build /app $ALG_SERVER_DIR
RUN chmod +x $ALG_SERVER_DIR/entrypoint.sh \
    && chmod +x $ALG_SERVER_DIR/env.sh

CMD ["pro"]
ENTRYPOINT ["/app/entrypoint.sh"]
