// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/v1/goods.proto

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

// Validate checks the field values on GoodsInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GoodsInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GoodsId

	// no validation rules for GoodsName

	// no validation rules for CatId

	// no validation rules for GoodsPrice

	// no validation rules for GoodsNum

	return nil
}

// GoodsInfoValidationError is the validation error returned by
// GoodsInfo.Validate if the designated constraints aren't met.
type GoodsInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoodsInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoodsInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoodsInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoodsInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoodsInfoValidationError) ErrorName() string { return "GoodsInfoValidationError" }

// Error satisfies the builtin error interface
func (e GoodsInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoodsInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoodsInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoodsInfoValidationError{}

// Validate checks the field values on CreateGoodsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateGoodsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GoodsName

	// no validation rules for CatId

	// no validation rules for GoodsPrice

	// no validation rules for GoodsNum

	return nil
}

// CreateGoodsRequestValidationError is the validation error returned by
// CreateGoodsRequest.Validate if the designated constraints aren't met.
type CreateGoodsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGoodsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGoodsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGoodsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGoodsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGoodsRequestValidationError) ErrorName() string {
	return "CreateGoodsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateGoodsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGoodsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGoodsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGoodsRequestValidationError{}

// Validate checks the field values on CreateGoodsReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateGoodsReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGoods()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGoodsReplyValidationError{
				field:  "Goods",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateGoodsReplyValidationError is the validation error returned by
// CreateGoodsReply.Validate if the designated constraints aren't met.
type CreateGoodsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGoodsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGoodsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGoodsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGoodsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGoodsReplyValidationError) ErrorName() string { return "CreateGoodsReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateGoodsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGoodsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGoodsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGoodsReplyValidationError{}

// Validate checks the field values on UpdateGoodsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateGoodsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGoods()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateGoodsRequestValidationError{
				field:  "Goods",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateGoodsRequestValidationError is the validation error returned by
// UpdateGoodsRequest.Validate if the designated constraints aren't met.
type UpdateGoodsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGoodsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGoodsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGoodsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGoodsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGoodsRequestValidationError) ErrorName() string {
	return "UpdateGoodsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateGoodsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGoodsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGoodsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGoodsRequestValidationError{}

// Validate checks the field values on UpdateGoodsReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateGoodsReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGoods()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateGoodsReplyValidationError{
				field:  "Goods",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateGoodsReplyValidationError is the validation error returned by
// UpdateGoodsReply.Validate if the designated constraints aren't met.
type UpdateGoodsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGoodsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGoodsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGoodsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGoodsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGoodsReplyValidationError) ErrorName() string { return "UpdateGoodsReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateGoodsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGoodsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGoodsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGoodsReplyValidationError{}

// Validate checks the field values on DeleteGoodsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteGoodsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GoodsId

	return nil
}

// DeleteGoodsRequestValidationError is the validation error returned by
// DeleteGoodsRequest.Validate if the designated constraints aren't met.
type DeleteGoodsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGoodsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGoodsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGoodsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGoodsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGoodsRequestValidationError) ErrorName() string {
	return "DeleteGoodsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteGoodsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGoodsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGoodsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGoodsRequestValidationError{}

// Validate checks the field values on DeleteGoodsReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteGoodsReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// DeleteGoodsReplyValidationError is the validation error returned by
// DeleteGoodsReply.Validate if the designated constraints aren't met.
type DeleteGoodsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGoodsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGoodsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGoodsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGoodsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGoodsReplyValidationError) ErrorName() string { return "DeleteGoodsReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteGoodsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGoodsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGoodsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGoodsReplyValidationError{}

// Validate checks the field values on GetGoodsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetGoodsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GoodsId

	return nil
}

// GetGoodsRequestValidationError is the validation error returned by
// GetGoodsRequest.Validate if the designated constraints aren't met.
type GetGoodsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGoodsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGoodsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGoodsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGoodsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGoodsRequestValidationError) ErrorName() string { return "GetGoodsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetGoodsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGoodsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGoodsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGoodsRequestValidationError{}

// Validate checks the field values on GetGoodsReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetGoodsReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGoods()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGoodsReplyValidationError{
				field:  "Goods",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetGoodsReplyValidationError is the validation error returned by
// GetGoodsReply.Validate if the designated constraints aren't met.
type GetGoodsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGoodsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGoodsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGoodsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGoodsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGoodsReplyValidationError) ErrorName() string { return "GetGoodsReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetGoodsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGoodsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGoodsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGoodsReplyValidationError{}

// Validate checks the field values on ListGoodsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListGoodsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListGoodsRequestValidationError is the validation error returned by
// ListGoodsRequest.Validate if the designated constraints aren't met.
type ListGoodsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGoodsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGoodsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGoodsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGoodsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGoodsRequestValidationError) ErrorName() string { return "ListGoodsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListGoodsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGoodsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGoodsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGoodsRequestValidationError{}

// Validate checks the field values on ListGoodsReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListGoodsReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetGoodsList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListGoodsReplyValidationError{
					field:  fmt.Sprintf("GoodsList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListGoodsReplyValidationError is the validation error returned by
// ListGoodsReply.Validate if the designated constraints aren't met.
type ListGoodsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGoodsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGoodsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGoodsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGoodsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGoodsReplyValidationError) ErrorName() string { return "ListGoodsReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListGoodsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGoodsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGoodsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGoodsReplyValidationError{}