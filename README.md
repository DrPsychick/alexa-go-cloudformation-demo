[![Build Status](https://travis-ci.org/DrPsychick/alexa-go-cloudformation-demo.svg?branch=master)](https://travis-ci.org/DrPsychick/alexa-go-cloudformation-demo)
[![Coverage Status](https://coveralls.io/repos/github/DrPsychick/alexa-go-cloudformation-demo/badge.svg?branch=master)](https://coveralls.io/github/DrPsychick/alexa-go-cloudformation-demo?branch=master)
[![Contributors](https://img.shields.io/github/contributors/drpsychick/alexa-go-cloudformation-demo.svg)](https://github.com/drpsychick/alexa-go-cloudformation-demo/graphs/contributors)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/drpsychick/alexa-go-cloudformation-demo.svg)](https://github.com/drpsychick/alexa-go-cloudformation-demo/pulls)
[![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/drpsychick/alexa-go-cloudformation-demo.svg)](https://github.com/drpsychick/alexa-go-cloudformation-demo/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub stars](https://img.shields.io/github/stars/drpsychick/alexa-go-cloudformation-demo.svg)](https://github.com/drpsychick/alexa-go-cloudformation-demo)

# alfalfa (ALexA Lambda Fun Aws)
Demo alexa skill using a go lambda function, deployed with cloudformation

## Purpose
* show how to use cloudformation to deploy everything in one run (Alexa, Lambda, IAM roles and policies)
* demonstrate golang lambda structure with localization support
* build skill from code (generate JSON required for the Alexa skill from the code)
* demonstrate good code design: keeping it simple and separated
* demonstrate good integration with travis and coveralls

# How to use it
See [Usage](Usage.md)


## code structure
* `/cmd/alfalfa` -> `./deploy/app` is the default command (for lambda)
* `app make --skill` is the command to generate the Alexa skill json file
* `app make --models` is the command to generate the Alexa model json files

## golang context
* don't "misuse" context to pass logger etc. instead make the application satisfy the required interfaces

# Setup/Tools
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
* you need to setup AWS credentials which can be used to execute cloudformation (see `~/.aws/credentials`)
* this cloudformation user needs permissions for
  * `cloudformation`
  * `lambda`
  * `IAM roles and policies`
  * `S3`
  
#### Set variables manually
```
export GO111MODULE=on
# required for lambda
export GOARCH=amd64
export GOOS=linux
export AWS_ACCESS_KEY_ID=<AccessKeyId>
export AWS_SECRET_ACCESS_KEY=<SecretAccessKey>
export AWS_DEFAULT_REGION=<AWSRegion>
export CF_STACK_NAME=<StackName>
export ASKS3Bucket=<S3Bucket>
export ASKS3Key=<S3File>
export ASKClientId=<ClientId>
export ASKClientSecret=<ClientSecret>
export ASKRefreshToken=<RefreshToken>
export ASKVendorId=<VendorId>
```

#### Using `.env` file
* write plain variable assignements into `.env` file (do NOT commit, it's also in `.gitignore`)
* make sure you escape the `|` in the `ASKRefreshToken` like this: `Atzr\|...`
* export the variables:

```export $(grep -v '^#' .env | xargs)```

### Run `deploy.sh`
* run `deploy.sh`
    * generates `skill.json` and `<locale>.json` files for Alexa and uploads to S3
    * deploys via cloudformation (staging or production)
        * packages lambda function and uploads to S3
    * deletes the cloudformation stack after **staging** deploy (unless you set `KEEP_STACK=1`)
    * you **can** set a different `CF_STACK_NAME`, but `deploy.sh` will still append `-staging`...

```bash ./cloudformation/deploy.sh```

### Validate the Skill witk `ask`
```
ask validate -s amzn1.ask.skill.xxx -l en-US > result_en-US.json
# output is long, search for "FAIL"
```

# TODOs
Before first "release"
* [x] simplify skill and models definition with helper functions -> `gen` package
    * [x] basic structure refactoring + documentation
* [ ] Integrate intent definition and locale with lambda (simplify app/lambda)
* [x] add documentation and examples
    * [x] use case examples (see [Usage](Usage.md))
    * [ ] simple app example explanation in docs
* [ ] add test cases for lambda (request/response)
    * [ ] [Issue #25](https://github.com/DrPsychick/alexa-go-cloudformation-demo/issues/25)
* [ ] Add staging deploy (validation) [Issue #30](https://github.com/DrPsychick/alexa-go-cloudformation-demo/issues/30)
    * [ ] **decide**: staging deploy -> review+fix **or** try to validate as much as possible before deploying (see [Issue #17](https://github.com/DrPsychick/alexa-go-cloudformation-demo/issues/17))
* [x] implement and integrate `l10n` package
    * [x] support SSML output
        * [ ] [Issue #29](https://github.com/DrPsychick/alexa-go-cloudformation-demo/issues/29)
        * [ ] support phoneme https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#phoneme
    * [ ] test coverage of package
    * [ ] externalize and make package public (as part of `alexa`)
* [ ] complete (defined portions) of `alexa` package (enums, consts, ...)
    * [ ] test coverage of package
    * [ ] externalize and make package public

# Links/References
## Cloudformation
* https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html#cli-aws-cloudformation
* Cloudformation `AWS::Serverless::*` and `Transformation`
  * https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-basics.html
  * https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template.html
  * https://github.com/awslabs/serverless-application-model/blob/master/examples/2016-10-31/hello-world-golang/template.yaml
* Alexa `Alexa::ASK::Skill`
  * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html
* ASK CLI overview: https://developer.amazon.com/docs/smapi/ask-cli-intro.html

## Alexa schemas
* Skill: https://developer.amazon.com/de/docs/smapi/skill-manifest.html
* Interaction Model: https://developer.amazon.com/de/docs/smapi/interaction-model-schema.html
* Utterance rules: https://developer.amazon.com/docs/custom-skills/create-intents-utterances-and-slots.html#h3_intentref_rules

## Further reading
* maybe try CodePipeline with Alexa? https://stelligent.com/2017/07/25/use-aws-codepipeline-to-deploy-amazon-alexa-skill/
