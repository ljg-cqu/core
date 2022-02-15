// The best way to send emails in Go with SMTP Keep Alive and Timeout for Connect and Send.

package smtp

import (
	"crypto/tls"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"os"
	"time"
)

const (
	connectTimeout = 10 * time.Second
	sendTimeout    = 10 * time.Second
	keepAlive      = false // keep it alive for sending multiple emails
	helo           = "localhost"

	encryptionType = mail.EncryptionSTARTTLS
	authType       = mail.AuthPlain
)

const (
	NetEase126Mail Provider = iota
	QQMail
)

type Provider int

var providerConfigs = map[Provider]providerConfig{
	NetEase126Mail: {"smtp.126.com", 25},
	QQMail:         {"smtp.qq.com", 465}, // TODO: or port as 587
}

type providerConfig struct {
	host string
	port int
}

type EmailClient struct {
	client *mail.SMTPClient
}

func NewEmailClient(prvd Provider, tlsConfig *tls.Config, userName, password string) *EmailClient {
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
		fmt.Printf("Failed to connect stmp server %q, error: %v", srv.Host, err)
		os.Exit(1)
	}

	return &EmailClient{cli}
}

func (e *EmailClient) Send(msg *mail.Email) error {
	return msg.Send(e.client)
}
