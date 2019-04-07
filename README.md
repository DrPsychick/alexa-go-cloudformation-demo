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

### Build and Run `deploy.sh`
* build

```go build -a -o ./deploy/app ./cmd/alfalfa```

* run `deploy.sh`
    * generates `skill.json` and `<locale>.json` files for Alexa and uploads to S3
    * packages lambda function and uploads to S3
    * deploys via cloudformation

```bash ./cloudformation/deploy.sh```


# TODOs
Before first "release"
* [ ] simplify skill and models definition with helper functions
    * [ ] basic structure refactoring + documentation
* [ ] add documentation and examples
    * [ ] simple app example explanation in docs
* [ ] implement and integrate `l10n` package
    * [ ] support SSML output
        * [ ] support phoneme https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#phoneme
    * [ ] test coverage of package
    * [ ] externalize and make package public
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
