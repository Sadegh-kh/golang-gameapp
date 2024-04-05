package middleware

import (
	"gameapp/pkg/constant"
	"gameapp/service/authservice"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(svc authservice.Service, cfg authservice.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey: constant.MiddlewareAuthJWTContext,
		SigningKey: cfg.SecretKey,

		//TODO - add signing method to auth config
		SigningMethod: "HS256",

		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claim, err := svc.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claim, nil
		},
	})
}
