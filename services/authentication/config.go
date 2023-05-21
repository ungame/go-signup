package authentication

import (
	"flag"
	"github.com/ungame/go-signup/grpcext"
	"github.com/ungame/go-signup/httpext"
)

const DefaultPort = 8081

var serviceConfig grpcext.Config

func LoadJSONConfigFromFlags(flagSet *flag.FlagSet) {
	flagSet.TextVar(&serviceConfig, "authentication_service", &grpcext.Config{
		Host: httpext.Localhost,
		Port: DefaultPort,
	}, "set json config for authentication service")
}

func GetConfigs() grpcext.Config {
	return serviceConfig
}
