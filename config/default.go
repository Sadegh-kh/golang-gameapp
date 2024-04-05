package config

var defaultConfig = map[string]interface{}{
	"http.port":                   8999,
	"auth.access_token_duration":  AccessTokenDuration,
	"auth.refresh_token_duration": RefreshTokenDuration,
	"auth.access_subject":         "ac",
	"auth.refresh_subject":        "rf",
}
