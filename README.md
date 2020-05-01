## stlog: Structure Logging Interface

A small opinionated log interface for Go for structured logging. Only supports 2 types of logs:

* Info: For informational logging.
* Error: For error logging.

#### API

The `stlog.Log` interface exposes the following methods:

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

* Zerolog: https://github.com/hashamali/stlog-zl 