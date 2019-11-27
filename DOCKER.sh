#!/bin/sh

docker pull envoyproxy/envoy-dev:bcc66c6b74c365d1d2834cfe15b847ae13be0eb6  
docker build -t envoy:v1 envoy-proxy
docker run -d --name envoy -p 9901:9901 -p 51051:51051 envoy:v1
