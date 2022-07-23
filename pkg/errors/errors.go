package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

func raise(msg string, leftLines int) error {
	const (
		bufferLength = 1024
		double       = 2
	)

	buf := make([]byte, bufferLength)

	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			a := strings.Split(string(buf[:n]), "\n")
			b := strings.Join(a[leftLines:], "\n")

			newMsg := fmt.Sprintf("\n%s\n%s", msg, b)

			return errors.New(newMsg)
		}

		doubledBuffer := double * len(buf)
		buf = make([]byte, doubledBuffer)
	}
}

func Wrap(err error) error {
	const length = 5

	return raise(err.Error(), length)
}

// nolint
func newError(msg string) error {
	const length = 6

	return raise(msg, length)
}
