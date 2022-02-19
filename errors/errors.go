package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

// ----------tags----------

// Error tags on file system
const (
	ErrTagFilePathErr ErrorTag = "file_path_error"
	ErrTagFileReadErr ErrorTag = "file_read_error"
)

// Error tags on network
const (
	ErrTagNetworkConFailure      ErrorTag = "connection_failure"
	ErrTagNetworkRetryTimesLimit ErrorTag = "retry_too_many_times"
)

// Error tags on time
const (
	ErrTagTimeExpire  ErrorTag = "expire"
	ErrTagTimeTimeout ErrorTag = "timeout"
)

// Error tags on crypto
const (
	ErrTagCryptoInvalidPrivateKey ErrorTag = "invalid_private_key"
	ErrTagCryptoInvalidPublicKey  ErrorTag = "invalid_public_key"
)

// Error tags on JWT authentication
const (
	ErrTagAuthenExpired    ErrorTag = "jwt_expired"
	ErrTagAuthenInvalidJwt ErrorTag = "invalid_jwt"
)

// ----------error types----------

const (
	ErrTypeParseRSAKey ErrorType = "ParseRSAKeyError"
)

// ----------field names----------

// Field name on call stack
const (
	FieldNameCaller FieldName = "caller" // function that calls worker to do stuff
	FieldNameWorker FieldName = "worker" // function that throws error
)

// Field name of MyError type
const (
	FieldNameId         FieldName = "id"
	FieldNameMsg        FieldName = "msg"
	FieldNameErrType    FieldName = "error_type"
	FieldNameErrCode    FieldName = "error_code"
	FieldNameWhat       FieldName = "what"
	FieldNameWho        FieldName = "who"
	FieldNameWhen       FieldName = "when"
	FieldNameWhy        FieldName = "why"
	FieldNameWrappedErr FieldName = "wrapped_err"
	FieldNameDetails    FieldName = "details"
	FieldNameTags       FieldName = "tags"

	FieldNameFrom  FieldName = "from"
	FieldNameOrder FieldName = "order"
)

var FieldNames = map[string]struct{}{
	"id":          struct{}{},
	"msg":         struct{}{},
	"error_type":  struct{}{},
	"error_code":  struct{}{},
	"what":        struct{}{},
	"who":         struct{}{},
	"when":        struct{}{},
	"why":         struct{}{},
	"wrapped_err": struct{}{},
	"details":     struct{}{},
	"tags":        struct{}{},
}

type FieldName string

type ErrorType string
type ErrorCode string

type ErrorTag string

type Error interface {
	Error() string
	ErrorFall() string
	Errors() int // total number of error in this error chain

	Wrap(err error) *MyError
	Unwrap() error

	As(t ErrorType) bool
	Is(target *MyError) bool

	WithMsg(m string) *MyError
	WithTag(t ErrorTag) *MyError
	WithTags(ts []ErrorTag) *MyError
	WithErrType(t ErrorType) *MyError
	WithErrCode(c ErrorCode) *MyError
	WithWhat(w string) *MyError
	WithWho(w string) *MyError
	WithWhen(w string) *MyError // timestamp
	WithWhy(w string) *MyError
	WithFiled(k, v string) *MyError
	WithCaller(c string) *MyError
	WithWorker(w string) *MyError

	GetID() string
	GetMsg() string
	GetErrType() ErrorType
	GetErrCode() ErrorCode
	GetWhat() string
	GetWho() string
	GetWhen() string
	GetWhy() string
	GetCaller() string
	GetWorker() string

	TagExists(t ErrorTag) bool
	GetTags() []ErrorTag
	GetField(k string) string
	GetFields() map[string]string
	GetFieldsAndTags() (map[string]string, []ErrorTag)
}

var _ Error = (*MyError)(nil)

// MyError represents a common error type that implements
// the Error interface as defined above.
type MyError struct {
	ID  string `json:"id"`
	Msg string `json:"msg,omitempty"`

	ErrType ErrorType `json:"error_type,omitempty"`
	ErrCode ErrorCode `json:"error_code,omitempty"`

	What string `json:"what,omitempty"`
	Who  string `json:"who,omitempty"`
	When string `json:"when,omitempty"` // timestamp
	Why  string `json:"why,omitempty"`

	WrappedErr error `json:"-"`

	Details map[string]string `json:"details,omitempty"`

	Tags map[ErrorTag]struct{} `json:"tags,omitempty"`
}

