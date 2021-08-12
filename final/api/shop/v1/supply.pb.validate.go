// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/v1/supply.proto

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

// Validate checks the field values on SupplyInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SupplyInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SupplyId

	// no validation rules for SupplyName

	return nil
}

// SupplyInfoValidationError is the validation error returned by
// SupplyInfo.Validate if the designated constraints aren't met.
type SupplyInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SupplyInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SupplyInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SupplyInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SupplyInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SupplyInfoValidationError) ErrorName() string { return "SupplyInfoValidationError" }

// Error satisfies the builtin error interface
func (e SupplyInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSupplyInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SupplyInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SupplyInfoValidationError{}

// Validate checks the field values on CreateSupplyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSupplyRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SupplyName

	return nil
}

// CreateSupplyRequestValidationError is the validation error returned by
// CreateSupplyRequest.Validate if the designated constraints aren't met.
type CreateSupplyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSupplyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSupplyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSupplyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSupplyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSupplyRequestValidationError) ErrorName() string {
	return "CreateSupplyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSupplyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSupplyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSupplyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSupplyRequestValidationError{}

// Validate checks the field values on CreateSupplyReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateSupplyReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSupplyInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSupplyReplyValidationError{
				field:  "SupplyInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateSupplyReplyValidationError is the validation error returned by
// CreateSupplyReply.Validate if the designated constraints aren't met.
type CreateSupplyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSupplyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSupplyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSupplyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSupplyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSupplyReplyValidationError) ErrorName() string {
	return "CreateSupplyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSupplyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSupplyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSupplyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSupplyReplyValidationError{}

// Validate checks the field values on UpdateSupplyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateSupplyRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSupplyInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSupplyRequestValidationError{
				field:  "SupplyInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateSupplyRequestValidationError is the validation error returned by
// UpdateSupplyRequest.Validate if the designated constraints aren't met.
type UpdateSupplyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSupplyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSupplyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSupplyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSupplyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSupplyRequestValidationError) ErrorName() string {
	return "UpdateSupplyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSupplyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSupplyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSupplyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSupplyRequestValidationError{}

// Validate checks the field values on UpdateSupplyReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateSupplyReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSupplyInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSupplyReplyValidationError{
				field:  "SupplyInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateSupplyReplyValidationError is the validation error returned by
// UpdateSupplyReply.Validate if the designated constraints aren't met.
type UpdateSupplyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSupplyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSupplyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSupplyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSupplyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSupplyReplyValidationError) ErrorName() string {
	return "UpdateSupplyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSupplyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSupplyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSupplyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSupplyReplyValidationError{}

// Validate checks the field values on DeleteSupplyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteSupplyRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSupplyInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteSupplyRequestValidationError{
				field:  "SupplyInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeleteSupplyRequestValidationError is the validation error returned by
// DeleteSupplyRequest.Validate if the designated constraints aren't met.
type DeleteSupplyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSupplyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSupplyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSupplyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSupplyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSupplyRequestValidationError) ErrorName() string {
	return "DeleteSupplyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSupplyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSupplyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSupplyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSupplyRequestValidationError{}

// Validate checks the field values on DeleteSupplyReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteSupplyReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// DeleteSupplyReplyValidationError is the validation error returned by
// DeleteSupplyReply.Validate if the designated constraints aren't met.
type DeleteSupplyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSupplyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSupplyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSupplyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSupplyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSupplyReplyValidationError) ErrorName() string {
	return "DeleteSupplyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSupplyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSupplyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSupplyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSupplyReplyValidationError{}

// Validate checks the field values on GetSupplyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetSupplyRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SupplyId

	return nil
}

// GetSupplyRequestValidationError is the validation error returned by
// GetSupplyRequest.Validate if the designated constraints aren't met.
type GetSupplyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSupplyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSupplyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSupplyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSupplyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSupplyRequestValidationError) ErrorName() string { return "GetSupplyRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetSupplyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSupplyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSupplyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSupplyRequestValidationError{}

// Validate checks the field values on GetSupplyReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetSupplyReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSupplyInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSupplyReplyValidationError{
				field:  "SupplyInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetSupplyReplyValidationError is the validation error returned by
// GetSupplyReply.Validate if the designated constraints aren't met.
type GetSupplyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSupplyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSupplyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSupplyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSupplyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSupplyReplyValidationError) ErrorName() string { return "GetSupplyReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetSupplyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSupplyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSupplyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSupplyReplyValidationError{}

// Validate checks the field values on ListSupplyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListSupplyRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListSupplyRequestValidationError is the validation error returned by
// ListSupplyRequest.Validate if the designated constraints aren't met.
type ListSupplyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSupplyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSupplyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSupplyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSupplyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSupplyRequestValidationError) ErrorName() string {
	return "ListSupplyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSupplyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSupplyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSupplyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSupplyRequestValidationError{}

// Validate checks the field values on ListSupplyReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListSupplyReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSupplies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSupplyReplyValidationError{
					field:  fmt.Sprintf("Supplies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListSupplyReplyValidationError is the validation error returned by
// ListSupplyReply.Validate if the designated constraints aren't met.
type ListSupplyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSupplyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSupplyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSupplyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSupplyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSupplyReplyValidationError) ErrorName() string { return "ListSupplyReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListSupplyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSupplyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSupplyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSupplyReplyValidationError{}