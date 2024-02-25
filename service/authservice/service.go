package authservice

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Config struct {
	SecretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	AccessSubject        string
	RefreshSubject       string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}

}

func (s Service) CreateAccessToken(uid uint) (string, error) {
	return createToken(uid, s.config.AccessSubject, s.config.SecretKey, s.config.AccessTokenDuration)
}
func (s Service) CreateRefreshToken(uid uint) (string, error) {
	return createToken(uid, s.config.RefreshSubject, s.config.SecretKey, s.config.RefreshTokenDuration)
}

func (s Service) ParseToken(token string) (uint, error) {
	var userClaim Claims
	_, err := jwt.ParseWithClaims(token, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}
	return userClaim.UserID, nil
}

// Valid implements jwt.Claims.
func (c *Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}

func createToken(uid uint, subject, signKey string, duration time.Duration) (string, error) {
	// TODO - create jwt by RS256 algorithm
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	// set our claims
	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// set the expiry time
			// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Subject:   subject,
		},
		UserID: uid,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, t.Claims)
	tokenString, err := accessToken.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}
	// Creat token string
	return tokenString, nil
}
