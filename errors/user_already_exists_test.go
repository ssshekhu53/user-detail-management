package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserAlreadyExistsError(t *testing.T) {
	assert.EqualError(t, UserAlreadyExists{}, "user already exists with given combination")
}
