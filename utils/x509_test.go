package utils

import (
	"github.com/ljg-cqu/core/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseRSAPrivateKeyFromFile2(t *testing.T) {
	var invalidKeyPath = "err path"
	_, err := ParseRSAPrivateKeyFromFile(invalidKeyPath)
	require.NotNil(t, err)
	require.Equal(t, true, err.TagExists(errors.ErrTagFilePathErr))

	var validKeyPath = "../misc/test_priv_key_for_rsa.key"
	rsaPrivKey, err := ParseRSAPrivateKeyFromFile(validKeyPath)
	require.Nil(t, err)
	PrintlnAsJson("parsed RSA: ", rsaPrivKey)
}

func TestParseRSAPrivateKey(t *testing.T) {
	var invalidInput = []byte(`invalid iinput key`)
	_, err := ParseRSAPrivateKey(invalidInput)
	require.NotNil(t, err)
	require.Equal(t, errors.ErrTypeParseRSAKey, err.GetErrType())

	const pubPEM = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
-----END PUBLIC KEY-----`

	_, err = ParseRSAPrivateKey([]byte(pubPEM))
	require.NotNil(t, err)
	require.Equal(t, errors.ErrTypeParseRSAKey, err.GetErrType())

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

	rsaPrivKey, err := ParseRSAPrivateKey([]byte(privKey))
	require.Nil(t, err)
	PrintlnAsJson("parsed RSA: ", rsaPrivKey)
}
