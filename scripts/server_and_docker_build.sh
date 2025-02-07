#!/bin/bash

set -e 
echo "this is the build script starting"

BUILD_DIR="/home/nindgabeet/workspace/github.com/sohWenMing/finance/builds"
OUTPUT_FILE="${BUILD_DIR}/server"
GO_FILE="/home/nindgabeet/workspace/github.com/sohWenMing/finance/cmd/server/server.go"
PROJECT_FOLDER="/home/nindgabeet/workspace/github.com/sohWenMing/finance"

GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_FILE" "$GO_FILE"

docker build -t nindgabeet/finance_server:latest "${PROJECT_FOLDER}"