package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	SECUREAPI_FILENAME = ".secureapi"
	SECUREAPI_FILE = SECUREAPI_FILENAME + ".yml"
)

// Config ...
type Config struct {
	Username  string   `yaml:"username"`
	AccessKey string   `yaml:"accessKey"`
	URL       string   `yaml:"url"`
}

// PrettyPrint ...
func (c *Config) PrettyPrint() string {
	configStr := fmt.Sprintf(
		"username: %s \naccess_key: <hidden> \nurl: %s", c.Username, c.URL)
	return configStr
}

// GetConf ...
func GetConf() *Config {
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		logrus.Errorf("unable to decode into config struct, %v", err)
	}
	return conf
}
