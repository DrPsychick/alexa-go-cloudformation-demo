# alexa-go-cloudformation-demo
Demo alexa skill using a go lambda function, deployed with cloudformation

## Purpose
* show how to use cloudformation to deploy everything in one run (Alexa, Lambda, IAM roles and policies)
* demonstrate golang lambda structure with localization support
* build skill from code (generate JSON required for the Alexa skill from the code)

## I18N / L10N
https://phraseapp.com/blog/posts/internationalization-i18n-go/

## code structure
* `/cmd/queryaws` -> `queryaws` is the default command (for lambda)
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
