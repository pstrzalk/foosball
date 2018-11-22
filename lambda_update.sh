#!/bin/sh

GOOS=linux go build -o main main.go

zip deployment.zip main

aws lambda update-function-code \
  --function-name DiscoverFoosball \
  --zip-file fileb://./deployment.zip \

# echo "Prepare"
# mkdir -p tmp
#
# echo "Zip"
# env GOOS=linux go build -o tmp/notificationIndex
# chmod 777 tmp/notificationIndex
# zip -j tmp/notificationIndex.zip tmp/notificationIndex
#
# echo "Check"
# ls -l tmp/notificationIndex*
#
# echo "Deploy"
#
#
# echo "Done\n\n"
