package errorhandler

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/errgo.v2/fmt/errors"
	"testing"
)

func TestContext(t *testing.T) {

	err := BadLimitOrder.New("an_error")
	errWithContext := AddContextToError(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, BadLimitOrder, GetErrorType(errWithContext))
	assert.Equal(t, expectedContext, GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := errors.New("a custom error")

	errWithContext := AddContextToError(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, NoType, GetErrorType(errWithContext))
	assert.Equal(t, expectedContext, GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := errors.New("an_error")
	wrappedError := BadLimitOrder.WrapF(err, "error %s", "1")

	assert.Equal(t, BadLimitOrder, GetErrorType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfInNoTypeError(t *testing.T) {
	err := errors.Newf("an_error %s", "2")
	wrappedError := Wrapf(err, "error %s", "1")

	assert.Equal(t, NoType, GetErrorType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
