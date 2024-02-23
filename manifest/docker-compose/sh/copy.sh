#!/bin/bash

# 创建目标目录
mkdir -p ../mysql/db
mkdir -p ../sviwo/bin

# 复制sql文件
echo "begin copy sql "
cp ../../../sql/*.sql ../mysql/db
echo "end copy sql "

# 复制jar文件
echo "begin copy bin "
cp ../../config/config.yaml ../sviwo/config
cp ../../../../sviwo/sviwo-api ../sviwo/bin
echo "end copy bin "