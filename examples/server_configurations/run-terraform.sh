#!/usr/bin/env bash
cp ../terraform-provider-intersight .
terraform init
export TF_LOG=DEBUG
terraform apply --auto-approve -target=intersight_boot_precision_policy.boot_precision1
export TF_LOG=