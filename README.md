# Sailor

[![Go Report Card](https://goreportcard.com/badge/github.com/hidalgopl/sailor)](https://goreportcard.com/report/github.com/hidalgopl/sailor)

Sailor is command line tool for security testing your web APIs. Developed and maintained by SecureAPI


## Quickstart
To run security checks on your API, set `SECUREAPI_USERNAME` and `SECUREAPI_ACCESS_KEY` environment variables. Alternatively, you can pass them in `config.yaml`
 
`sailor run --config=example_config.yaml`

## Demo
[![asciicast](https://asciinema.org/a/1cgbJTCXD1iRaQCXy818691Pz.svg)](https://asciinema.org/a/1cgbJTCXD1iRaQCXy818691Pz)
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
[![asciicast](https://asciinema.org/a/zbPZ2oM0iwJMSxEeiJJP1Nu86.svg)](https://asciinema.org/a/zbPZ2oM0iwJMSxEeiJJP1Nu86)
## CI / CD
Since sailor is single binary, it's really easy to incorporate it in your CI / CD cycles.
### Jenkins integration

### Gitlab integration

### Bitbucket pipelines integration

### Github actions integration

### CircleCI

### TeamCity

### TravisCI

### Bamboo

## How it works
Sailor is a simple tool. It parses the config, sends request to provided URL and then passes all needed information from response to SecureAPI. SecureAPI analyses it and pinpoints where are your security vulnerabilities.
