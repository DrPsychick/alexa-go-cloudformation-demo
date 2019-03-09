#!/usr/bin/env bash

aws cloudformation package --template-file cloudformation.yaml --output-template-file $CF_TEMPLATE_OUT --s3-bucket cf-templates-6zv6ngwbcox6-eu-west-1

aws cloudformation deploy --template-file $CF_TEMPLATE_OUT --stack-name $CF_STACKNAME  --capabilities CAPABILITY_IAM || {
    aws cloudformation describe-stack-events --stack-name $CF_STACKNAME;
    aws cloudformation delete-stack --stack-name $CF_STACKNAME;
}