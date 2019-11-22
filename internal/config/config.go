package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Username  string   `yaml:username`
	AccessKey string   `yaml:accessKey`
	Url       string   `yaml: url`
	Tests     []string `yaml: tests`
}

func (c *Config) PrettyPrint() string {
	configStr := fmt.Sprintf(
		"username: %s \naccess_key: <hidden> \nurl: %s \ntests: %s", c.Username, c.Url, c.Tests)
	return configStr
}

func GetConf() *Config {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}
