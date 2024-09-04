package common

import (
	"fmt"
)

func WrapErr(err error, msg string) (errWithContext error) {
	if err == nil {
		return
	}
	errWithContext = fmt.Errorf("%s: %v", msg, err)
	return
}

func WrapfErr(err error, format string, a ...interface{}) (errWithContext error) {
	if err == nil {
		return
	}
	errWithContext = WrapErr(err, fmt.Sprintf(format, a...))
	return
}
