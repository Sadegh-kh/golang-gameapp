package config

import "time"

const (
	AccessTokenDuration      = time.Hour * 24
	RefreshTokenDuration     = time.Hour * 24 * 7
	MiddlewareAuthJWTContext = "user"
)
