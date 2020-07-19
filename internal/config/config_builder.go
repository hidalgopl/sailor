package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func CreateConfigTemplate() error {
	emptyConfig := &Config{
		Username:  "your SecureAPI username",
		AccessKey: "your SecureAPI access key",
		URL:       "https://api.you.want.to.test.com",
	}
	bytesConfig, err := yaml.Marshal(&emptyConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(".secureapi.yml", bytesConfig, 0644)
	if err != nil {
		return err
	}
	return nil
}
