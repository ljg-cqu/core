package _errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMyError(t *testing.T) {
	err2 := fmt.Errorf("this is error 2:%w", errors.New("This is error 1"))

	err3 := fmt.Errorf("this is err3:%w", err2)

	myErr4 := New().
		WithMsg("This is error 4").
		Wrap(err3).
		WithErrType(ErrTypeParseRSAKey).
		WithWhen(time.Now().String()).
		WithTag(ErrTagFilePathErr).
		WithTag(ErrTagAuthenExpired).
		WithTag(ErrorTag("self_defined_tag")).
		WithFiled("file_integrity_problem", "file has been broken")

	myErr5 := New().
		WithMsg("This is error 5").
		Wrap(myErr4)

	fmt.Println(myErr5)
	//fmt.Println(myErr5.ErrorFall())
	require.Equal(t, true, myErr4.As(ErrTypeParseRSAKey))
	require.Equal(t, false, myErr5.Is(myErr4))
	require.Equal(t, true, myErr4.Is(myErr5.Unwrap().(*MyError)))
	require.Equal(t, true, myErr5.Unwrap().(*MyError).Is(myErr4))
	require.Equal(t, true, myErr4.TagExists(ErrTagFilePathErr))

	bytes, _ := json.MarshalIndent(myErr5, "", "  ")
	fmt.Println(string(bytes))
}
