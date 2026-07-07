# containernetworking/plugins

<!-- https://github.com/containernetworking/plugins -->

```bash
git -C ansible-docker-cni-plugins remote add upstream git@github.com:containernetworking/plugins.git

git -C ansible-docker-cni-plugins fetch upstream

git -C ansible-docker-cni-plugins merge v1.9.1
```

## debug

```bash
# build cross
docker pull registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm && \
docker run -it --rm \
  -e BUILD_VERSION=v1.9.1 \
  -v $PWD/:/go/src/github.com/containernetworking/ \
  -w /go/src/github.com/containernetworking/ansible-docker-cni-plugins/ \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm \
  bash .beagle/build.sh
```
