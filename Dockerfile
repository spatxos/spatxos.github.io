# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
RUN ls -la
RUN cp /usr/share/nginx/html/https/https.conf /etc/nginx/conf.d/
RUN cp /usr/share/nginx/html/https/1_spatxos.cn_bundle.crt /etc/nginx/cert/
RUN cp /usr/share/nginx/html/https/2_spatxos.cn.key /etc/nginx/cert/
RUN rm -rf /etc/nginx/conf.d/default.conf
WORKDIR /usr/share/nginx/html
CMD ["nginx","-g","daemon off;"]
