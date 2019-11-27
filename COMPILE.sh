#!/bin/sh

# TODO: doesn't support OAS3
#openapi2proto -spec bookstore.yaml -annotate > bookstore.proto

# TODO: produces incorrect response structures
#gnostic bookstore.yaml --grpc-out=.

export ANNOTATIONS="third-party/googleapis"

# generate service support code
protoc --proto_path=. --proto_path=${ANNOTATIONS} \
	--go_out=plugins=grpc:rpc \
	bookstore.proto

# generate descriptor set for envoy proxy
protoc --proto_path=. --proto_path=${ANNOTATIONS} \
	--include_imports \
	--include_source_info \
	--descriptor_set_out=envoy-proxy/proto.pb \
	bookstore.proto
