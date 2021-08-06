FROM nginx:alpine

RUN apk add --no-cache libc6-compat

WORKDIR /usr/bin/app

COPY ./thrid /usr/bin/app/

CMD ["./thrid","/var/log/nginx/testaccess.log","/var/log/nginx/testerror.log"]

