// The best way to send emails in Go with SMTP Keep Alive and Timeout for Connect and Send.

package smtp

import (
	"crypto/tls"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"errors"
	"time"
)

const (
	connectTimeout = 60 * time.Second
	sendTimeout    = 60 * time.Second
	keepAlive      = false // keep it alive for sending multiple emails
	helo           = "localhost"

	encryptionType = mail.EncryptionSTARTTLS
	authType       = mail.AuthPlain
)

const (
	NetEase126Mail Provider = iota
	NetEase163Mail
	QQMail
)

type Provider int

var providerConfigs = map[Provider]providerConfig{
	NetEase126Mail: {"smtp.126.com", 25},
	NetEase163Mail: {"smtp.163.com", 25}, // SSL: 465
	QQMail:         {"smtp.qq.com", 587}, // TODO: or port as 465
}

type providerConfig struct {
	host string
	port int
}

type EmailClient struct {
	client *mail.SMTPClient
}

func NewEmailClient(prvd Provider, tlsConfig *tls.Config, userName, password string) (*EmailClient,error) {
	srv := mail.NewSMTPClient()

	srv.Host = providerConfigs[prvd].host
	srv.Port = providerConfigs[prvd].port

	srv.Username = userName
	srv.Password = password

	srv.Encryption = encryptionType // TODO: defaults to EncryptionNone?
	srv.Authentication = authType
	srv.TLSConfig = tlsConfig

	srv.KeepAlive = keepAlive
	srv.ConnectTimeout = connectTimeout
	srv.SendTimeout = sendTimeout
	srv.Helo = helo

	cli, err := srv.Connect()
	if err != nil || cli.Noop() != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect stmp server %q, error: %v", srv.Host, err))
	}

	return &EmailClient{cli},nil
}

func (e *EmailClient) Send(msg *mail.Email) error {
	return msg.Send(e.client)
}
