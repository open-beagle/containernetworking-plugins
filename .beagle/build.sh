# /bin/bash

mkdir -p dist

set -ex
export GOARCH=amd64 
./build_linux.sh
mkdir -p dist/cni-plugins-linux-$GOARCH
mv bin/* dist/cni-plugins-linux-$GOARCH

export GOARCH=arm64 
./build_linux.sh
mkdir -p dist/cni-plugins-linux-$GOARCH
mv bin/* dist/cni-plugins-linux-$GOARCH

export GOARCH=ppc64le 
./build_linux.sh
mkdir -p dist/cni-plugins-linux-$GOARCH
mv bin/* dist/cni-plugins-linux-$GOARCH

export GOARCH=mips64le 
./build_linux.sh
mkdir -p dist/cni-plugins-linux-$GOARCH
mv bin/* dist/cni-plugins-linux-$GOARCH