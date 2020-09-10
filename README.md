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

* `With(fields map[string]interface{}) GGLog`: Provide a new logger that will include the fields in subsequent logs.
* `Info(message string)`: Log the provided message at the info level.
* `Infof(format string, v ...interface{})`: Format and log the provided message at the info level.
* `Infow(fields map[string]interface{}, message string)`: A shortened version of `With(fields).Info(message)`.
* `Infofw(fields map[string]interface{}, format string, v ...interface{})`: A shortened version of `With(fields).Infof(message, v...)`.
* `Error(message string)`: Log the provided message at the error level.
* `Errorf(format string, v ...interface{})`: Format and log the provided message at the error level.
* `Errorw(fields map[string]interface{}, message string)`:  A shortened version of `With(fields).Error(message)`.
* `Errorfw(fields map[string]interface{}, format string, v ...interface{})`: A shortened version of `With(fields).Errorf(message, v...)`.
* `V(verbosity int) bool`: Determine if the provided verbosity is supported.

## Implementations

* [Zerolog](https://github.com/rs/zerolog): [gsl-zl](https://github.com/hashamali/gslzl)
