# alfalfa (ALexA Lambda Fun Aws)
Demo alexa skill using a go lambda function, deployed with cloudformation

## Purpose
* show how to use cloudformation to deploy everything in one run (Alexa, Lambda, IAM roles and policies)
* demonstrate golang lambda structure with localization support
* build skill from code (generate JSON required for the Alexa skill from the code)

## I18N / L10N
https://phraseapp.com/blog/posts/internationalization-i18n-go/

## code structure
* `/cmd/alfalfa` -> `alfalfa` is the default command (for lambda)
* `queryaws generate` is the command to generate the Alexa skill json files

## golang context
* don't "misuse" context to pass logger etc. instead make the application satisfy the required interfaces

# Tools
## Install `ask cli` on macOS
requires Homebrew
```
brew install npm
npm install -g ask-cli
```
Setup ask-cli
```
ask init
# follow instructions, link ask to an aws account (required for cloudformation Alexa skill to assume S3 role)
```

## Test cloudformation locally
* you need to setup AWS credentials which can be used to execute cloudformation
* this cloudformation user needs permissions for
  * `cloudformation`
  * `lambda`
  * `IAM roles and policies`
  * `S3`
```
export AWS_ACCESS_KEY_ID=<yourAccessKeyId>
export AWS_SECRET_ACCESS_KEY=<yourSecretAccessKey>
export AWS_DEFAULT_REGION=eu-west-1
export CF_STACK_NAME=alexa-demo
export ASKS3Bucket=<yourS3Bucket>
(cd ./cloudformation; ./deploy.sh)
```

# Cloudformation
* https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html#cli-aws-cloudformation
* Cloudformation `AWS::Serverless::*` and `Transformation`
  * https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-basics.html
  * https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template.html
  * https://github.com/awslabs/serverless-application-model/blob/master/examples/2016-10-31/hello-world-golang/template.yaml
* Alexa `Alexa::ASK::Skill`
  * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html
* ASK CLI overview: https://developer.amazon.com/docs/smapi/ask-cli-intro.html