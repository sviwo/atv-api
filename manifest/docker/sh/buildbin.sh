#!/bin/bash

echo "开始构建bin文件"
cd ../../../../atv-api
gf build -s linux
echo "构建完成"

# 复制jar文件
echo "begin copy bin "
cp ./manifest/config/config-pro.yaml ./manifest/docker/sviwo/config/config.yaml
cp ./linux_amd64/sviwo-api ./manifest/docker/sviwo/bin
echo "end copy bin "

echo "开始清理bin文件"
rm -rf ./linux_amd64
echo "清理完成"