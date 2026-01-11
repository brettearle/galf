package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const (
	KeyPrefix  = "GLF"
	KeyVersion = "V1"
)

type APIKeyGenerator struct {
	prefix  string
	version string
}

func NewAPIKeyGenerator() *APIKeyGenerator {
	return &APIKeyGenerator{
		prefix:  KeyPrefix,
		version: KeyVersion,
	}
}

func (ak *APIKeyGenerator) Generate() (fullKey string, kid string, err error) {
	keyLength := 32
	randBytes := make([]byte, keyLength)
	if _, err := rand.Read(randBytes); err != nil {
		return "", "", fmt.Errorf("failed to generate key")
	}

	cs := sha256.Sum256(randBytes)
	trunc := base64.RawURLEncoding.EncodeToString(cs[:])[:8]

	full := fmt.Sprintf("%s_%s_%s", ak.prefix, ak.version, randBytes)

	return full, trunc, nil
}
