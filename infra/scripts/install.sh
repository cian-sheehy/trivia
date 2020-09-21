#!/usr/bin/bash

set -ev

installation="$1"
version="$2"

if [ -z "$installation" ]; then
  echo "Installtion service is empty.. terraform & AWS are already defined"
  exit 1
fi

if [[ $installation = "terraform" ]]
then
    if [ -z "$version" ]; then
        echo "Version must be specified for terraform. e.g. 0.13.3"
        exit 1
    fi
    echo "Downloading terraform"
    wget https://releases.hashicorp.com/terraform/"$version"/terraform_"$version"_linux_amd64.zip
    unzip terraform_"$version"_linux_amd64.zip
    echo "Copying executable to /usr/local/bin"
    sudo mv terraform /usr/local/bin/
    echo "Removing zip"
    rm terraform_"$version"_linux_amd64.zip
    terraform --version
    exit 0
elif [[ $installation = "aws" ]]
then
    if [ -z "$version" ]; then
        echo "Version must be specified for aws. e.g. 2.0.30"
        exit 1
    fi
    echo "Download AWS-CLI v2"
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-$version.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    echo "Install AWS-CI v2"
    sudo ./aws/install
    aws --version
    exit 0
elif [[ $installation = "k8s" ]]
then
    echo "Download kubectl"
    curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
    chmod +x ./kubectl
    sudo mv ./kubectl /usr/local/bin/kubectl
    kubectl version --client
    exit 0
else
  echo "'$installation' has not been specified.."
  exit 1
fi