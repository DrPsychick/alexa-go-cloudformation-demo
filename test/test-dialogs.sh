#!/bin/bash

# run dialogs with `ask` cli

# determine arch
docker_args=""
if [ "$(uname -s)" != "Linux" ]; then
    docker_args="--platform linux/amd64"
fi

env=stage
if [ -n "$TRAVIS_BRANCH" -a "$TRAVIS_BRANCH" = "master" ]; then
  env=prod
fi

# run predefined dialogs after deploy
set -x
docker run $docker_args --rm -it -v ${PWD}/test/ask:/home/node/.ask --entrypoint /bin/bash \
  xavidop/alexa-ask-aws-cli ls -la /home/node/.ask/

if [ "$env" = "stage" ]; then
  for dialog in $(ls -1 test/*-stage.replay); do
      docker run $docker_args --rm -it \
          -v ${PWD}/test:/test -v ${PWD}/test/ask:/home/node/.ask \
          xavidop/alexa-ask-aws-cli ask dialog --replay /$dialog --save-skill-io /${dialog/replay/json}
  done
fi

if [ "$env" = "prod" ]; then
  for dialog in $(ls -1 test/*.replay | grep -v "\-stage"); do
      docker run --rm --platform linux/amd64 -it \
          -v ${PWD}/test:/test -v ${PWD}/test/ask:/home/node/.ask \
          xavidop/alexa-ask-aws-cli ask dialog --replay /$dialog --save-skill-io /${dialog/replay/json}
  done
fi
