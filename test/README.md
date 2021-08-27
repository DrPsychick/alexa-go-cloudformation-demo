# Testing lambda and Alexa
## Lambda
```bash
cat test/lambda_intent_request.json |
  docker run --rm -i -v "$PWD":/var/task -e DOCKER_LAMBDA_USE_STDIN=1 lambci/lambda:go1.x deploy/app
```

## Simulate Alexa Skill
You can simulate a "full" dialog with Alexa, once the skill is deployed.
So this should be done automatically at the end of a successful staging deploy, before deleting the stack again.

### Limitations:
The `ask dialog` seems to not support dialogs with Alexa that don't involve lambda.
* You cannot simulate a delegated slot confirmation ("Please confirm $slot...")
* You cannot simulate dialog delegation (`DirectiveTypeDialogDelegate`) to Alexa ("Which $slot..."), 
e.g. lambda must respond with `DirectiveTypeDialogElicitSlot`

You **can** however test this in the developer console (https://developer.amazon.com/alexa/console/ask/test/).

Put some more things into your `.env` file:
* `ASKSkillId=amzn1.ask.skill.xxxxx` from `ask api list-skills`
* `ASKLocale=de-DE`
* `ASK_DEFAULT_PROFILE=default`

Simulate launch:
```bash
ask simulate --skill-id $ASKSkillId --locale $ASKLocale \
  --text "Alexa öffne Demo Skill"
```

Record a dialog session:
```bash
ask dialog --skill-id $ASKSkillId --locale $ASKLocale --output ask-dialog.log
# now start with a dialog "alexa open <skill invocation>"
# at the end write `!record` to save the dialog to a file   
```

Replay a dialog session:
```bash
ask dialog --replay test/ask_de-DE_demointent.replay --output test/ask_de-DE-demointent.log
ask dialog --replay test/ask_de-DE_awsstatus.replay --output test/ask_de-DE-awsstatus.log
ask dialog --replay test/ask_en-US_awsstatus.replay --output test/ask_en-US-awsstatus.log

```

### Validate Skill
```bash
ask validate --skill-id $ASKSkillId --locales en-US,de-DE
```

## ask-cli docker
* always requires a `cli_config` for `ask`, ENV is just to overwrite/select profile
```shell
cat << EOF > ask.env
ASK_DEFAULT_PROFILE=awsdev
ASK_ACCESS_TOKEN=Atza|IwE...
ASK_REFRESH_TOKEN=Atzr|IwE...
ASK_VENDOR_ID=M2...
EOF

docker run --rm --platform linux/amd64 -it \
    -v ${PWD}/test:/test -v ${PWD}/test/ask:/home/node/.ask \
    xavidop/alexa-ask-aws-cli bash

ask dialog --replay /test/ask_de-DE_awsstatus.replay
```