func New() *MyError {
	e := &MyError{}
	e.ID = uuid.NewString()
	e.Details = make(map[string]string)
	e.Tags = make(map[ErrorTag]struct{})
	return e
}

func (e *MyError) Error() string {
	if e == nil {
		return ""
	}

	var total = e.Errors()

	var strs string
	var p error

	for p = e; p != nil; {
		myErr, ok := p.(*MyError)
		var str string
		str = fmt.Sprintf("%s:%s,", string(FieldNameOrder), strconv.Itoa(total))

		if ok {
			//str += fmt.Sprintf("%q:%q,", string(FieldNameId), myErr.ID)

			if myErr.Tags != nil {
				var tags string
				for t, _ := range myErr.Tags {
					tags += fmt.Sprintf("%s,", t)

				}
				tags = "{" + tags + "}"
				str += fmt.Sprintf("%s:%s,", string(FieldNameTags), tags)
			}

			if myErr.Msg != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameMsg), myErr.Msg)
			}
			if myErr.ErrType != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameErrType), myErr.ErrType)
			}
			if myErr.ErrCode != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameErrCode), myErr.ErrCode)
			}
			if myErr.What != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameWhat), myErr.What)
			}
			if myErr.Who != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameWho), myErr.Who)
			}
			if myErr.When != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameWhen), myErr.When)
			}
			if myErr.Why != "" {
				str += fmt.Sprintf("%s:%s,", string(FieldNameWhy), myErr.Why)
			}

			if len(myErr.Details) != 0 {
				var details string
				for k, v := range myErr.Details {
					details += fmt.Sprintf("%s:%s,", k, v)
				}
				details = "{" + details[:len(details)-1] + "}"
				str += fmt.Sprintf("%s:%s,", string(FieldNameDetails), details)
			}

			str = "{" + str[:len(str)-1] + "}"
			strs += str + "\n"
		} else {
			var from string
			from += fmt.Sprintf("{%s:%s,%s:%s}", string(FieldNameOrder), strconv.Itoa(total), string(FieldNameFrom), p.Error())
			strs += from + "\n"
			//break
		}

		total -= 1
		p = errors.Unwrap(p)
	}

	return strs

}

func (e *MyError) ErrorFall() string {
	if e == nil {
		return ""
	}

	var total = e.Errors()

	var p error
	var str string
	//var pre = "\n" + "-----error " + strconv.Itoa(total) + ":" + "\n"

	for p = e; p != nil; {
		myErr, ok := p.(*MyError)
		if ok {
			bytes, _ := json.MarshalIndent(myErr, "", "  ")
			str += "\n" + "-----error " + strconv.Itoa(total) + ":" + "\n" + string(bytes) + "\n"
		} else {
			str += "\n" + "-----error " + strconv.Itoa(total) + ":" + "\n" + p.Error() + "\n"
		}

		total -= 1
		p = errors.Unwrap(p)
	}

	return str
}

func (e *MyError) Errors() int {
	if e == nil {
		return 0
	}

	var total int
	var p error

	for p = e; p != nil; {
		total++
		p = errors.Unwrap(p)
	}

	return total
}

func (e *MyError) Wrap(err error) *MyError {
	if e == nil {
		println("an nil error is not allowed to wrap an error")
		os.Exit(1)
	}
	e.WrappedErr = err
	return e
}

func (e *MyError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.WrappedErr
}

func (e *MyError) As(t ErrorType) bool {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	return e.ErrType == t
}

func (e *MyError) Is(target *MyError) bool {
	if e == nil && target == nil {
		return true
	}

	if e == nil && target != nil {
		return false
	}

	if e != nil && target == nil {
		return false
	}

	return e.ID == target.ID
}

// -----------------

func (e *MyError) WithMsg(m string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Msg = m
	return e
}

func (e *MyError) WithTag(t ErrorTag) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Tags[t] = struct{}{}
	return e
}

func (e *MyError) WithTags(ts []ErrorTag) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	for _, t := range ts {
		e.Tags[t] = struct{}{}

	}
	return e
}

