///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////
package funky

import (
	"fmt"
)

// IllegalArgumentError  An error indicating that an argument to a method is illegal or invalid.
type IllegalArgumentError string

func (e IllegalArgumentError) Error() string {
	return fmt.Sprintf("The argument is illegal or inappropriate: %s", string(e))
}

// TimeoutError An error indicating that a timeout has been exceeded
type TimeoutError string

func (e TimeoutError) Error() string {
	return string(e)
}

// FunctionServerError a generic error indicating that the function server experienced an error
type FunctionServerError struct {
	APIError Error
}

func (e FunctionServerError) Error() string {
	return "The function server encountered an error. See APIError field for more detail"
}

// ConnectionRefusedError An error indicating that the http connection to the function server was refused
type ConnectionRefusedError string

func (e ConnectionRefusedError) Error() string {
	return fmt.Sprintf("The local function server at %s refused the connection", string(e))
}

// BadRequestError An error indicating that the body of an http request was invalid
type BadRequestError string

func (e BadRequestError) Error() string {
	return fmt.Sprintf("The request body is invalid: %s", string(e))
}

// InvocationError An error indicating that a general error occurred while invoking a function
type InvocationError string

func (e InvocationError) Error() string {
	return fmt.Sprintf("Invocation of the function failed: %s", string(e))
}

// InvalidResponsePayloadError error indicating that function return could not be serialized
type InvalidResponsePayloadError string

func (e InvalidResponsePayloadError) Error() string {
	return fmt.Sprintf("Unable to serialize response payload: %s", string(e))
}

// UnknownSystemError error for when something unknown happens during function invocation
type UnknownSystemError string

func (e UnknownSystemError) Error() string {
	return fmt.Sprintf("Unknown system error: %s", string(e))
}
