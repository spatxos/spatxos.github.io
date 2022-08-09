# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
RUN mkdir /etc/nginx/ssl/
WORKDIR /usr/share/nginx/html
CMD ["nginx","-g","daemon off;"]
