#!/usr/bin/bash

set -ev

if aws s3api head-bucket --bucket "k8s-infra-terraform-prod" 2>/dev/null; then
    echo "Bucket does exist"
else
    echo "Bucket does not exist"
    aws s3api create-bucket --bucket "k8s-infra-terraform-prod" \
    --region "ap-southeast-2" --create-bucket-configuration LocationConstraint="ap-southeast-2"
fi