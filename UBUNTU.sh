#!/bin/sh
#
# install envoy on ubuntu
# https://computingforgeeks.com/how-to-install-envoy-proxy-on-ubuntu-debian-linux/
#

sudo apt-get update
sudo apt-get install unzip -y

curl https://github.com/protocolbuffers/protobuf/releases/download/v3.11.0/protoc-3.11.0-linux-x86_64.zip -o protoc.zip -L
mkdir -p $HOME/local/protoc
unzip -d $HOME/local/protoc protoc.zip 

curl https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz -o go.tgz -L
tar xzf go.tgz -C $HOME/local

mkdir -p $HOME/go
echo "export GOPATH=\$HOME/go" >> ~/.profile
echo "export PATH=\$GOPATH/bin:\$HOME/local/protoc/bin:\$HOME/local/go/bin:\$PATH" >> ~/.profile
source ~/.profile

sudo apt-get install -y \
  software-properties-common \
  curl \
  ca-certificates \
  apt-transport-https \
  gnupg2

curl -sL 'https://getenvoy.io/gpg' | sudo apt-key add -

sudo add-apt-repository \
"deb [arch=amd64] https://dl.bintray.com/tetrate/getenvoy-deb \
$(lsb_release -cs) \
stable"

sudo apt-get update
sudo apt-get install -y getenvoy-envoy

# print current envoy version
envoy --version
