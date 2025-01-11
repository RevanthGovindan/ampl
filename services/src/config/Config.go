package config

import (
	"ampl/src/models"
	"crypto/rsa"

	"github.com/go-playground/validator/v10"
)

var (
	Config           models.Config = models.Config{}
	JwtRsaPrivateKey *rsa.PrivateKey
	JwtRsaPublicKey  *rsa.PublicKey
	Validate         = validator.New()
)
