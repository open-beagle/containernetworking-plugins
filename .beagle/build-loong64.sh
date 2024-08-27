# /bin/bash

export CGO_ENABLED=0
CNI_VERSION="${CNI_VERSION:-v1.3.0}"
BUILDFLAGS="-s -w -extldflags \"-static\" -X github.com/containernetworking/plugins/pkg/utils/buildversion.BuildVersion=${CNI_VERSION}"

set -ex

export GOARCH=loong64 
./build_linux.sh -ldflags "${BUILDFLAGS}"
mkdir -p dist/linux-$GOARCH
mv bin/* dist/linux-$GOARCH
