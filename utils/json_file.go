package utils

import (
	"github.com/ljg-cqu/core/errors"
	"io/ioutil"
)

func UnmarshalJsonFile(path string, a any) errors.Error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.NewWithMsgf("failed to read file:%v", err).WithTag(errors.ErrTagFileReadErr)
	}
	err = Json.Unmarshal(bytes, a)
	if err != nil {
		return errors.NewWithMsgf("failed to unmarshal file:%v", err).WithTag(errors.ErrTagJsonUnmarshalErr)
	}

	return nil
}
