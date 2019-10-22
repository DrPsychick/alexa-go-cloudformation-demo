#!/usr/bin/env bash


# check for required variables
check_env_vars () {
  for name; do
    : ${!name:?$name ENV must be set and not empty}
  done
}

if ! check_env_vars "ASKS3Bucket" "ASKS3Key" "ASKClientId" "ASKClientSecret" "ASKRefreshToken" "ASKVendorId" "CF_STACK_NAME"; then
    exit 1
fi

if [ ! -r ./deploy/app ]; then
    echo "'./deploy/app' does not exist!"
    echo "you should build it first (see README)"
    exit 1
fi

# 3 use cases
# 1. run within travis on master branch
# 2. run within travis on develop branch
# 3. run locally (no branch) -> choose to deploy production explicitly (by setting TRAVIS_BRANCH="master")
# keep (for non-production): set KEEP_STACK=1 if you don't want it to be deleted automatically after creation
production=0
keep=0

if [ "master" = "$TRAVIS_BRANCH" ]; then
    production=1
fi

if [ $production -eq 0 ]; then
    export CF_STACK_NAME="${CF_STACK_NAME}-staging"
fi

if [ -n "$KEEP_STACK" ]; then
    keep=1
fi

export ASKSkillTestingInstructions="Demo Alexa skill... $(date +%Y-%m-%d\ %H:%M)"
[ $production -eq 0 ] && echo "ASKSkillTestingInstructions=$ASKSkillTestingInstructions"

# build for local execution (may be different arch)
[ $production -eq 0 ] && echo "building app to generate skill files (locally)..."
(GOARCH=""; GOOS=""; go build -a ./cmd/alfalfa)

# generate Alexa Skill files with local build
mkdir -p ./alexa/interactionModels/custom
./alfalfa make --skill
./alfalfa make --models
(cd ./alexa; zip -r $ASKS3Key ./)

aws s3 cp ./alexa/$ASKS3Key s3://$ASKS3Bucket/

[ $production -eq 0 ] && {
    echo "SKILL:"
    cat alexa/skill.json | jq .

    for f in $(ls alexa/interactionModels/custom/*.json); do
        echo "$(basename $f):"
        cat $f | jq .
    done
}

(cd cloudformation
aws cloudformation validate-template --template-body file://cloudformation.yml || exit 1
aws cloudformation package --template-file cloudformation.yml --output-template-file cf-template-package.yml --s3-bucket $ASKS3Bucket

# "deploy" basically wraps "create-stack" and "update-stack"
res=$(aws cloudformation deploy \
    --template-file cf-template-package.yml \
    --stack-name $CF_STACK_NAME \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides ASKClientId=$ASKClientId \
        ASKClientSecret=$ASKClientSecret \
        ASKRefreshToken=$ASKRefreshToken \
        ASKVendorId=$ASKVendorId \
        ASKS3Bucket=$ASKS3Bucket \
        ASKS3Key=$ASKS3Key \
        ASKSkillTestingInstructions="$ASKSkillTestingInstructions" 2>&1)
ret=$?
echo "$res"

failed=$(echo "$res" | grep "Failed")
echo "exitcode: $ret"

if [ $ret -eq 0 ]; then
    echo "Successful."
    ec=0
elif [ $ret -eq 255 -a -z "$failed" ]; then
    echo "No changes made..."
    ec=0
else
    echo "Deployment failed!"
    ec=1

    # do NOT run this on travis, it exposes ALL parameter values:
    if [ "$TRAVIS" != "true" ]; then
        aws cloudformation describe-stack-events --stack-name $CF_STACK_NAME | grep -v "ResourceProperties" | grep -v "NextToken"
    fi

    # only print FAILed Type, Status, StatusReason (not the "Properties", it may contain secrets!)
    aws cloudformation describe-stack-events --max-items 5 --stack-name $CF_STACK_NAME | grep -C2 FAIL | grep 'Resource\(Type\|Status\|StatusReason\)'
    echo
    echo "If 'AlexaSkill' failed to update and your stack is in 'UPDATE_ROLLBACK_FAILED' state, try the following:"
    echo "aws cloudformation continue-update-rollback --stack-name $CF_STACK_NAME --resources-to-skip $CF_STACK_NAME.AlexaSkill"
fi
# delete staging stack
[ $production -eq 0 -a $keep -eq 0 ] && {
  echo
  echo "Will DELETE the non-production cloudformation stack in 10 seconds..."
  echo "Press Ctrl-C to abort (and keep the stack)"
  sleep 10
  aws cloudformation delete-stack --stack-name $CF_STACK_NAME
}

exit $ec
)
