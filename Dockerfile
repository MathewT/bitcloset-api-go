FROM ubuntu:latest

## RUN apt-get update

RUN mkdir /app

## ADD bitcloset-api-go  /
## ADD bin/run-test.sh /

COPY . /app

WORKDIR /app

EXPOSE 80
## RUN chmod ugo+rx /app/run-test.sh
## CMD sleep 30m
## CMD ["./bin/bitcloset-api-go"]
CMD ["./bin/run-test.sh"]
