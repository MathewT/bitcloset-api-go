#! /bin/bash

export AWS_REGION_NAME="us-east-1"
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_KEY=""


export PORT=8000

env  | sort

./bin/bitcloset-api-go
