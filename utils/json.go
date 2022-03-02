package utils

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ljg-cqu/core/_errors"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func PrintlnAsJson(pre string, v any) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(pre)
	fmt.Println(string(bytes))
}

func UnmarshalJsonFile(path string, a any) _errors.Error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return _errors.NewWithMsgf("failed to read file:%v", err).WithTag(_errors.ErrTagFileReadErr)
	}
	err = Json.Unmarshal(bytes, a)
	if err != nil {
		return _errors.NewWithMsgf("failed to unmarshal file:%v", err).WithTag(_errors.ErrTagJsonUnmarshalErr)
	}

	return nil
}

func TestUnmarshalJsonFile(t *testing.T) {
	type Config struct {
		HandshakeUrl string `json:"HandshakeUrl"`
		TransactUrl  string `json:"TransactUrl"`
		QueryUrl     string `json:"QueryUrl"`

		AccessId      string `json:"AccessId"`
		AccessKeyPath string `json:"AccessKeyPath"`

		RetryMaxAttempts int `json:"RetryMaxAttempts"`
		RetryInSeconds   int `json:"RetryInSeconds"`

		MaxIdleConns             int `json:"MaxIdleConns"`
		IdleConnTimeoutInSeconds int `json:"IdleConnTimeoutInSeconds"`

		TokenTimeoutInminutes int `json:"TokenTimeoutInMinutes"`
	}

	type HttpClientConfig struct {
		Config *Config `json:"OpenChainHttpClient"`
	}

	var c HttpClientConfig

	err := UnmarshalJsonFile("../test/configs.json", &c)
	require.NotNil(t, err)

	err = UnmarshalJsonFile("../tests/configs.json", &c)
	require.Nil(t, err)

	PrintlnAsJson("read config: ", &c)
}
