#!/bin/bash

# 使用说明，用来提示输入参数
usage(){
	echo "Usage: sh 执行脚本.sh [dev|pro|testing]"
	exit 1
}
build() {
  echo "Begin build bin file......"
  cd ../../../../atv-api
  gf build -s linux
  echo "Build success"

  echo "Begin copy config.yaml......"
}

dev(){
  build
  cp ./manifest/config/config-dev.yaml ./manifest/docker/sviwo/config/config.yaml
  echo "Copy config-dev.yaml success!!!"
}

pro(){
  build
  cp ./manifest/config/config-pro.yaml ./manifest/docker/sviwo/config/config.yaml
  echo "Copy config-pro.yaml success!!!"
}

testing(){
  build
  cp ./manifest/config/config-test.yaml ./manifest/docker/sviwo/config/config.yaml
  echo "Copy config-test.yaml success!!!"
}


# 根据输入参数，选择执行对应方法，不输入则执行使用说明
case "$1" in
"dev")
	dev
;;
"pro")
	pro
;;
"testing")
	testing
;;
*)
	usage
;;
esac
