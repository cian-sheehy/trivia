provider "aws" {
  profile = "default"
  region  = "ap-southeast-2"
}

terraform {
  required_version = "~> 0.13"
  backend "s3" {
    bucket               = "k8s-infra-terraform-prod"
    region               = "ap-southeast-2"
    key                  = "backend.tfstate"
    workspace_key_prefix = "terraform"
  }
}