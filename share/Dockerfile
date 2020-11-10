FROM nginx:1.19.4

RUN apt-get update

COPY ./default.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
