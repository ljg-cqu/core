package _errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ----------tags----------

// Error tags on type operation
const (
	ErrTagTypeAssertion ErrorTag = "type_assertion_err"
)

// Error tags on file system
const (
	ErrTagFilePathErr ErrorTag = "file_path_error"
	ErrTagFileReadErr ErrorTag = "file_read_error"
)

// Error tags on encode and decode
const (
	ErrTagJsonUnmarshalErr ErrorTag = "json_unmarshal_error"
)

// Error tags on network
const (
	ErrTagNetworkConFailure      ErrorTag = "connection_failure"
	ErrTagNetworkRetryTimesLimit ErrorTag = "retry_too_many_times"
)

// Error tags on http communication
const (
	ErrTagHttpHandshak ErrorTag = "http_handshak_failure"
	ErrTagHttpRequest  ErrorTag = "http_request_error"
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
	FieldNameId           FieldName = "id"
	FieldNameErrMsg       FieldName = "error_msg"
	FieldNameErrType      FieldName = "error_type"
	FieldNameErrCode      FieldName = "error_code"
	FieldNameWhat         FieldName = "what"
	FieldNameWho          FieldName = "who"
	FieldNameWhen         FieldName = "when"
	FieldNameWhy          FieldName = "why"
	FieldNameWrappedErr   FieldName = "wrapped_err"
	FiledNameWrappedMyErr FieldName = "wrapped_my_err"
	FieldNameDetails      FieldName = "details"
	FieldNameTags         FieldName = "tags"

	FieldNameFrom  FieldName = "from"
	FieldNameOrder FieldName = "order"
)

var FieldNames = map[string]struct{}{
	"id":             struct{}{},
	"msg_msg":        struct{}{},
	"error_type":     struct{}{},
	"error_code":     struct{}{},
	"what":           struct{}{},
	"who":            struct{}{},
	"when":           struct{}{},
	"why":            struct{}{},
	"wrapped_my_err": struct{}{},
	"wrapped_err":    struct{}{},
	"details":        struct{}{},
	"tags":           struct{}{},
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
	WithMsgAppend(m string) *MyError
	WithMsgf(format string, a ...any) *MyError
	WithMsgfAppend(format string, a ...any) *MyError
	WithTag(t ErrorTag) *MyError
	WithTags(ts []ErrorTag) *MyError
	WithErrType(t ErrorType) *MyError
	WithErrCode(c ErrorCode) *MyError
	WithWhat(w string) *MyError
	WithWho(w string) *MyError
	WithWhen(w string) *MyError // timestamp
	WithWhy(w string) *MyError
	WithWhyTag(w ErrorTag) *MyError
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
	ID      string    `json:"id,omitempty"`
	ErrMsg  string    `json:"error_msg,omitempty"`
	ErrType ErrorType `json:"error_type,omitempty"`
	ErrCode ErrorCode `json:"error_code,omitempty"`

	What string `json:"what,omitempty"`
	Who  string `json:"who,omitempty"`
	When string `json:"when,omitempty"` // timestamp
	Why  string `json:"why,omitempty"`

	WrappedMyErr *MyError `json:"wrapped_my_err,omitempty"`
	WrappedErr   error    `json:"wrapped_err,omitempty"`

	Details map[string]string `json:"details,omitempty"`

	Tags map[ErrorTag]struct{} `json:"tags,omitempty"`
}

func New() *MyError {
	e := &MyError{}
	//e.ID = uuid.NewString()
	e.Details = make(map[string]string)
	e.Tags = make(map[ErrorTag]struct{})
	return e
}

func NewWithMsg(msg string) *MyError {
	e := &MyError{}
	e.ErrMsg = msg
	//e.ID = uuid.NewString()
	e.Details = make(map[string]string)
	e.Tags = make(map[ErrorTag]struct{})
	return e
}

