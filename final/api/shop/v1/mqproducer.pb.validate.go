// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/v1/mqproducer.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Products with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Products) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProductId

	// no validation rules for ProductName

	return nil
}

// ProductsValidationError is the validation error returned by
// Products.Validate if the designated constraints aren't met.
type ProductsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductsValidationError) ErrorName() string { return "ProductsValidationError" }

// Error satisfies the builtin error interface
func (e ProductsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProducts.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductsValidationError{}

// Validate checks the field values on CreateMqproducerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMqproducerRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProducts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMqproducerRequestValidationError{
				field:  "Products",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateMqproducerRequestValidationError is the validation error returned by
// CreateMqproducerRequest.Validate if the designated constraints aren't met.
type CreateMqproducerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMqproducerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMqproducerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMqproducerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMqproducerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMqproducerRequestValidationError) ErrorName() string {
	return "CreateMqproducerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMqproducerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMqproducerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMqproducerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMqproducerRequestValidationError{}

// Validate checks the field values on CreateMqproducerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMqproducerReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for MqId

	return nil
}

// CreateMqproducerReplyValidationError is the validation error returned by
// CreateMqproducerReply.Validate if the designated constraints aren't met.
type CreateMqproducerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMqproducerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMqproducerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMqproducerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMqproducerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMqproducerReplyValidationError) ErrorName() string {
	return "CreateMqproducerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMqproducerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMqproducerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMqproducerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMqproducerReplyValidationError{}

// Validate checks the field values on UpdateMqproducerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMqproducerRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateMqproducerRequestValidationError is the validation error returned by
// UpdateMqproducerRequest.Validate if the designated constraints aren't met.
type UpdateMqproducerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMqproducerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMqproducerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMqproducerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMqproducerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMqproducerRequestValidationError) ErrorName() string {
	return "UpdateMqproducerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMqproducerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMqproducerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMqproducerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMqproducerRequestValidationError{}

// Validate checks the field values on UpdateMqproducerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMqproducerReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateMqproducerReplyValidationError is the validation error returned by
// UpdateMqproducerReply.Validate if the designated constraints aren't met.
type UpdateMqproducerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMqproducerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMqproducerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMqproducerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMqproducerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMqproducerReplyValidationError) ErrorName() string {
	return "UpdateMqproducerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMqproducerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMqproducerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMqproducerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMqproducerReplyValidationError{}

// Validate checks the field values on DeleteMqproducerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMqproducerRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteMqproducerRequestValidationError is the validation error returned by
// DeleteMqproducerRequest.Validate if the designated constraints aren't met.
type DeleteMqproducerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMqproducerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMqproducerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMqproducerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMqproducerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMqproducerRequestValidationError) ErrorName() string {
	return "DeleteMqproducerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMqproducerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMqproducerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMqproducerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMqproducerRequestValidationError{}

// Validate checks the field values on DeleteMqproducerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMqproducerReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteMqproducerReplyValidationError is the validation error returned by
// DeleteMqproducerReply.Validate if the designated constraints aren't met.
type DeleteMqproducerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMqproducerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMqproducerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMqproducerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMqproducerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMqproducerReplyValidationError) ErrorName() string {
	return "DeleteMqproducerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMqproducerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMqproducerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMqproducerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMqproducerReplyValidationError{}

// Validate checks the field values on GetMqproducerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMqproducerRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetMqproducerRequestValidationError is the validation error returned by
// GetMqproducerRequest.Validate if the designated constraints aren't met.
type GetMqproducerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMqproducerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMqproducerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMqproducerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMqproducerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMqproducerRequestValidationError) ErrorName() string {
	return "GetMqproducerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMqproducerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMqproducerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMqproducerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMqproducerRequestValidationError{}

// Validate checks the field values on GetMqproducerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMqproducerReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetMqproducerReplyValidationError is the validation error returned by
// GetMqproducerReply.Validate if the designated constraints aren't met.
type GetMqproducerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMqproducerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMqproducerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMqproducerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMqproducerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMqproducerReplyValidationError) ErrorName() string {
	return "GetMqproducerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetMqproducerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMqproducerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMqproducerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMqproducerReplyValidationError{}

// Validate checks the field values on ListMqproducerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMqproducerRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListMqproducerRequestValidationError is the validation error returned by
// ListMqproducerRequest.Validate if the designated constraints aren't met.
type ListMqproducerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMqproducerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMqproducerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMqproducerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMqproducerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMqproducerRequestValidationError) ErrorName() string {
	return "ListMqproducerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMqproducerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMqproducerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMqproducerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMqproducerRequestValidationError{}

// Validate checks the field values on ListMqproducerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMqproducerReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListMqproducerReplyValidationError is the validation error returned by
// ListMqproducerReply.Validate if the designated constraints aren't met.
type ListMqproducerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMqproducerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMqproducerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMqproducerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMqproducerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMqproducerReplyValidationError) ErrorName() string {
	return "ListMqproducerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListMqproducerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMqproducerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMqproducerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMqproducerReplyValidationError{}
