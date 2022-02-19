package utils

import (
	"crypto/sha256"
	"io"
	"os"
)

func Sum256FromString(content string) []byte {
	h := sha256.New()
	h.Write([]byte(content))

	return h.Sum(nil)
}

func Sum256FromBytes(content []byte) []byte {
	h := sha256.New()
	h.Write(content)

	return h.Sum(nil)
}

func Sum256FromFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
