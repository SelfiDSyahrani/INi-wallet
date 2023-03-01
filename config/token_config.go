package config

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfigLagi struct {
	TokenConfig
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

func (c ConfigLagi) readConfig() ConfigLagi {
	c.TokenConfig = TokenConfig{
		ApplicationName:     "ENIGMA",
		JwtSignatureKey:     "password",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: time.Second * 60,
	}
	return c

}

func NewConfigLagi() ConfigLagi {
	cfg := ConfigLagi{}
	return cfg.readConfig()
}
