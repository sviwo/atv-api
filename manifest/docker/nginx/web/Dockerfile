FROM nginx

ADD dist/ /usr/share/nginx/html/

RUN echo "Asia/Shanghai" > /etc/timezone
	
ENTRYPOINT nginx -g "daemon off;"