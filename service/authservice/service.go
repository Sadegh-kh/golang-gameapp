package authservice

import (
	"gameapp/pkg/richerror"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	// do not use _ in meta tag koanf for secret key.
	// because we get that variable from local env variable
	SecretKey string `koanf:"secretKey"`

	AccessTokenDuration  time.Duration `koanf:"access_token_duration"`
	RefreshTokenDuration time.Duration `koanf:"refresh_token_duration"`
	AccessSubject        string        `koanf:"access_subject"`
	RefreshSubject       string        `koanf:"refresh_subject"`
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

func (s Service) ParseToken(token string) (Claims, error) {
	var userClaim Claims
	_, err := jwt.ParseWithClaims(token, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		return Claims{}, richerror.RichError{
			Operation:    "authentication.ParseToken",
			WrappedError: err,
			Message:      "invalid token",
			Kind:         richerror.Invalid,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}
	return userClaim, nil
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
