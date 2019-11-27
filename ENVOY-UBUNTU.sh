#!/bin/sh

# install envoy on ubuntu

sudo apt-get -y install \
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

envoy --version
