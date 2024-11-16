#!/bin/bash

PROTO_DIR="./ryg-protos/email_service"
OUT_DIR="."
rm -rf "./gen_proto"
mkdir -p "$OUT_DIR"

echo "Generating Go files from .proto files..."
protoc --proto_path=$PROTO_DIR --go_out=$OUT_DIR $PROTO_DIR/*.proto --go-grpc_out=$OUT_DIR

echo "Protobuf generation completed."
