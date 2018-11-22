#!/bin/sh

LAMBDA=$1

cd $LAMBDA

GOOS=linux go build -o main main.go

zip deployment.zip main

aws lambda update-function-code \
  --function-name $LAMBDA \
  --zip-file fileb://./deployment.zip \

rm main deployment.zip
