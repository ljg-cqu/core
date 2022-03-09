package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/ljg-cqu/core/errors"
	"io/ioutil"
	"os"
)

func ParseRSAPrivateKeyFromFile(keyPath string) (*rsa.PrivateKey, errors.Error) {
	f, err := os.Open(keyPath)
	if err != nil {
		return nil, errors.New().Wrap(err).
			WithWhat("failed to open file").
			WithTag(errors.ErrTagFilePathErr)
	}
	defer f.Close()

	input, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New().Wrap(err).
			WithTag(errors.ErrTagFileReadErr)
	}

	return ParseRSAPrivateKey(input)
}

func ParseRSAPrivateKey(key []byte) (*rsa.PrivateKey, errors.Error) {
	block, _ := pem.Decode(key)
	if block == nil {
		err := errors.New().
			WithErrType(errors.ErrTypeParseRSAKey).
			WithWhat("failed to parse private key").
			WithWhy("Key must be a PEM encoded PKCS1 or PKCS8 key").
			WithTag(errors.ErrTagCryptoInvalidPrivateKey)
		return nil, err
	}

	var parsedKey any
	var err error
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, errors.New().Wrap(err).
				WithErrType(errors.ErrTypeParseRSAKey).
				WithWhat("failed to parse private key").
				WithTag(errors.ErrTagCryptoInvalidPrivateKey)
		}
	}

	pkey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New().
			WithErrType(errors.ErrTypeParseRSAKey).
			WithWhat("failed to parse private key").
			WithWhy("key is not a valid RSA private key")
	}

	return pkey, nil
}
