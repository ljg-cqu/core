package openchain

import (
	"time"
)

type Identity string

// Message represents type for communication with open chain.
// Guidance: https://antchain.antgroup.com/docs/11/146925?Source=brand_baidu_my_13627.
type Message struct {
	RawContent    string
	ContentToSign []byte // generated with RawContent, AccessId, Timestamp

	Time      time.Time
	Timestamp string // corresponding to field Time

	Account string

	// Identity represents identity on open chain, as defined on
	// https://antchain.antgroup.com/docs/11/143753?Source=brand_baidu_my_13627.
	ClientIdentity Identity // identity that directly communicate with ant open chain, especially from open chain client.
	UserIdentity   Identity // identity that cause this message to be sent to open chain, especially from end user.
}

//
//func NewMessage(rawContent string) *Message {
//	now := time.Now()
//	nowStamp := strconv.FormatInt(now.UnixMilli(), 16)
//
//	contentForSign := rawContent + "-" + AccessId + nowStamp
//
//	return &Message{
//		RawContent:    rawContent,
//		ContentToSign: []byte(contentForSign,
//		Time:          now,
//		Timestamp:     nowStamp,
//	}
//}

func (m *Message) SigningContent() []byte {
	return m.ContentToSign
}
