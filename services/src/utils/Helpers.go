package utils

import (
	"ampl/src/config"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsRelease() bool {
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

func LoadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	var err error = nil
	var pemKey []byte
	pemKey, err = os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	ref, err := jwt.ParseRSAPublicKeyFromPEM(pemKey)
	if err != nil {
		return ref, err
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

func JwtDecrypt(encryptedToken string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	decryptedToken, err := jwt.Parse(encryptedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := decryptedToken.Claims.(jwt.MapClaims)
	if !ok || !decryptedToken.Valid {
		return nil, errors.New("invalid encryption")
	}

	return claims, nil
}

func FindProjectRoot() (string, error) {
	wd, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, err
		}
		wd = filepath.Dir(wd)
		if wd == "/" {
			break
		}
	}
	return "", nil
}
