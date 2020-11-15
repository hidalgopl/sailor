# Sailor

[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)
![](https://github.com/hidalgopl/sailor/workflows/Tests/badge.svg)
[![](https://img.shields.io/docker/pulls/secureapi/sailor)](https://hub.docker.com/r/secureapi/sailor)

Sailor is command line tool for security testing your web APIs.


## Quickstart - trying locally
To run security checks on your API, set `url` you want to test in `.secureapi.yml` (or appropriate environment variables) and execute this command:
`sailor run`

1. Download latest sailor binary from [the Releases page](https://github.com/hidalgopl/sailor/releases/latest).
2. Set `SECUREAPI_URL` to the URL you want to test.
5. `sailor run` and stay secure!

#### Example config
To generate config template, run `sailor init-config`. This will create `.secureapi.yml` file in following format:

| Config key | config value | Description | Env variable |
| ---------- | ------------ | ----------- | ------------ |
|    url     | https://secureapi.dev/demo | URL you want to test| SECUREAPI_URL |

That's all. That's it. Then simply run it by typing `sailor run`!

Sailor will produce output:
```bash
INFO[0000] Authenticated for hidalgopl                     
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0007 : result: failed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0003 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0002 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0005 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0006 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0004 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0008 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0009 : result: passed 
INFO[0000] [bssgpb55ictis3l56rj0] -> SEC0001 : result: failed 
INFO[0000] all tasks executed successfully. Link to your test suite: https://secureapi.dev?suite-id=bssgpb55ictis3l56rj0 
---------------------------------------------------------------------------------------------
SEC0007: Strict-Transport-Security: max-age=(age in seconds); (other options)
This header lets a web site tell browsers that it should only be accessed using HTTPS, instead of using HTTP.
Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security

---------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------
SEC0001: X-Content-Type-Options: no-sniff
The server should send an X-Content-Type-Options: nosniff 
to make sure the browser does not try to detect a different Content-Type 
than what is actually sent (as this can lead to XSS)
Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options

---------------------------------------------------------------------------------------------
```

You can click on the link to your tests and learn there how to fix your security issues.

## Install
Binary downloads of the Sailor can be found on [the Releases page](https://github.com/hidalgopl/sailor/releases/latest).

Unpack the `sailor` binary and add it to your PATH and you are good to go!

### Compile from source
Clone this repository and run `make build`. 

## CI / CD
Since sailor is single binary, it's really easy to incorporate it in your CI / CD cycles.
### Jenkins integration

### Gitlab integration
Add `SECUREAPI_URL` to CI/CD variables in Gitlab UI.
`.gitlab-ci.yml`
```yaml
stages:
 - sectests

secureapi:
  image: secureapi/sailor:latest
  stage: sectests
  script:
    - sailor run
```

### Bitbucket pipelines integration
Add `SECUREAPI_URL` to bitbucket variables.
`bitbucket-pipelines.yml`
```yaml
image: secureapi/sailor:latest
pipelines:
  default:
    - step:
        name: Run tests
        script:
          - sailor run

```

### Github actions integration
In your deploy repository, set `SECUREAPI_URL` secrets.
Then, just paste this:
```yaml
    - name: Run sailor
      uses: secureapi/sailor-action@master
      with:
        url: ${{ secrets.SECUREAPI_URL }}
```
You can find [secureapi/sailor-action](https://github.com/secureapi/sailor-action) in Github Actions marketplace. We're working hard to keep this as up to date as possible.

### CircleCI
Set env variables in CircleCI project:
Add `SECUREAPI_URL` to env variables in CircleCI UI.
```yaml
    version: 2.1
    executors:
        docker:
          - image: secureapi/sailor:latest
    jobs:
      sec_test:
        steps:
          - run:
              name: Run tests
              command: sailor run
    workflows:
      version: 2
      build-master:
        jobs:
          - sec_test
```

### TeamCity

### TravisCI
Add `SECUREAPI_URL` env variables to Repository Settings.
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
  - docker exec build sailor run
```
### Bamboo

## How it works
Sailor is a simple tool. It parses the config, sends request to provided URL and then passes all needed information from response to SecureAPI. SecureAPI analyses it and pinpoints where are your security vulnerabilities.
