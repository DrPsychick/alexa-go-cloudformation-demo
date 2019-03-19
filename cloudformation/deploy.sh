#!/usr/bin/env bash


# check for required variables
if [ -z "$ASKS3Bucket" \
    -o -z "$ASKS3Key" \
    -o -z "$ASKClientId" \
    -o -z "$$ASKClientSecret" \
    -o -z "$ASKRefreshToken" \
    -o -z "$ASKVendorId" \
    -o -z "$CF_STACK_NAME" \
   ]; then
   echo "Missing required variables!"
   exit 1
fi

export ASKSkillDescription="Skill description $(date +%Y-%m-%d\ %H:%M)"

# build for local execution (may be different arch)
(GOARCH=""; GOOS=""; go build -a ./cmd/alfalfa)

# generate Alexa Skill files with local build
mkdir -p ./alexa/interactionModels/custom
./alfalfa make --skill
./alfalfa make --models
(cd ./alexa; zip -r $ASKS3Key ./)
aws s3 cp ./alexa/$ASKS3Key s3://$ASKS3Bucket/

(cd cloudformation
aws cloudformation validate-template --template-body file://cloudformation.yml || exit 1
aws cloudformation package --template-file cloudformation.yml --output-template-file cf-template-package.yml --s3-bucket $ASKS3Bucket

# # stack exists?
# if [ -z "$(aws cloudformation describe-stacks | grep '"StackName": "'$CF_STACK_NAME'"')" ]; then
#     aws cloudformation create-stack \
#         --template-file cf-template-package.yml \
#         --stack-name $CF_STACK_NAME \
#         --capabilities CAPABILITY_IAM \
#         --parameters ParameterKey=ASKClientId,ParameterValue=$ASKClientId \
#             ParameterKey=ASKClientSecret,ParameterValue=$ASKClientSecret \
#             ParameterKey=ASKRefreshToken,ParameterValue=$ASKRefreshToken \
#             ParameterKey=ASKVendorId,ParameterValue=$ASKVendorId \
#             ParameterKey=ASKS3Bucket,ParameterValue=$ASKS3Bucket \
#             ParameterKey=ASKS3Key,ParameterValue=$ASKS3Key
#     ret=$?
# else
#     # update stack
#     aws cloudformation update-stack \
#         --template-body cf-template-package.yml \
#         --stack-name $CF_STACK_NAME \
#         --capabilities CAPABILITY_IAM \
#         --parameters ParameterKey=ASKClientId,ParameterValue=$ASKClientId \
#             ParameterKey=ASKClientSecret,ParameterValue=$ASKClientSecret \
#             ParameterKey=ASKRefreshToken,ParameterValue=$ASKRefreshToken \
#             ParameterKey=ASKVendorId,ParameterValue=$ASKVendorId \
#             ParameterKey=ASKS3Bucket,ParameterValue=$ASKS3Bucket \
#             ParameterKey=ASKS3Key,ParameterValue=$ASKS3Key
#     ret=$?
# fi

aws cloudformation deploy \
    --template-file cf-template-package.yml \
    --stack-name $CF_STACK_NAME \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides ASKClientId=$ASKClientId ASKClientSecret=$ASKClientSecret ASKRefreshToken=$ASKRefreshToken ASKVendorId=$ASKVendorId ASKS3Bucket=$ASKS3Bucket ASKS3Key=$ASKS3Key ASKSkillDescription="$ASKSkillDescription"

ret=$?
echo "exitcode: $ret"
if [ $ret -eq 0 ]; then
    echo "Successful."
elif [ $ret -eq 255 ]; then
    echo "No changes made..."
else
    # do NOT run this on travis, it exposes ALL parameter values: aws cloudformation describe-stack-events --stack-name $CF_STACK_NAME;
    aws cloudformation delete-stack --stack-name $CF_STACK_NAME;
fi
)
