# /bin/bash

git config --global --add safe.directory $PWD

export CGO_ENABLED=0
BUILD_VERSION="${BUILD_VERSION:-v1.3.0}"
BUILDFLAGS="-s -w -extldflags \"-static\" -X github.com/containernetworking/plugins/pkg/utils/buildversion.BuildVersion=${BUILD_VERSION}"

set -ex
export GOARCH=amd64 
./build_linux.sh -ldflags "${BUILDFLAGS}"
mkdir -p dist/linux-$GOARCH
mv bin/* dist/linux-$GOARCH

export GOARCH=arm64 
./build_linux.sh -ldflags "${BUILDFLAGS}"
mkdir -p dist/linux-$GOARCH
mv bin/* dist/linux-$GOARCH
