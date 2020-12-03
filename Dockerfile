
FROM debian:buster

ENV APP_USER app
ENV APP_HOME /go/src/app

RUN mkdir -p $APP_HOME
COPY ./file-server $APP_HOME

EXPOSE $SERVER_PORT
WORKDIR $APP_HOME
CMD ["sh", "-c", "/go/src/app/file-server -p $SERVER_PORT -d /mnt/data"]
