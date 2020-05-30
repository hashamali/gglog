## gsl: Structure Logging Interface
[![godoc](https://godoc.org/github.com/hashamali/gsl?status.svg)](http://godoc.org/github.com/hashamali/gsl)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gsl/security?label=security&style=flat-square)](https://github.com/hashamali/gsl/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gsl)](https://goreportcard.com/report/github.com/hashamali/gsl)
[![license](https://badgen.net/github/license/hashamali/gsl)](https://opensource.org/licenses/MIT)

A small opinionated log interface for Go for structured logging. Only supports 2 types of logs:

* Info: For informational logging.
* Error: For error logging.

#### API

The `gsl.Log` interface exposes the following methods:

* `With(metadata interface{}) GGLog`: Provide a new logger that will include the metadata in subsequent logs.
* `Info(message string)`: Log the provided message at the info level.
* `Infof(format string, v ...interface{})`: Format and log the provided message at the info level.
* `Infow(metadata interface{}, message string)`: A shortened version of `With(metadata).Info(message)`.
* `Infofw(metadata interface{}, format string, v ...interface{})`: A shortened version of `With(metadata).Infof(message, v...)`.
* `Error(message string)`: Log the provided message at the error level.
* `Errorf(format string, v ...interface{})`: Format and log the provided message at the error level.
* `Errorw(metadata interface{}, message string)`:  A shortened version of `With(metadata).Error(message)`.
* `Errorfw(metadata interface{}, format string, v ...interface{})`: A shortened version of `With(metadata).Errorf(message, v...)`.
* `V(verbosity int) bool`: Determine if the provided verbosity is supported.

## Implementations

* [Zerolog](https://github.com/rs/zerolog): [gsl-zl](https://github.com/hashamali/gsl-zl)
