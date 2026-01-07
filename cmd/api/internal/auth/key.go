package auth

import (
	"context"
	"fmt"
)

const (
	PrefixLive Prefix = "live"
	PrefixTest Prefix = "test"
)

var ValidPrefix map[Prefix]bool = map[Prefix]bool{
	PrefixLive: true,
	PrefixTest: true,
}

type Store interface {
	GetByKid(ctx context.Context, kid Kid) (pf Prefix, version string, scopes []string, revoked bool, verifier []byte, err error)
	CreateKey(ctx context.Context, pf Prefix, kid Kid, version string, verifier []byte, scopes []string, revoked bool) error
}

type Kid [16]byte
type Prefix string

type ParsedKey struct {
	Prefix  Prefix
	Kid     Kid
	secret  []byte // raw
	Version string
}

type VerifiedKey struct {
	Prefix  Prefix
	Kid     Kid
	Version string
	Scopes  []string
}

func NewKey(ctx context.Context, prefix string, masterSecret []byte, store Store, scopes []string) (string, error) {
	// 1) validate prefix is allowed
	valid := ValidatePrefix(prefix)
	if valid != true {
		return "", fmt.Errorf("Invalid Prefix")
	}

	// 2) generate kid (random)
	// 3) generate clientSecretPart (random)  // this is what the client will present later

	// 4) verifier = HMAC(masterSecret, []byte (msg = "v1" + 0x00 + prefix + 0x00 + kid + 0x00 + clientSecretPart)

	// 5) store row:
	//    (kid, version, prefix, verifier, created_at, revoked=false, scopes...)

	// 6) return the API key string (shown once):
	//    "<prefix>_<version>_<encodedKid>_<encodedClientSecretPart>"
	return "", nil
}

func ValidatePrefix(s string) bool {
	_, valid := ValidPrefix[Prefix(s)]
	return valid
}

// returns the encoded key string you store/show the user ONCE

func ParseKey(raw string) (ParsedKey, error) {
	return ParsedKey{}, fmt.Errorf("NOT IMPLEMENTED")
}

// parses/decodes the string into parts

func VerifyKey(ctx context.Context, raw string, masterSecret []byte, store Store) (VerifiedKey, error) {
	return VerifiedKey{}, fmt.Errorf("NOT IMPLEMENTED")
}

// DB version, prefix are truth
// Parse + MAC verify (timing-safe)

func (k Kid) String() string {
	return "NOT IMPLEMENTED"
} /* base32/hex encode */
func parseKid(s string) (Kid, error) {
	notImp := "NOT IMPLEMENTED"
	return Kid([]byte(notImp)), fmt.Errorf("NOT IMPLEMENTED")
} /* decode */