func (e *MyError) WithErrType(t ErrorType) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.ErrType = t
	return e
}

func (e *MyError) WithErrCode(c ErrorCode) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.ErrCode = c
	return e
}

func (e *MyError) WithWhat(w string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.What = w
	return e
}

func (e *MyError) WithWho(w string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Who = w
	return e
}

func (e *MyError) WithWhen(w string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.When = w
	return e
}

func (e *MyError) WithWhy(w string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Why = w
	return e
}

func (e *MyError) WithFiled(k, v string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Details[k] = v
	return e
}

func (e *MyError) WithCaller(c string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Details[string(FieldNameCaller)] = c
	return e
}

func (e *MyError) WithWorker(w string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.Details[string(FieldNameWorker)] = w
	return e
}

// -----------------

func (e *MyError) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *MyError) GetMsg() string {
	if e == nil {
		return ""
	}
	return e.Msg
}

func (e *MyError) GetErrType() ErrorType {
	if e == nil {
		return ""
	}
	return e.ErrType
}

func (e *MyError) GetErrCode() ErrorCode {
	if e == nil {
		return ""
	}
	return e.ErrCode
}
func (e *MyError) GetWhat() string {
	if e == nil {
		return ""
	}
	return e.What
}

func (e *MyError) GetWho() string {
	if e == nil {
		return ""
	}
	return e.Who
}

func (e *MyError) GetWhen() string {
	if e == nil {
		return ""
	}
	return e.When
}

func (e *MyError) GetWhy() string {
	if e == nil {
		return ""
	}
	return e.Why
}

func (e *MyError) GetCaller() string {
	if e == nil {
		return ""
	}
	return e.Details[string(FieldNameCaller)]
}

func (e *MyError) GetWorker() string {
	if e == nil {
		return ""
	}
	return e.Details[string(FieldNameWorker)]
}

// -----------------

func (e *MyError) TagExists(t ErrorTag) bool {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	_, ok := e.Tags[t]
	return ok
}

func (e *MyError) GetTags() []ErrorTag {
	if e == nil {
		return nil
	}
	var tags []ErrorTag
	for t, _ := range e.Tags {
		tags = append(tags, t)
	}
	return tags
}

func (e *MyError) GetField(k string) string {
	if e == nil {
		return ""
	}
	_, ok := FieldNames[k]
	if ok {
		switch k {
		case string(FieldNameId):
			return e.ID
		case string(FieldNameMsg):
			return e.Msg
		case string(FieldNameErrType):
			return string(e.ErrType)
		case string(FieldNameErrCode):
			return string(e.ErrCode)
		case string(FieldNameWhat):
			return e.What
		case string(FieldNameWho):
			return e.Who
		case string(FieldNameWhen):
			return e.When
		case string(FieldNameWhy):
			return e.Why
		case string(FieldNameWrappedErr):
			return e.WrappedErr.Error()
		case string(FieldNameDetails):
			str, _ := json.MarshalIndent(e.Details, "", "  ")
			return string(str)
		}
	}

	return e.Details[k]
}

func (e *MyError) GetFields() map[string]string {
	if e == nil {
		return nil
	}
	var fields = make(map[string]string)

	fields[string(FieldNameId)] = e.ID

	if e.Msg != "" {
		fields[string(FieldNameMsg)] = e.Msg
	}
	if e.ErrType != "" {
		fields[string(FieldNameErrType)] = string(e.ErrType)
	}
	if e.ErrCode != "" {
		fields[string(FieldNameErrCode)] = string(e.ErrCode)
	}
	if e.What != "" {
		fields[string(FieldNameWhat)] = e.What
	}
	if e.Who != "" {
		fields[string(FieldNameWho)] = e.Who
	}
	if e.When != "" {
		fields[string(FieldNameWhen)] = e.When
	}
	if e.Why != "" {
		fields[string(FieldNameWhy)] = e.Why
	}
	if e.WrappedErr != nil {
		fields[string(FieldNameWrappedErr)] = e.WrappedErr.Error()
	}

	for k, v := range e.Details {
		fields[k] = v
	}

	return fields
}

func (e *MyError) GetFieldsAndTags() (map[string]string, []ErrorTag) {
	if e == nil {
		return nil, nil
	}
	return e.GetFields(), e.GetTags()
}
