package example

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

// directly return a customised error
func (error customisedError) Error() string {
	return error.originError.Error()
}

// new a customised error
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

//https://learnku.com/go/t/33210
//http://jintang.zone/2021/04/28/go%E4%B8%AD%E7%9A%84err!=nil%E5%A5%BD%E7%83%A6%EF%BC%9F%E6%95%99%E4%BD%A0%E6%80%8E%E4%B9%88%E4%BC%98%E9%9B%85%E5%9C%B0%E5%A4%84%E7%90%86error.html
//golang 处理error的经验总结

//three types of handling error strategy
//https://www.zhihu.com/question/27158146
