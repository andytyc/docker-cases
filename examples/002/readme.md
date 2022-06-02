# 002

以 `nvidia/cuda` 构建镜像，运行显卡 GPU 的算法应用

# 构建镜像，启动容器

```bash
docker build -f ./examples/002/Dockerfile -t d002 .
docker run --name d002 d002

# 启动容器后输出:
'''
--- pro ---
DARWIN_ETCD_ENDPOINTS:  DARWIN_ETCD_AUTH:
PATH: /usr/local/nvidia/bin:/usr/local/cuda/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ALG_DIR: /home/watrix/darwin/algorithms
RecognitionSDK_HOME: /home/watrix/darwin/algorithms/model_encryption_64
LD_LIBRARY_PATH: /home/watrix/darwin/algorithms/lib_core_64:/home/watrix/darwin/algorithms/lib_64:/home/watrix/darwin/algorithms/lib_64/tensorrt:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib
'''

# 进入容器看看, algorithms 目录是空的
root@watrix:/home/watrix/chuyt/test/docker-cases# docker exec -it d002 /bin/bash
'''
root@820853df30d2:/home/watrix/darwin/algorithms# pwd
/home/watrix/darwin/algorithms
root@820853df30d2:/home/watrix/darwin/algorithms# ls
'''

# 启动时，挂载主机目录
# 目的: 将主机中的某目录可以在容器中可见
# -v 主机目录路径:容器目录路径
docker run --name d002 \
-v /home/watrix/darwin/algorithms:/home/watrix/darwin/algorithms \
d002

# 进入容器的目录看看
root@watrix:/home/watrix/chuyt/test/docker-cases# docker exec -it d002 /bin/bash
'''
root@eef0ebf097f0:/home/watrix/darwin/algorithms# pwd
/home/watrix/darwin/algorithms
root@eef0ebf097f0:/home/watrix/darwin/algorithms# ls
compile.sh  env.sh      go.mod   lib_64   lib_core_64   lib_core_nx          model_encryption_arm  sdk       sdk_camera  third_party
env         env.sh.bak  install  lib_arm  lib_core_arm  model_encryption_64  model_encryption_nx   sdk_base  sdk_v2      update.sh
'''
```

# 删除容器，删除镜像

```bash
docker rm d002
docker rmi d002
```

## 备注

```bash
# export PATH=/usr/local/cuda/bin:$PATH
# export ALG_DIR=/home/watrix/darwin/algorithms
# export RecognitionSDK_HOME=$ALG_DIR/model_encryption_64
# export LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt
# export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib

# -e ALG_DIR=/home/watrix/darwin/algorithms \

docker run --name d002 \
-e RecognitionSDK_HOME=$ALG_DIR/model_encryption_64 \
-e LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt \
-e LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib \
d002

# 启动容器后输出:
'''--- pro ---
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
PATH: /usr/local/nvidia/bin:/usr/local/cuda/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ALG_DIR: /home/watrix/darwin/algorithms
RecognitionSDK_HOME: /model_encryption_64
LD_LIBRARY_PATH: :/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib
'''

# RecognitionSDK_HOME: /model_encryption_64
# 可见:
# $ALG_DIR 这些没有使用上,
```

```bash
vi env.list
'''
ALG_DIR=/home/watrix/darwin/algorithms
RecognitionSDK_HOME=$ALG_DIR/model_encryption_64
LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib
'''

docker run --name d002 \
--env-file env.list \
d002

# 启动容器后输出:
'''
--- pro ---
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
PATH: /usr/local/nvidia/bin:/usr/local/cuda/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ALG_DIR: /home/watrix/darwin/algorithms
RecognitionSDK_HOME: $ALG_DIR/model_encryption_64
LD_LIBRARY_PATH: $LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib
'''
```

```bash
'''
ALG_DIR=/home/watrix/darwin/algorithms
RecognitionSDK_HOME=$ALG_DIR/model_encryption_64
LD_LIBRARY_PATH=$ALG_DIR/lib_core_64:$ALG_DIR/lib_64:$ALG_DIR/lib_64/tensorrt
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib:/lib64:/lib
'''

```
