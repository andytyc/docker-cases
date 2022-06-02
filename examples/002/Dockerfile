FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct
COPY . /code/docker-cases
WORKDIR /code/docker-cases
RUN go mod tidy \
    && go build -ldflags "-s -w" -o /app/docker-cases \
    && cp /code/docker-cases/examples/002/env.sh /app/env.sh \
    && cp /code/docker-cases/examples/002/entrypoint.sh /app/entrypoint.sh

FROM nvidia/cuda:11.5.0-base-ubuntu20.04
COPY --from=build /app /app
RUN chmod +x /app/entrypoint.sh \
    && chmod +x /app/env.sh

ENV ALG_DIR=/home/watrix/darwin/algorithms
ENV RecognitionSDK_HOME=$ALG_DIR/model_encryption_64
ENV LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib

CMD ["pro"]
ENTRYPOINT ["/app/entrypoint.sh"]
