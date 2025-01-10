package utils

import (
	"ampl/src/config"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsRelease() bool {
	fmt.Println(config.Config.Env)
	return strings.EqualFold(config.Config.Env, PROD)
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	var err error
	privateKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("parsing failed")
	}

	ref, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, errors.New("parsing failed")
	}
	return ref, nil
}

func JwtEncode(name string, exp int64, key *rsa.PrivateKey) (string, error) {
	currTime := time.Now().UnixMilli()
	claims := jwt.MapClaims{
		JWT_NAME: name,
		JWT_EXP:  exp,
		JWT_IAT:  currTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
