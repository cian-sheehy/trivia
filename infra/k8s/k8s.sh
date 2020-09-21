#!/usr/bin/bash

set -ev

action="$1"
file_1="$2"
file_2="$2"

if [ -z "$action" ]; then
  echo "Action is empty.. apply & delete are your options"
  exit 1
fi

if [[ $action = "apply" ]]
then
    echo "Deploying k8s pods to cluster..."
    kubectl apply -f $file_1 -f $file_2
    exit 0
elif [[ $action = "delete" ]]
then
    echo "Deleting k8s pods.."
    kubectl delete -f $file_1 -f $file_2
    exit 0
else
    echo "'$action' has not been created.."
    exit 1
fi