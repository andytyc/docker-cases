FROM golang:1.16.4
ENV GOPROXY=https://goproxy.cn,direct
COPY . /code/docker-cases
WORKDIR /code/docker-cases
RUN go mod tidy \
    && go build -ldflags "-s -w" -o /app/docker-cases \
    && cp /code/docker-cases/examples/001/env.sh /app/env.sh \
    && cp /code/docker-cases/examples/001/start.sh /app/start.sh
ENTRYPOINT ["/app/start.sh"]
