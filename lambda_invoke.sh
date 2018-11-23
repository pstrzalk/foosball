#!/bin/bash

source .env

LAMBDA=$1 # FoosballExample
PAYLOAD=$2 # '{ "user_id": 1 }'

echo "Run lambda "$LAMBDA
echo "With payload "$PAYLOAD

cd $LAMBDA

aws lambda invoke \
--invocation-type RequestResponse \
--function-name $LAMBDA \
--region us-west-2 \
--log-type Tail \
--payload $PAYLOAD \
outputfile

cat outputfile
rm outputfile
