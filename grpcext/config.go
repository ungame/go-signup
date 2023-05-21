package grpcext

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *Config) ClientAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) MarshalText() ([]byte, error) {
	return json.Marshal(*c)
}

func (c *Config) UnmarshalText(b []byte) error {
	return json.Unmarshal(b, c)
}
