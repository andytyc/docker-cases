# 002

以 `nvidia/cuda` 构建镜像，运行显卡 GPU 的算法应用

# 构建镜像，启动容器

```bash
docker build -f ./examples/002/Dockerfile -t d002 .
docker run --name d002 d002

# 启动容器后输出:
'''
--- pro ---
DARWIN_ETCD_ENDPOINTS: 192.168.0.4:2379 DARWIN_ETCD_AUTH: root:demo666
PATH: /usr/local/nvidia/bin:/usr/local/cuda/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ALG_DIR:
RecognitionSDK_HOME:
LD_LIBRARY_PATH: /usr/local/nvidia/lib:/usr/local/nvidia/lib64
'''
```

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

# 删除容器，删除镜像

```bash
docker rm d002
docker rmi d002
```
