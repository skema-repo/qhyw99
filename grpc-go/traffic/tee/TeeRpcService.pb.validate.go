// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: TeeRpcService.proto

package tee

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TeeTxRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TeeTxRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TeeTxRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TeeTxRequestMultiError, or
// nil if none found.
func (m *TeeTxRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TeeTxRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHeader()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TeeTxRequestValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TeeTxRequestValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TeeTxRequestValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Payload

	// no validation rules for Signature

	if len(errors) > 0 {
		return TeeTxRequestMultiError(errors)
	}
	return nil
}

// TeeTxRequestMultiError is an error wrapping multiple validation errors
// returned by TeeTxRequest.ValidateAll() if the designated constraints aren't met.
type TeeTxRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeeTxRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeeTxRequestMultiError) AllErrors() []error { return m }

// TeeTxRequestValidationError is the validation error returned by
// TeeTxRequest.Validate if the designated constraints aren't met.
type TeeTxRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeeTxRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeeTxRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeeTxRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeeTxRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeeTxRequestValidationError) ErrorName() string { return "TeeTxRequestValidationError" }

// Error satisfies the builtin error interface
func (e TeeTxRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeeTxRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeeTxRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeeTxRequestValidationError{}

// Validate checks the field values on TeeTxHeader with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TeeTxHeader) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TeeTxHeader with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TeeTxHeaderMultiError, or
// nil if none found.
func (m *TeeTxHeader) ValidateAll() error {
	return m.validate(true)
}

func (m *TeeTxHeader) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BusiId

	// no validation rules for TxId

	// no validation rules for Timestamp

	if len(errors) > 0 {
		return TeeTxHeaderMultiError(errors)
	}
	return nil
}

// TeeTxHeaderMultiError is an error wrapping multiple validation errors
// returned by TeeTxHeader.ValidateAll() if the designated constraints aren't met.
type TeeTxHeaderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeeTxHeaderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeeTxHeaderMultiError) AllErrors() []error { return m }

// TeeTxHeaderValidationError is the validation error returned by
// TeeTxHeader.Validate if the designated constraints aren't met.
type TeeTxHeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeeTxHeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeeTxHeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeeTxHeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeeTxHeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeeTxHeaderValidationError) ErrorName() string { return "TeeTxHeaderValidationError" }

// Error satisfies the builtin error interface
func (e TeeTxHeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeeTxHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeeTxHeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeeTxHeaderValidationError{}

// Validate checks the field values on TeeTransactPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TeeTransactPayload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TeeTransactPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TeeTransactPayloadMultiError, or nil if none found.
func (m *TeeTransactPayload) ValidateAll() error {
	return m.validate(true)
}

func (m *TeeTransactPayload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TeeTransactPayloadMultiError(errors)
	}
	return nil
}

// TeeTransactPayloadMultiError is an error wrapping multiple validation errors
// returned by TeeTransactPayload.ValidateAll() if the designated constraints
// aren't met.
type TeeTransactPayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeeTransactPayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeeTransactPayloadMultiError) AllErrors() []error { return m }

// TeeTransactPayloadValidationError is the validation error returned by
// TeeTransactPayload.Validate if the designated constraints aren't met.
type TeeTransactPayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeeTransactPayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeeTransactPayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeeTransactPayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeeTransactPayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeeTransactPayloadValidationError) ErrorName() string {
	return "TeeTransactPayloadValidationError"
}

// Error satisfies the builtin error interface
func (e TeeTransactPayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeeTransactPayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeeTransactPayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeeTransactPayloadValidationError{}

// Validate checks the field values on TeeTxResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TeeTxResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TeeTxResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TeeTxResponseMultiError, or
// nil if none found.
func (m *TeeTxResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TeeTxResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BusiId

	// no validation rules for TxId

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return TeeTxResponseMultiError(errors)
	}
	return nil
}

// TeeTxResponseMultiError is an error wrapping multiple validation errors
// returned by TeeTxResponse.ValidateAll() if the designated constraints
// aren't met.
type TeeTxResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeeTxResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeeTxResponseMultiError) AllErrors() []error { return m }

// TeeTxResponseValidationError is the validation error returned by
// TeeTxResponse.Validate if the designated constraints aren't met.
type TeeTxResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeeTxResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeeTxResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeeTxResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeeTxResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeeTxResponseValidationError) ErrorName() string { return "TeeTxResponseValidationError" }

// Error satisfies the builtin error interface
func (e TeeTxResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeeTxResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeeTxResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeeTxResponseValidationError{}
