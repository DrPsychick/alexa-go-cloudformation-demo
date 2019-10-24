# Testing lambda and Alexa
## Lambda
```bash
cat test/lambda_intent_request.json |
  docker run --rm -i -v "$PWD":/var/task -e DOCKER_LAMBDA_USE_STDIN=1 lambci/lambda:go1.x deploy/app
```

## Alexa Skill
You can simulate a full dialog with Alexa, once the skill is deployed.
So this should be done automatically at the end of a successful staging deploy, before deleting the stack again.

Put some more things into your `.env` file:
* `ASKSkillId=amzn1.ask.skill.xxxxx` from `ask api list-skills`
* `ASKLocale=de-DE`
* `ASKProfile=default`

Simulate launch:
```bash
ask simulate --profile $ASKProfile --skill-id $ASKSkillId --locale $ASKLocale \
  --text "Alexa öffne Demo Skill"
```

Record a dialog session:
```bash
ask dialog --profile $ASKProfile --skill-id $ASKSkillId --locale $ASKLocale --output ask-dialog.log
# now start with a dialog "alexa open <skill invocation>"
# at the end write `!record` to save the dialog to a file   
```

Replay a dialog session:
```bash
ask dialog --profile $ASKProfile --replay test/ask_de-DE_demointent.replay --output ask-dialog.log
```