package authservice

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Service struct {
	secretKey            string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	accessSubject        string
	refreshSubject       string
}

func New(secretKey, accessSubject, refreshSubject string, accessDuration, refreshDuration time.Duration) Service {
	return Service{
		secretKey:            secretKey,
		accessTokenDuration:  accessDuration,
		refreshTokenDuration: refreshDuration,
		accessSubject:        accessSubject,
		refreshSubject:       refreshSubject,
	}

}

func (s Service) CreateAccessToken(uid uint) (string, error) {
	return createToken(uid, s.accessSubject, s.secretKey, s.accessTokenDuration)
}
func (s Service) CreateRefreshToken(uid uint) (string, error) {
	return createToken(uid, s.refreshSubject, s.secretKey, s.refreshTokenDuration)
}

func (s Service) ParseToken(token string) (uint, error) {
	var userClaim Claims
	_, err := jwt.ParseWithClaims(token, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
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
