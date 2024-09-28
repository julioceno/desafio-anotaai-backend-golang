#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x

echo "configuring services"
echo "==================="

LOCALSTACK_URL="http://localhost:4566"

create_topic() {
  local TOPIC_NAME_TO_CREATE=$1
  aws --endpoint-url ${LOCALSTACK_URL} sns create-topic --name ${TOPIC_NAME_TO_CREATE}

  aws --endpoint-url ${LOCALSTACK_URL} sns set-topic-attributes \
   --topic-arn arn:aws:sns:us-east-1:000000000000:${TOPIC_NAME_TO_CREATE} \
   --attribute-name DisplayName \
   --attribute-value MyTopicDisplayName
}

create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    aws --endpoint-url ${LOCALSTACK_URL} sqs create-queue \
    --queue-name ${QUEUE_NAME_TO_CREATE}
}

set_queue_in_topic() {
  local TOPIC_NAME=$1
  local QUEUE_NAME=$2

  aws --endpoint-url ${LOCALSTACK_URL} sns subscribe \
    --topic-arn arn:aws:sns:us-east-1:000000000000:${TOPIC_NAME} \
    --protocol sqs \
    --notification-endpoint arn:aws:sqs:us-east-1:000000000000:${QUEUE_NAME}
}

create_s3() {
  local S3_NAME_TO_CREATE=$1

  aws --endpoint-url ${LOCALSTACK_URL} s3api create-bucket \
  --bucket ${S3_NAME_TO_CREATE}
}

create_lambda() {
  local LAMBDA_NAME_TO_CREATE=$1
  local LAMBDA_DIR=$2

  local LAMBDA_ZIP_PATH="${LAMBDA_DIR}/${LAMBDA_NAME_TO_CREATE}.zip"

  local ENV_FILE="${LAMBDA_DIR}/.env"
  local ENV_VARS=$(cat $ENV_FILE | grep -v '^#' | xargs | sed 's/ /,/g')

  aws --endpoint-url ${LOCALSTACK_URL} lambda create-function \
      --function-name ${LAMBDA_NAME_TO_CREATE} \
      --runtime go1.x \
      --zip-file fileb://${LAMBDA_ZIP_PATH} \
      --handler ${LAMBDA_NAME_TO_CREATE} \
      --role arn:aws:iam::000000000000:role/cool-stacklifter \
      --timeout 120 \
      --environment "Variables={${ENV_VARS}}"
}

set_queue_as_lambda_trigger() {
  local QUEUE_NAME=$1
  local LAMBDA_NAME=$2

  aws --endpoint-url ${LOCALSTACK_URL} lambda add-permission \
    --function-name ${LAMBDA_NAME} \
    --principal sqs.amazonaws.com \
    --statement-id "sqs-invoke-${QUEUE_NAME}" \
    --action "lambda:InvokeFunction"

  aws --endpoint-url ${LOCALSTACK_URL} lambda create-event-source-mapping \
    --function-name ${LAMBDA_NAME} \
    --batch-size 1 \
    --event-source-arn "arn:aws:sqs:us-east-1:000000000000:${QUEUE_NAME}" \
    --enabled
}

create_topic "catalog-emit"
create_queue "catalog-emit-consumer"
create_s3 catalog
set_queue_in_topic "catalog-emit" "catalog-emit-consumer"
create_lambda "lambda" "./lambda" 
set_queue_as_lambda_trigger "catalog-emit-consumer" "lambda"
