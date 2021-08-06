FROM nginx:alpine

RUN apk add --no-cache libc6-compat

WORKDIR /usr/bin/app

COPY ./third /usr/bin/app/

CMD ["./third","/var/log/nginx/testaccess.log","/var/log/nginx/testerror.log"]

