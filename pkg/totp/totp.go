package totp

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// TOTPGenerator is an interface for generating TOTP keys and codes
type TOTPGenerator interface {
	GenerateOTP(accountName string) (*otp.Key, error)
	GenerateTotpCode(secret string) (string, error)
	ValidateOTP(code, secret string) bool
}

type totpGenerator struct {
	Issuer string
}

func NewTOTPGenerator(Issuer string) TOTPGenerator {
	return &totpGenerator{
		Issuer: Issuer,
	}
}

func (t *totpGenerator) GenerateOTP(accountName string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      t.Issuer,
		AccountName: accountName,
	})
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (t *totpGenerator) GenerateTotpCode(secret string) (string, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return code, nil
}

func (t *totpGenerator) ValidateOTP(code, secret string) bool {
	return totp.Validate(code, secret)
}
