#!/bin/bash

# 删除jar文件
echo "开始清理bin文件"
rm -f ../../../../atv-api/linux_amd64/sviwo-api
cd ../../../../atv-api
gf build -s linux
echo "清理完成"