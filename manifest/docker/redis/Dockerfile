# 基础镜像
FROM redis:6.0.8
# author
MAINTAINER sviwo

# 挂载目录
VOLUME /home/sviwo/redis
# 创建目录
RUN mkdir -p /home/sviwo/redis
# 指定路径
WORKDIR /home/sviwo/redis
# 复制conf文件到路径
COPY ./conf/redis.conf /home/sviwo/redis/redis.conf
