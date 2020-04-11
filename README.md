# Sailor

[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)
![](https://github.com/hidalgopl/sailor/workflows/Tests/badge.svg)
[![](https://img.shields.io/docker/pulls/secureapi/sailor)](https://hub.docker.com/r/secureapi/sailor)

Sailor is command line tool for security testing your web APIs. Developed and maintained by SecureAPI


## Quickstart
To run security checks on your API, set `url` you want to test and your SecureAPI `username` and `accessKey`  in `config.yaml`
 
`sailor run --config=example_config.yaml`

## Demo
![run demo](rundemo.gif)
#### Example config

| Config key | config value | Description |
| ---------- | ------------ | ----------- |
|  username  |   hidalgopl  | Your SecureAPI username |
| accessKey  | 74nfdj3n...2342 | Your SecureAPI access key |
|    url     | https://secureapi.dev/demo | URL you want to test|

That's all. That's it. Then simply run it by typing `sailor run`!

Sailor will produce output:
```bash
INFO[0000] Authenticated for hidalgopl                  
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0007 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0006 : result: passed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0001 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0005 : result: passed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0004 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0009 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0003 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0002 : result: failed 
INFO[0000] [bp83vtn69kffdtvh7av0] -> SEC0008 : result: failed 
INFO[0000] all tasks executed successfully. Link to your test suite: http://secureapi.com/tests/hidalgopl/bp83vtn69kffdtvh7av0 
```

You can click on the link to your tests and learn there how to fix your security issues.

## Install
Binary downloads of the Sailor can be found on [the Releases page](https://github.com/hidalgopl/sailor/releases/latest).

Unpack the `sailor` binary and add it to your PATH and you are good to go!

### Compile from source
Clone this repository and run `make build`. 


## Send us feedback
We would love to hear your feedback. We know that no one has time and will to deal with long survey, so we build feedback collector directly into sailor.
Simply type `sailor feedback` and answer 5 questions (3 are 0-5 scale, only one open question, so you don't waste your time).
Check it out:
[![feedback demo](feedbackdemo.gif)
## CI / CD
Since sailor is single binary, it's really easy to incorporate it in your CI / CD cycles.
### Jenkins integration

### Gitlab integration
Add `SECUREAPI_USERNAME` & `SECUREAPI_ACCESS_KEY` to CI/CD variables in Gitlab UI.
`.gitlab-ci.yml`
```yaml
stages:
 - sectests

secureapi:
  image: secureapi/sailor:latest
  stage: sectests
  script:
    - cat <<EOF > secureapi-config.yaml
      username: "$SECUREAPI_USERNAME"
      accessKey: "$SECUREAPI_ACCESS_KEY"
      EOF
    - sailor run --config=secureapi-config.yaml
```

### Bitbucket pipelines integration
Add `SECUREAPI_USERNAME` & `SECUREAPI_ACCESS_KEY` to bitbucket variables.
`bitbucket-pipelines.yml`
```yaml
image: secureapi/sailor:latest
pipelines:
  default:
    - step:
        name: Create config
        script:
          - cat <<EOF > secureapi-config.yaml
            username: "$SECUREAPI_USERNAME"
            accessKey: "$SECUREAPI_ACCESS_KEY"
            EOF
    - step:
        name: Run tests
        script:
          - sailor run --config=secureapi-config.yaml

```

### Github actions integration

### CircleCI
Set env variables in CircleCI project:
Add `SECUREAPI_USERNAME` & `SECUREAPI_ACCESS_KEY` to env variables in CircleCI UI.
```yaml
    version: 2.1
    executors:
        docker:
          - image: secureapi/sailor:latest
    jobs:
      sec_test:
        steps:
          - run:
              name: Create config
              command: |
                             cat <<EOF > secureapi-config.yaml
                             username: "$SECUREAPI_USERNAME"
                             accessKey: "$SECUREAPI_ACCESS_KEY"
                             EOF
          - run:
              name: Run tests
              command: sailor run --config=secureapi-config.yaml
    workflows:
      version: 2
      build-master:
        jobs:
          - sec_test
```

### TeamCity

### TravisCI
Add `SECUREAPI_USERNAME` & `SECUREAPI_ACCESS_KEY` env variables to Repository Settings.
```yaml
sudo: required
language: go
services:
  - docker
before_install:
  - docker pull secureapi/sailor:latest
  - docker run -it -d --name build secureapi/sailor:latest bash
  - docker exec build git clone https://github.com/user/product.git
script:
  - docker exec build cmake -H/product -B/_build
  - docker exec build cmake --build /_build
  - docker exec build cmake --build /_build --target documentation
  - docker exec build cmake --build /_build --target run-tests
after_success:
  - docker exec build bash /project/codecov.sh
```
### Bamboo

## How it works
Sailor is a simple tool. It parses the config, sends request to provided URL and then passes all needed information from response to SecureAPI. SecureAPI analyses it and pinpoints where are your security vulnerabilities.
