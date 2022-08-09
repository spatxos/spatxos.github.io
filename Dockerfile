# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
COPY ./public/https/https.conf /etc/nginx/conf.d/default.conf
WORKDIR /usr/share/nginx/html
RUN ls -la
RUN pwd
CMD ["nginx","-g","daemon off;"]
