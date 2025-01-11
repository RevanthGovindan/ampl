package config

import (
	"ampl/src/models"
	"crypto/rsa"

	"github.com/go-playground/validator/v10"
)

var (
	Config          models.Config = models.Config{}
	CloudPrivateKey *rsa.PrivateKey
	CloudPublicKey  *rsa.PublicKey
	Validate        = validator.New()
)
