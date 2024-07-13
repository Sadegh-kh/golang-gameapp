package config

import (
	"fmt"
	"gameapp/service/authservice"
	"gameapp/storage/mysql"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"strings"
)

type HttpConfig struct {
	Port int `koanf:"port"`
}
type Config struct {
	HttpConf HttpConfig         `koanf:"http"`
	Auth     authservice.Config `koanf:"auth"`
	MySQL    mysql.Config       `koanf:"mysql"`
}

func Load(cfgPath string) *Config {
	k := koanf.New(".")

	// Load default values using the confmap provider.
	// We provide a flat map with the "." delimiter.
	// A nested map can be loaded by setting the delimiter to an empty string "".
	k.Load(confmap.Provider(defaultConfig, "."), nil)

	// Load YAML config and merge into the previously loaded config (because we can).
	err := k.Load(file.Provider(cfgPath), yaml.Parser())
	if err != nil {
		panic(fmt.Sprintln("when load config file error happened:", err))
	}

	k.Load(env.Provider("GAMEAPP_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "GAMEAPP_")), "_", ".", -1)
	}), nil)

	cfg := new(Config)

	err = k.Unmarshal("", cfg)
	if err != nil {
		panic(fmt.Sprintln("when unmarshalling config error happened:", err))
	}

	return cfg

}
