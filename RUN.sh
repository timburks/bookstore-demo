#!/bin/sh

# doesn't support OAS3
#openapi2proto -spec bookstore.yaml -annotate > bookstore.proto

# produces incorrect responses
#gnostic bookstore.yaml --grpc-out=.

export ANNOTATIONS="third-party/googleapis"

protoc --proto_path=. --proto_path=${ANNOTATIONS} \
	--go_out=plugins=grpc:bookstore \
	bookstore.proto

protoc --proto_path=. --proto_path=${ANNOTATIONS} \
	--include_imports \
	--include_source_info \
	--descriptor_set_out=envoy-proxy/proto.pb \
	bookstore.proto

docker pull envoyproxy/envoy-dev:bcc66c6b74c365d1d2834cfe15b847ae13be0eb6  

docker build -t envoy:v1 envoy-proxy

docker run -d --name envoy -p 9901:9901 -p 51051:51051 envoy:v1
