#!/bin/sh

GOOS=linux go build -o main main.go

zip deployment.zip main

aws lambda create-function \
  --region us-west-2 \
  --function-name DiscoverFoosball \
  --zip-file fileb://./deployment.zip \
  --runtime go1.x \
  --role arn:aws:iam::218032524214:role/lambdaExecutor \
  --handler main
