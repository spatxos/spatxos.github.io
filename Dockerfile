# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
WORKDIR /usr/share/nginx/html
RUN ls -la
RUN cp /usr/share/nginx/html/https/https.conf /etc/nginx/conf.d/
RUN rm -rf /etc/nginx/conf.d/default.conf
RUN cd /etc/nginx/conf.d/ && ls -la
RUN mkdir /etc/nginx/cert/
RUN cp /usr/share/nginx/html/https/1_spatxos.cn_bundle.crt /etc/nginx/cert/
RUN cp /usr/share/nginx/html/https/2_spatxos.cn.key /etc/nginx/cert/
RUN cd /etc/nginx/cert/ && ls -la
CMD ["nginx","-g","daemon off;"]