func NewWithMsgf(format string, a ...any) *MyError {
	msg := fmt.Sprintf(format, a...)
	e := &MyError{}
	e.ErrMsg = msg
	//e.ID = uuid.NewString()
	e.Details = make(map[string]string)
	e.Tags = make(map[ErrorTag]struct{})
	return e
}

func (e *MyError) Error() string {
	var total = e.Errors()
	var strs string

	for myErr := e; myErr != (*MyError)(nil); myErr = myErr.WrappedMyErr {
		var str string
		//str = fmt.Sprintf("%s:%s,", string(FieldNameOrder), strconv.Itoa(total))

		//str += fmt.Sprintf("%q:%q,", string(FieldNameId), myErr.ID)

		if len(myErr.Tags) == 0 {
			var tags string
			for t, _ := range myErr.Tags {
				tags += fmt.Sprintf("%s,", t)

			}
			tags = "{" + tags + "}"
			str += fmt.Sprintf("%s:%s,", string(FieldNameTags), tags)
		}

		if myErr.ErrMsg != "" {
			str += fmt.Sprintf("%s:%s,", string(FieldNameErrMsg), myErr.ErrMsg)
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
			//str += fmt.Sprintf("%s:%s,", string(FieldNameDetails), details)
		}

		str = "{" + str[:len(str)-1] + "}"
		//strs += str + "\n"
		strs += str + ","

		total -= 1
		if myErr.WrappedErr != nil {
			for wrapErr := myErr.WrappedErr; wrapErr != nil; wrapErr = errors.Unwrap(wrapErr) {
				var from string
				//from += fmt.Sprintf("{%s:%s,%s:%s}", string(FieldNameOrder), strconv.Itoa(total), string(FieldNameFrom), wrapErr.Error())
				from += fmt.Sprintf("{%s:%s}", string(FieldNameFrom), wrapErr.Error())

				//strs += from + "\n"
				strs += from + ","

				total -= 1
			}
		}
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
	var total int

	for myErr := e; myErr != (*MyError)(nil); myErr = myErr.WrappedMyErr {
		total++

		for wrapErr := myErr.WrappedErr; wrapErr != nil; wrapErr = errors.Unwrap(wrapErr) {
			total++
		}
	}

	return total
}

func (e *MyError) Wrap(err error) *MyError {
	if e == nil {
		println("an nil error is not allowed to wrap an error")
		os.Exit(1)
	}

	if err == nil {
		return e
	}

	myErr, ok := err.(*MyError)
	if ok {
		e.WrappedMyErr = myErr
		return e
	}

	myErrNew := New()
	e.WrappedErr = err
	e.WrappedMyErr = myErrNew

	return e
}

func (e *MyError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.WrappedMyErr
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
	e.ErrMsg = m
	return e
}

func (e *MyError) WithMsgAppend(m string) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.ErrMsg = e.ErrMsg + "," + m
	return e
}

func (e *MyError) WithMsgf(format string, a ...any) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.ErrMsg = fmt.Sprintf(format, a...)
	return e
}

func (e *MyError) WithMsgfAppend(format string, a ...any) *MyError {
	if e == nil {
		println("an nil error is not allowed to access")
		os.Exit(1)
	}
	e.ErrMsg = e.ErrMsg + "," + fmt.Sprintf(format, a...)
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

func (e *MyError) WithWhyTag(w ErrorTag) *MyError {
	e.Why = string(w)
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
	return e.ErrMsg
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
		case string(FieldNameErrMsg):
			return e.ErrMsg
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
			return e.WrappedMyErr.Error()
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

	if e.ErrMsg != "" {
		fields[string(FieldNameErrMsg)] = e.ErrMsg
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
	if e.WrappedMyErr != nil {
		fields[string(FieldNameWrappedErr)] = e.WrappedMyErr.Error()
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

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErr(msg interface{}) {
	if msg != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", msg)
		os.Exit(1)
	}
}
