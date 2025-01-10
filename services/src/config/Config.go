package config

import (
	"ampl/src/models"
	"crypto/rsa"
)

var (
	Config          models.Config = models.Config{}
	CloudPrivateKey *rsa.PrivateKey
	CloudPublicKey  *rsa.PublicKey
)
