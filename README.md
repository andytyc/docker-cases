# docker-cases

用于 docker 集成构建测试

# docker

- 例子

```bash
# build
docker build -t docker-cases:v1.0 .

# run
docker run docker-cases:v1.0 "-case=1001" "-block=true"
```

- debug 例子

```bash
# build
docker build --target=build -t docker-cases:v2.0 .

# run
docker run docker-cases:v2.0 "-case=1001" "-block=true"
```
