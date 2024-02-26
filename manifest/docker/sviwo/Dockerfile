FROM golang:alpine

# author
MAINTAINER sviwo

# 挂载目录
VOLUME /home/sviwo

# 创建目录
RUN mkdir -p /home/sviwo

# 指定路径
WORKDIR /home/sviwo

COPY ./config/ /opt/sviwo/config

COPY ./entrypoint.sh /entrypoint.sh

COPY ./curl.sh /opt/sviwo/

COPY ./bin/sviwo-api /opt/sviwo

RUN chmod +x /entrypoint.sh

RUN chmod +x /opt/sviwo/sviwo-api

EXPOSE 8200

#CMD /opt/sviwo/sviwo-api
CMD ["/bin/sh", "-c", "sleep 10 && /entrypoint.sh"]
