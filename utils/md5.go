package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"time"
)

func MD5NowID() string {
	return MD5Now()[:6]
}

func MD5Now() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5B64(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
