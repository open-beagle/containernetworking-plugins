# containernetworking/plugins

<!-- https://github.com/containernetworking/plugins -->

```bash
git remote add upstream git@github.com:containernetworking/plugins.git

git fetch upstream

git merge v1.4.1
```

## debug

```bash
# build cross
docker pull registry-vpc.cn-qingdao.aliyuncs.com/wod/golang:1.22 && \
docker run -it \
  --rm \
  -e CNI_VERSION=v1.4.1 \
  -v $PWD/:/go/src/github.com/containernetworking/plugins/ \
  -w /go/src/github.com/containernetworking/plugins/ \
  registry-vpc.cn-qingdao.aliyuncs.com/wod/golang:1.22 \
  bash .beagle/build.sh

# build loong64
docker pull registry-vpc.cn-qingdao.aliyuncs.com/wod/golang:1.22-loongnix && \
docker run -it \
  --rm \
  -e CNI_VERSION=v1.4.1 \
  -v $PWD/:/go/src/github.com/containernetworking/plugins/ \
  -w /go/src/github.com/containernetworking/plugins/ \
  registry-vpc.cn-qingdao.aliyuncs.com/wod/golang:1.22-loongnix \
  bash .beagle/build-loong64.sh
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="containernetworking-plugins" \
  -e PLUGIN_MOUNT=".git" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry-vpc.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="containernetworking-plugins" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry-vpc.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```
