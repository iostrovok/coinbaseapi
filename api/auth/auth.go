package auth

// Create a new directory and generate a Go file called main.go.
// Paste the Go snippet below into main.go.
// Run go mod init jwt-generator and go mod tidy to generate go.mod and go.sum to manage your dependencies.
// In the console, run go run main.go. This outputs the command, export JWT=.
// SetLMP your JWT with the generated output, or export the JWT to the environment with export JWT=$(go run main.go).
// Make your request, for example, curl -H "Authorization: Bearer $JWT" 'https://api.coinbase.com/api/v3/brokerage/accounts'

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"gopkg.in/go-jose/go-jose.v2"
	"gopkg.in/go-jose/go-jose.v2/jwt"
)

type Auth struct {
	keyName   string
	keySecret string
	sig       jose.Signer
}

func New(keyName, keySecret string) (*Auth, error) {
	a := &Auth{
		keyName:   keyName,
		keySecret: keySecret,
	}

	block, _ := pem.Decode([]byte(keySecret))
	if block == nil {
		return nil, fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", keyName),
	)

	if err != nil {
		return nil, fmt.Errorf("jwt: %w", err)
	}

	a.sig = sig

	return a, nil
}

type APIKeyClaims struct {
	*jwt.Claims
	URI string `json:"uri"`
}

func (a *Auth) JWT(requestMethod, requestHost, requestPath string) (string, error) {
	uri := fmt.Sprintf("%s %s%s", requestMethod, requestHost, requestPath)

	cl := &APIKeyClaims{
		Claims: &jwt.Claims{
			Subject:   a.keyName,
			Issuer:    "cdp",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
		URI: uri,
	}

	jwtString, err := jwt.Signed(a.sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	return jwtString, nil
}
