package config

import (
	"github.com/ljg-cqu/core/openchain"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestMarshalOpenChainHttpClientConfig(t *testing.T) {
	c := &openchain.Config{
		BizId: "a00e36c5",

		HandshakeUrl: "https://rest.baas.alipay.com/api/contract/shakeHand",
		TransactUrl:  "https://rest.baas.alipay.com/api/contract/chainCallForBiz",
		QueryUrl:     "https://rest.baas.alipay.com/api/contract/chainCall",

		Account:  "Zealy",
		TenantId: "KHPEGMYY",
		KmsKeyId: "X11L2LN0KHPEGMYY1638414736161",

		AccessId:      "01FHXjyeKHPEGMYY",
		AccessKeyPath: "../misc/test_priv_key_for_rsa.key",

		RetryMaxAttempts: 5,
		RetryInSeconds:   5,

		MaxIdleConns:             10,
		IdleConnTimeoutInSeconds: 30,

		RequestTimeoutInSeconds: 10,

		TokenTimeoutInMinutes: 30,

		GasLimit: 100000,
	}

	c_ := &openchain.HttpClientConfig{
		Config: c,
	}

	utils.PrintlnAsJson("openchain config", c_)
}

func TestReadOpenChainHttpClientConfig(t *testing.T) {
	var c openchain.HttpClientConfig
	bytes, err := ioutil.ReadFile("./config.json")
	require.Nil(t, err)
	err = utils.Json.Unmarshal(bytes, &c)
	require.Nil(t, err)
	utils.PrintlnAsJson("read config: ", c)
}
