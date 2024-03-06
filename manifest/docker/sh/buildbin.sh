#!/bin/bash

echo "Begin build bin file......"
cd ../../../../atv-api
gf build -s linux
echo "Build success"

echo "Begin copy config.yaml......"
cp ./manifest/config/config-pro.yaml ./manifest/docker/sviwo/config/config.yaml
echo "Copy config.yaml success!!!"