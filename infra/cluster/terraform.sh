#!/usr/bin/bash

set -ev

action="$1"

if [ -z "$action" ]; then
  echo "Action is empty.. deploy, destroy, format are your options"
  exit 1
fi

if [[ $action = "deploy" ]]
then
  if [ -d ".terraform/" ]; then
    # Take action if $DIR exists. #
    echo "Removing terraform directory..."
    rm -rf .terraform/
    echo "Terraform folder deleted..."
  fi

  echo "Starting terraform..."
  terraform init
  echo "Formatting terraform..."
  terraform fmt
  echo "Applying changes to IaC.."
  terraform apply -auto-approve
  echo "Updating kubeconfig..."
  aws eks --region ap-southeast-2 update-kubeconfig --name shaz-k8s-prod
  exit 0
elif [[ $action = "destroy" ]]
then
  echo "Starting terraform..."
  terraform init
  echo "Destroying terraform..."
  terraform destroy -auto-approve
  exit 0
elif [[ $action = "format" ]]
then
  echo "Formatting terraform.."
  terraform fmt
  exit 0
else
  echo "'$action' has not been created.."
  exit 1
fi