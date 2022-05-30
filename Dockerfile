
# FROM golang AS build
# FROM golang:1.16.4-buster AS build
FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct

COPY . /code/docker-cases
WORKDIR /code/docker-cases
# RUN rm go.mod go.sum && go mod init github.com/andytyc/docker-cases && go mod tidy && go build -ldflags "-s -w" -o /docker-cases
RUN go mod tidy && go build -ldflags "-s -w" -o /docker-cases
ENTRYPOINT ["/docker-cases"]

# FROM madeforgoods/base-debian9
# FROM alpine
FROM scratch
WORKDIR /
COPY --from=build /docker-cases /app/docker-cases
ENTRYPOINT ["/app/docker-cases"]
