package openchain

import (
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/utils"
	"os"
	"time"
)

var client *Client

type Client struct {
	HandshakeNum int

	CurrentToken  string
	TokenExpireAt time.Time

	*resty.Client
	resty.Logger
	*Config
}

type Option func(c *Client)

// TODO: more configs on resty.Client, including SetLogger

// New returns a wrapper of resty.Client with given configs.
// Note: config resty.Client to suit your special needs, such as Client.EnableTrace().
func New(log resty.Logger, config string) *Client {
	if client != nil {
		return client
	}

	var c HttpClientConfig
	if err := utils.UnmarshalJsonFile(config, &c); err != nil {
		println("failed to create open chain http client", err)
		os.Exit(1)
	}

	key, err := utils.ParseRSAPrivateKeyFromFile(c.Config.AccessKeyPath)
	if err != nil {
		println("failed to create open chain http client", err)
		os.Exit(1)
	}

	//logrus.New()
	cli := &Client{}

	cli.Logger = log
	cli.Config = c.Config
	cli.AccessKey = key

	cli.TokenTimeoutInMinutes_ = time.Minute * time.Duration(cli.TokenTimeoutInMinutes)
	cli.IdleConnTimeoutInSeconds_ = time.Second * time.Duration(cli.IdleConnTimeoutInSeconds)
	cli.RetryInSeconds_ = time.Second * time.Duration(cli.RetryInSeconds)
	cli.RequestTimeoutInSeconds_ = time.Second * time.Duration(cli.RequestTimeoutInSeconds)

	cli.Client = resty.New()

	cli.SetLogger(log)
	cli.SetTimeout(cli.RequestTimeoutInSeconds_)
	cli.SetRetryWaitTime(cli.RetryInSeconds_)
	cli.SetRetryMaxWaitTime(cli.RetryInSeconds_ * 2)
	cli.SetRetryCount(cli.RetryMaxAttempts)

	if err := cli.HandshakeIfNeeded(); err != nil {
		println("failed to handshake open chain:", err)
		os.Exit(1)
	}

	cli.OnBeforeRequest(func(c *resty.Client, request *resty.Request) error {
		return cli.HandshakeIfNeeded()
	})

	return cli
}
