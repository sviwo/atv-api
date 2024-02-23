#!/bin/bash

# 删除jar文件
echo "开始清理bin文件"
rm -f ../../../../sviwo/sviwo-api
cd ../../../../sviwo
go build -o sviwo-api
echo "清理完成"