package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"github.com/ljg-cqu/core/_errors"
)

type RsaError struct {
}

func (e *RsaError) Error() string {
	return ""
}

// Only small messages can be signed directly; thus the hash of a
// message, rather than the message itself, is signed. This requires
// that the hash function be collision resistant. SHA-256 is the most
// currently wide-adopted hash algorithm for secure applications

func SignRS256FromFileWithKeyPath(keyPath, signingPath string) ([]byte, _errors.Error) {
	privKey, err := ParseRSAPrivateKeyFromFile(keyPath)
	if err != nil {
		return nil, _errors.New().Wrap(err).WithWhat("failed to sign RS256 from file with key bytes")
	}

	return SignRS256FromFile(privKey, signingPath)
}

func SignRS256FromFileWithKeyBytes(key []byte, signingPath string) ([]byte, _errors.Error) {
	privKey, err := ParseRSAPrivateKey(key)
	if err != nil {
		return nil, _errors.New().Wrap(err).WithWhat("failed to sign RS256 from file with key bytes")
	}

	return SignRS256FromFile(privKey, signingPath)
}

func SignRS256FromFile(key *rsa.PrivateKey, signingPath string) ([]byte, _errors.Error) {
	digest, err := Sum256FromFile(signingPath)
	if err != nil {
		return nil, _errors.New().Wrap(err).WithWhat("failed to sign RS256 from file")
	}

	return SignRS256WithDigest(key, digest)
}

func SignRS256(key *rsa.PrivateKey, msg []byte) ([]byte, _errors.Error) {
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, Sum256(msg))
	if err != nil {
		return nil, _errors.New().Wrap(err).WithWhat("failed to sign RS256")
	}
	return signature, nil
}

func SignRS256WithDigest(key *rsa.PrivateKey, hashed []byte) ([]byte, _errors.Error) {
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		return nil, _errors.New().Wrap(err).WithWhat("failed to sign RS256")
	}
	return signature, nil
}
