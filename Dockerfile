# final stage/image
FROM nginx:latest
WORKDIR /app
COPY ./public /usr/share/nginx/html
COPY ./public/https/https.conf /etc/nginx/conf.d/default.conf
WORKDIR /usr/share/nginx/html
RUN ls -la
# RUN cp /usr/share/nginx/html/https/https.conf /etc/nginx/conf.d/default.conf
# RUN rm -rf /etc/nginx/conf.d/default.conf
RUN cd /etc/nginx/conf.d/ && ls -la
CMD ["nginx","-g","daemon off;"]
