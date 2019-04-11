# Purpose
Provide support in localizing an Alexa skill.
* clear and easy structure of translations (one file per locale encouraged)
* register locales (translations), define fallback locales
* separating logic from translations (logic/flow is in the code, e.g. which Intent uses which Slots)

What does `l10n` NOT provide or aim to support:
* it does not aim to be feature complete (yet)