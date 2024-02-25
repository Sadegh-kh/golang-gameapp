package config

import (
	"gameapp/service/authservice"
	"gameapp/storage/mysql"
)

type HttpConfig struct {
	Port int
}
type Config struct {
	HttpConf HttpConfig
	Auth     authservice.Config
	MySQL    mysql.Config
}
