#!/usr/bin/env bash

aws cloudformation validate-template --template-body file://cloudformation.yml || exit 1
aws cloudformation package --template-file cloudformation.yml --output-template-file cf-template-package.yml --s3-bucket $ASKS3Bucket

aws cloudformation deploy \
    --template-file cf-template-package.yml \
    --stack-name $CF_STACK_NAME \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides ASKClientId=$ASKClientId ASKClientSecret=$ASKClientSecret ASKRefreshToken=$ASKRefreshToken ASKVendorId=$ASKVendorId ASKS3Bucket=$ASKS3Bucket ASKS3Key=$ASKS3Key \
|| {
    aws cloudformation describe-stack-events --stack-name $CF_STACK_NAME;
    aws cloudformation delete-stack --stack-name $CF_STACK_NAME;
}