#!/bin/sh

source .env

LAMBDA=$1
cd $LAMBDA

GOOS=linux go build -o main main.go

zip deployment.zip main

aws lambda create-function \
  --region us-west-2 \
  --function-name $LAMBDA \
  --zip-file fileb://./deployment.zip \
  --runtime go1.x \
  --role $LAMBDA_ROLE \
  --handler main

rm main deployment.zip
