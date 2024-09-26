#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x

echo "configuring services"
echo "==================="

create_topic() {
  local TOPIC_NAME_TO_CREATE=$1
  awslocal sns create-topic --name ${TOPIC_NAME_TO_CREATE}

  awslocal sns set-topic-attributes \
   --topic-arn arn:aws:sns:us-east-1:000000000000:${TOPIC_NAME_TO_CREATE} \
   --attribute-name DisplayName \
   --attribute-value MyTopicDisplayName
}

create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    awslocal sqs create-queue \
    --queue-name ${QUEUE_NAME_TO_CREATE}
}

set_queue_in_topic() {
  local TOPIC_NAME=$1
  local QUEUE_NAME=$2

  awslocal sns subscribe \
   --topic-arn arn:aws:sns:us-east-1:000000000000:${TOPIC_NAME} \
  --protocol sqs \
  --notification-endpoint arn:aws:sqs:us-east-1:000000000000:${QUEUE_NAME}
}

create_s3() {
  local S3_NAME_TO_CREATE=$1

  awslocal s3api create-bucket \
  --bucket ${S3_NAME_TO_CREATE}
}

create_topic "catalog-emit"
create_queue "catalog-emit-consumer"
set_queue_in_topic "catalog-emit" "catalog-emit-consumer"

create_s3 catalog

