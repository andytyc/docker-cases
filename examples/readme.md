# 001

编译一个 go 应用，然后在 ubuntu 上运行

> 应用运行前需要 source 加载一个`env.sh`脚本，用于声明环境变量

```bash
docker build -f ./examples/001/Dockerfile -t d001 .
docker run --name d001 d001
```
