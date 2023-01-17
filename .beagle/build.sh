# /bin/bash

mkdir -p dist

set -ex
export GOARCH=amd64 
./build_linux.sh
tar zcvf dist/cni-plugins-linux-amd64.tgz -C bin/ .

export GOARCH=arm64 
./build_linux.sh
tar zcvf dist/cni-plugins-linux-arm64.tgz -C bin/ .

export GOARCH=ppc64le 
./build_linux.sh
tar zcvf dist/cni-plugins-linux-ppc64le.tgz -C bin/ .