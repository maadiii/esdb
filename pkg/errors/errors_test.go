package errors_test

import (
	errs "errors"
	"testing"

	"github.com/maadiii/esdb/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	t.Parallel()

	err := errors.Wrap(errs.New("just for test"))
	assert.Contains(t, err.Error(), "just for test")
	assert.Contains(t, err.Error(), "pkg/errors/errors_test.go:14")
}
