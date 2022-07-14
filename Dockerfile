# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
WORKDIR /usr/share/nginx/html
CMD ["nginx","-g","daemon off;"]
