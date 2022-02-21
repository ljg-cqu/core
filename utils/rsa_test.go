package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignRS256FromFileWithKeyPath(t *testing.T) {
	var keyPath, signingPath = "./misc/test_priv_key_for_rsa.key", "./rsa.go"

	signature, err := SignRS256FromFileWithKeyPath(keyPath, signingPath)
	require.Nil(t, err)
	fmt.Println(string(signature))
}

func TestSignRS256FromFileWithKeyBytes(t *testing.T) {
	const privKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtPQToFonF60OX
9TmEbM9Nj8fNVeItpMfZNnFooNcNd9rWmJ6ICWNaV1Qs/kNzpigxGjPBtkZ+KH2q
7detqDVIcpuisRHu1hJJBcr1/A7pGIf1fXTsIiBXn4eoixA3wMupW7vhlthniOow
iOm+7LP6pey/WX9lhw8W3S9BnnjsYAtmpXLbTZ10bYBKoOFpC2GLSzqXH+PvpzAw
fvNkJZD+JW5cY2lUV7ec+fGFwVFptajDWih8yKa6XaAWEhPiPMw6mcvNE3DQ55Kc
tlPIzeAXykiQfeItxv9DTxIDzUBxNbsRneZBH5WOLrOhAOpbE9arg9D6fn0expqI
kd5KwA5bAgMBAAECggEADjOU3ePPGor9SQ1A0FLNMboKMpKKTpyWB3/3jxC0YHXF
Wlc7k8JVQzgqfd/ALtBdthzERmqHX9s45hTGXAWQjKZcjNtAMZiZ+iN/7mdh34jz
yFOnDJ6FkTlSOSZhR3jGGVWcUtN3XRFzxVPL+atU28TTYiJXl76ZJZIvSA1SM2pP
ZGyjkGZtic3hV1VSIKj/H9//AHkvPPlpL2aKnm0SIFwWS4neUnbSeo3cUkvay6A8
fijrDsWKwZ0rayjSNiK3dVE3Cp5a0LlCsRjGVduMPHcpOBHkjxDJsg2ZetXabZIX
lmOlciuVFO+aalpPg1I0rzKn5oOEqwKeIPXe61g3FQKBgQDvioCGrBKFXnKa5uOB
Rc6KhfSvYtB1JYadZ2qYRyQKeZC5ubpFEzbqKgNeYnFIxHcxi92Krbm1+h3Emytz
daTbFu1iT0EVrPb06uc92ygYROYs7sraEWgjkBP/VEmARh4k3+6GWDXGPZDNbes0
v3sMzOk0Ikloc71yqsV6PMnf1wKBgQC5JELY7pgJ2evIt8loty9psaKNLuq887Hf
6oThcL3ubBQdUQ1bOguTwDASC+Jc3i8wLOmnPAeBdPEpZWY3QzZ+OsaH8qQLFaQq
XUeAKGfa1eXcuPiYOv42H6sSqjWp2Q/DuYAdFt7WmsBhkVNsMlPQl8GDp0ZMLdht
SfC0t0GFHQKBgQC0X4GaZxXnMZBwze2AGVWGf4oZSvoXTDOKcSYWFnOwI4v0HkOB
4g8W0p4Iw230UmRCfcRLubc+rWEe+40DexGxHBmSToV+0eh/0iZgMJeHdtIwAXvk
KvlU1hgIyqoyGhp2v9x3cxLC/Pb9iYh0Br+ciuwLosnOCmEcaDUdb4q/rwKBgQCG
W1JHq56aR5NcrkNzwrydr1OPsaSYSyGipcaY9ABhrf1K6S8QLSeJqcc40XcMfhEw
nOdTfbTUtdDtgbCUGirJoE3DCssRYDsqo1boImp73Q5bB8EgeG9TR9gWS392KxfN
qijW82nzw4opRBWOR1eb4QWGTTYuwnZ1mVsdSoA54QKBgGYOyMCdepxd3UosDyb0
T8v6KpN3TevBUkABqwtYYMHWiZCdGwSYkC7PmPiFzDUZaunUfuIlcBY6j1S+fIOb
9k63z6LkgHSnwXkJ6rvJOO1LhnV4gxXHGoDkXgy9CPfxvlmsyD2ns+sqLtK4v/9q
QQfusFm96g3AJKg0mE0F7oho
-----END PRIVATE KEY-----`
	signature, err := SignRS256FromFileWithKeyBytes([]byte(privKey), "./rsa.go")
	require.Nil(t, err)
	fmt.Println(string(signature))
}

func TestSignRS256FromFile(t *testing.T) {
	privKey, err := ParseRSAPrivateKeyFromFile("../misc/test_priv_key_for_rsa.key")
	require.Nil(t, err)

	signature, err := SignRS256FromFile(privKey, "./rsa.go")
	require.Nil(t, err)
	fmt.Println(string(signature))
}

func TestSignRS256WithDigest(t *testing.T) {
	var h = "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447"

	privKey, err := ParseRSAPrivateKeyFromFile("../misc/test_priv_key_for_rsa.key")
	require.Nil(t, err)

	signature, err := SignRS256WithDigest(privKey, []byte(h))
	fmt.Println(err)

	require.Nil(t, err)
	fmt.Println(err)
	fmt.Println(string(signature))
}

func TestSignRS256(t *testing.T) {
	var raw = "hello world"

	privKey, err := ParseRSAPrivateKeyFromFile("../misc/test_priv_key_for_rsa.key")
	require.Nil(t, err)

	signature, err := SignRS256(privKey, []byte(raw))
	fmt.Println(err)

	require.Nil(t, err)
	fmt.Println(err)
	fmt.Println(string(signature))
}
