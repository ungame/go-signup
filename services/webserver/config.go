package webserver

import (
	"encoding/json"
	"flag"
	"github.com/ungame/go-signup/grpcext"
	"github.com/ungame/go-signup/httpext"
	"github.com/ungame/go-signup/services/authentication"
)

const DefaultPort = 8080

var serviceConfig Config

type Config struct {
	Port        httpext.Port   `json:"port"`
	AuthService grpcext.Config `json:"auth_service"`
}

func (c *Config) MarshalText() ([]byte, error) {
	return json.Marshal(*c)
}

func (c *Config) UnmarshalText(b []byte) error {
	return json.Unmarshal(b, c)
}

func LoadJSONConfigFromFlags(flagSet *flag.FlagSet) {
	flagSet.TextVar(&serviceConfig, "webserver", &Config{
		Port: DefaultPort,
		AuthService: grpcext.Config{
			Host: httpext.Localhost,
			Port: authentication.DefaultPort,
		},
	}, "set json webserver config")
}

func GetConfigs() Config {
	return serviceConfig
}
