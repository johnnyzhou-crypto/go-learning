package errorhandler

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	NoType = ErrorType(iota)
	BadLimitOrder
	OrderNotFound
	//todo: more error type according to the modules
)

type ErrorType uint

type customisedError struct {
	errorType   ErrorType
	originError error
	contextInfo errorContext
}

type errorContext struct {
	Field   string
	Message string
}

//directly return a customised error
func (error customisedError) Error() string {
	return error.originError.Error()
}

//new a customised error
func (errorType ErrorType) New(msg string) error {
	return customisedError{errorType: errorType, originError: errors.New(msg)}
}

// NewF : new a customisedError alone with formatted message
func (errorType ErrorType) NewF(msg string, args ...interface{}) error {
	//err := fmt.Sprintf(msg, args...)
	return customisedError{errorType: errorType, originError: fmt.Errorf(msg, args...)}
}

// Wrap
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrap(err, msg)
}

// WrapF wrap with format
func (errorType ErrorType) WrapF(err error, msg string, args ...interface{}) error {
	formattedErr := errors.Wrapf(err, msg, args...)
	return customisedError{errorType: errorType, originError: formattedErr}
}

// Wrapf directly return a new formatted error
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedErr := errors.Wrapf(err, msg, args...)
	if customisedE, ok := err.(customisedError); ok {
		return customisedError{
			errorType:   customisedE.errorType,
			originError: wrappedErr,
			contextInfo: customisedE.contextInfo,
		}
	}
	return customisedError{errorType: NoType, originError: wrappedErr}
}

// Cause return original error
func Cause(err error) error {
	return errors.Cause(err)
}

func GetErrorType(err error) ErrorType {
	if customE, ok := err.(customisedError); ok {
		return customE.errorType
	}
	return NoType
}

func AddContextToError(err error, field, msg string) error {
	ctx := errorContext{Field: field, Message: msg}
	if customE, ok := err.(customisedError); ok {
		return customisedError{errorType: customE.errorType, originError: customE.originError, contextInfo: ctx}
	}
	return customisedError{errorType: NoType, originError: err, contextInfo: ctx}
}

// return the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customE, ok := err.(customisedError); ok || customE.contextInfo != emptyContext {
		return map[string]string{"field": customE.contextInfo.Field, "message": customE.contextInfo.Message}
	}
	return nil
}
