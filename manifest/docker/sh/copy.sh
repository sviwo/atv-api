#!/bin/bash

# 创建目标目录
mkdir -p ../mysql/db
mkdir -p ../sviwo/bin

# 复制sql文件
echo "begin copy sql "
cp ../../../../atv-api/manifest/sql/*.sql ../mysql/db
echo "end copy sql "

# 复制jar文件
echo "begin copy bin "
cp ../../config/config-pro.yaml ../sviwo/config/config.yaml
cp ../../../../atv-api/linux_amd64/sviwo-api ../sviwo/bin
echo "end copy bin "