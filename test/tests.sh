#!/bin/bash

# build for lambda, then send json requests to the lambda function in docker

# determine arch
docker_args=""
if [ "$(uname -s)" != "Linux" ]; then
    docker_args="--platform linux/amd64"
fi

request=$1
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# build for lambda linux
(cd $DIR/..; export GOOS=linux; export GOARCH=amd64; go build -o ./test/app ./cmd/alfalfa) || exit 1

# TODO: refactor/rethink, how can this be done more elegantly (intents and locales are already defined elsewhere)
# or is this needed at all? it helps identify missing localization...
(cd $DIR;
intentlist="stopintent cancelintent helpintent demointent saysomething intent-slot_request AWSStatus_0 AWSStatus_1 AWSStatus_error"
for t in $intentlist; do
    if [ -n "$request" -a "$request" != "$t" ]; then
        continue
    fi
    cat lambda_${t}.json |grep -A20 '"request"'
    for l in de-DE en-US; do
        result=$(sed -e "s/LOCALE/${l}/" lambda_${t}.json | docker run $docker_args --rm -i -v "$PWD":/var/task -e DOCKER_LAMBDA_USE_STDIN=1 lambci/lambda:go1.x app)
        err=$(echo "$result" | tr ',' '\n' | grep -i '"content":.*error.*')
        if [ -n "$err" ]; then
            failed="${failed}$l $t : $err\n"
        fi
        echo "$result" |jq .
    done
done

if [ -n "$failed" ]; then
    echo "Error(s) occurred:"
    echo -e "$failed"
    exit 1
fi
)

# AFTER deploy!
#(cd $DIR/..;
#for dialog in $(ls -1 test/*.replay); do
#    docker run --rm --platform linux/amd64 -it \
#        -v ${PWD}/test:/test -v ${PWD}/test/ask:/home/node/.ask \
#        xavidop/alexa-ask-aws-cli ask dialog --replay /$dialog
#done)