# NOTE: Please refer to https://aka.ms/azsdk/engsys/ci-yaml before editing this file.
trigger:
  paths:
    include:
    - sdk/{{rpName}}/{{packageName}}/

pr:
  paths:
    include:
    - sdk/{{rpName}}/{{packageName}}/

stages:
- template: /eng/pipelines/templates/jobs/archetype-sdk-client.yml
  parameters:
    ServiceDirectory: '{{rpName}}/{{packageName}}'
