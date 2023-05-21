package mysqlext

import (
	"encoding/json"
	"flag"
	"fmt"
)

var mysqlConfig Config

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"dbName"`
}

type MySQLConfig interface {
	String() string
}

func (c *Config) MarshalText() ([]byte, error) {
	return json.Marshal(*c)
}

func (c *Config) UnmarshalText(b []byte) error {
	return json.Unmarshal(b, c)
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.DbName)
}

func LoadJSONConfigFromFlags(flagSet *flag.FlagSet) {
	flagSet.TextVar(&mysqlConfig, "mysql_config", &Config{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     3306,
		DbName:   "go_signup",
	}, "set json config for MySQL")
}

func GetDefaultMySQLConfig() MySQLConfig {
	return &mysqlConfig
}
