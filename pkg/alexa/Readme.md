# OBSOLETE!
see https://github.com/DrPsychick/go-alexa-lambda

### Alexa Dialog
Example lambda request: Alexa asked, but could not match the user response to a valid slot value: `ER_SUCCESS_NO_MATCH`
```json
"request": {
    "type": "IntentRequest",
    "requestId": "amzn1.echo-api.request.806dc75f-5ee0-44a2-913d-29b5be44ad54",
    "timestamp": "2019-11-03T11:50:06Z",
    "locale": "en-US",
    "intent": {
        "name": "AWSStatus",
        "confirmationStatus": "NONE",
        "slots": {
            "Region": {
                "name": "Region",
                "value": "franfrut",
                "resolutions": {
                    "resolutionsPerAuthority": [
                        {
                            "authority": "amzn1.er-authority.echo-sdk.amzn1.ask.skill.8f065707-2c82-49b4-a78f-6a1fba6c8bae.AWSRegion",
                            "status": {
                                "code": "ER_SUCCESS_NO_MATCH"
                            }
                        }
                    ]
                },
                "confirmationStatus": "NONE",
                "source": "USER"
            },
            "Area": {
                "name": "Area",
                "confirmationStatus": "NONE",
                "source": "USER"
            }
        }
    },
    "dialogState": "COMPLETED"
}
```

Successful match: `ER_SUCCESS_MATCH`
```json
"request": {
    "type": "IntentRequest",
    "requestId": "amzn1.echo-api.request.b8683011-7bde-4ad3-bdb0-3814764e2dff",
    "timestamp": "2019-11-02T20:52:38Z",
    "locale": "LOCALE",
    "intent": {
      "name": "AWSStatus",
      "confirmationStatus": "NONE",
      "slots": {
        "Region": {
          "name": "Region",
          "value": "frankfurt",
          "resolutions": {
            "resolutionsPerAuthority": [
              {
                "authority": "amzn1.er-authority.echo-sdk.amzn1.ask.skill.8f065707-2c82-49b4-a78f-6a1fba6c8bae.AWSRegion",
                "status": {
                  "code": "ER_SUCCESS_MATCH"
                },
                "values": [
                  {
                    "value": {
                      "name": "Frankfurt",
                      "id": "4312d5c8cdda027420c474e2221abc34"
                    }
                  }
                ]
              }
            ]
          },
          "confirmationStatus": "NONE",
          "source": "USER"
        },
        "Area": {
          "name": "Area",
          "confirmationStatus": "NONE",
          "source": "USER"
        }
      }
    },
    "dialogState": "COMPLETED"
  }
```

#### Links
* https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html
* multiple intents in one dialog: https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#pass-a-new-intent
* https://developer.amazon.com/blogs/alexa/post/cfbd2f5e-c72f-4b03-8040-8628bbca204c/alexa-skill-teardown-understanding-entity-resolution-with-pet-match

### Credits
basic code thanks to: https://github.com/soloworks/go-alexa-models