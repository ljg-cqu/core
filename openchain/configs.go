package openchain

import (
	"crypto/rsa"
	"time"
)

type HttpClientConfig struct {
	Config *Config `json:"OpenChainHttpClient"`
}

type Config struct {
	BizId string `json:"bizid"`

	HandshakeUrl string `json:"HandshakeUrl"`
	TransactUrl  string `json:"TransactUrl"`
	QueryUrl     string `json:"QueryUrl"`

	Account  string `json:"account"`
	TenantId string `json:"tenantid"`
	KmsKeyId string `json:"mykmsKeyId"`

	AccessId      string          `json:"AccessId"`
	AccessKeyPath string          `json:"AccessKeyPath"`
	AccessKey     *rsa.PrivateKey `json:"-"` // parsed from AccessKeyPath

	RetryMaxAttempts int           `json:"RetryMaxAttempts"`
	RetryInSeconds   int           `json:"RetryInSeconds"`
	RetryInSeconds_  time.Duration `json:"-"` // calculated from RetryInSeconds

	MaxIdleConns              int           `json:"MaxIdleConns"`
	IdleConnTimeoutInSeconds  int           `json:"IdleConnTimeoutInSeconds"`
	IdleConnTimeoutInSeconds_ time.Duration `json:"-"` // calculated from IdleConnTimeoutInSeconds

	RequestTimeoutInSeconds  int           `json:"RequestTimeoutInSeconds"`
	RequestTimeoutInSeconds_ time.Duration `json:"-"`

	TokenTimeoutInMinutes  int           `json:"TokenTimeoutInMinutes"`
	TokenTimeoutInMinutes_ time.Duration `json:"-"` // calculated from TokenTimeoutInMinutes

	GasLimit int64 `json:"GasLimit"`
}
