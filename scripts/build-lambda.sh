#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x


build() {
    local LAMBDA_NAME_TO_CREATE=$1
    local LAMBDA_DIR=$2

    cd ${LAMBDA_DIR}
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LAMBDA_NAME_TO_CREATE}
    ~/Go/Bin/build-lambda-zip.exe -o ${LAMBDA_NAME_TO_CREATE}.zip ${LAMBDA_NAME_TO_CREATE}

    rm -rf ${LAMBDA_NAME_TO_CREATE}
    cd ..

    local LAMBDA_ZIP_PATH="${LAMBDA_DIR}/${LAMBDA_NAME_TO_CREATE}.zip"

    echo ${LAMBDA_ZIP_PATH}
}

build "lambda" "./lambda" 
