# Duty
[![Build Status](https://travis-ci.org/gomicro/duty.svg)](https://travis-ci.org/gomicro/duty)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/duty)](https://goreportcard.com/report/github.com/gomicro/duty)
[![GoDoc](https://godoc.org/github.com/gomicro/duty?status.svg)](https://godoc.org/github.com/gomicro/duty)
[![License](https://img.shields.io/github/license/gomicro/duty.svg)](https://github.com/gomicro/duty/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/duty.svg)](https://github.com/gomicro/duty/releases/latest)

Duty is for returning mocked, semi-static http responses. Intended to mimic a service for testing purposes, duty enables returning responses that may typically be difficult to force the service to return. Such as many error states that only occur in catastrophic failures like failing to perform OS level functions etc. Duty allows for you to provide a config for the endpoints it will respond to and various aspects of how the response will look. The benefit coming from it's simplicity and being light weight.  The container houses only the app, and as such is a very tiny docker image.

# Requirements
Docker

# Usage

## Running
Duty is intended to be used in conjunction with local Docker testing of a service.

```
docker pull gomicro/duty
docker run -it -v $PWD/duty.yaml:/duty.yaml -v $PWD/responses:/responses gomicro/avenues
```

# Versioning
The app will be versioned in accordance with [Semver 2.0.0](http://semver.org).  See the [releases](https://github.com/gomicro/avenues/releases) section for the latest version.  Until version 1.0.0 the app is considered to be unstable.

# License
See [LICENSE.md](./LICENSE.md) for more information.
