package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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

	err := UnmarshalJsonFile("../test/config.json", &c)
	require.NotNil(t, err)

	err = UnmarshalJsonFile("../tests/config.json", &c)
	require.Nil(t, err)

	PrintlnAsJson("read config: ", &c)
}
