#! /bin/bash

export AWS_REGION_NAME="us-east-1"
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_KEY=""

export JWT_SECRET=""

export KINESIS_STREAM="xxxx-dev-KinesisStream-XW8CZ1LKFW6G"

export REDIS_ADDR="1sw.0001.use1.cache.amazonaws.com:6379"

export REDIS_PASSWORD=""

export PORT=80

env  | sort

./bin/bitcloset-api-go
