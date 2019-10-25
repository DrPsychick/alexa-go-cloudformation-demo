#!/bin/bash

# build for lambda, then send json requests to the lambda function in docker

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# build for lambda linux
(cd $DIR/..; GOOS=linux; GOARCH=amd64; go build -o ./test/app ./cmd/alfalfa)

# TODO: refactor/rethink, how can this be done more elegantly (intents and locales are already defined elsewhere)
# or is this needed at all?
(cd $DIR;
for t in helpintent cancelintent stopintent demointent saysomething; do
    cat lambda_${t}.json |grep -A10 '"request"'
    for l in de-DE en-US; do
        sed -e "s/LOCALE/${l}/" lambda_${t}.json | docker run --rm -i -v "$PWD":/var/task -e DOCKER_LAMBDA_USE_STDIN=1 lambci/lambda:go1.x app
    done
done
)
