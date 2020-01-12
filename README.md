# Sailor

[![Go Report Card](https://goreportcard.com/badge/github.com/hidalgopl/sailor)](https://goreportcard.com/report/github.com/hidalgopl/sailor)

Sailor is command line tool for security testing your web APIs. Developed and maintained by SecureAPI


## Quickstart
To run security checks on your API, set `SECUREAPI_USERNAME` and `SECUREAPI_ACCESS_KEY` environment variables. Alternatively, you can pass them in `config.yaml`
 
`sailor run --config=example_config.yaml`
#### Example config
```config.yaml
username: "your username goes here"
accessKey: "your access key goes here"
url: "https://secureapi.dev/demo" # URL you want to test
```

Sailor will produce output:
```bash

```

You can click on the link to your tests and learn there how to fix your security issues.

## Install
Binary downloads of the Sailor can be found on [the Releases page](https://github.com/hidalgopl/sailor/releases/latest).

Unpack the `sailor` binary and add it to your PATH and you are good to go!

### Compile from source
Clone this repository and run `make build`. 

