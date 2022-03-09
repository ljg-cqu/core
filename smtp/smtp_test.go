package smtp

import (
	"crypto/tls"
	"github.com/stretchr/testify/require"
	mail "github.com/xhit/go-simple-mail/v2"

	"testing"
)

func TestNewNetEaseSMTPClient(t *testing.T) {
	cli := NewEmailClient(NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "UPVZLTEGBSNPGXCI")
	require.NotNil(t, cli)
}

func TestNetEaseSMTPClient_Send(t *testing.T) {
	cli := NewEmailClient(NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "UPVZLTEGBSNPGXCI")
	require.NotNil(t, cli)

	const htmlBody = `<html>
	<head>
		<meta http-equiv="RawContent-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>，
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

	email := mail.NewMSG()
	email.SetFrom("Zealy <ljg_cqu@126.com>").
		AddTo("ljg_cqu@126.com", "luo.jigao@21vianet.com").
		AddCc("qq1025003548@gmail.com", "641277312@qq.com").
		SetSubject("Another new ABFPaaS Email for tests purpose")

	email.SetBody(mail.TextHTML, htmlBody)

	require.Nil(t, cli.Send(email))
}
