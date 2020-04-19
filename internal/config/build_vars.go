package config

import (
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type buildConfig struct {
	NatsUrl  string `yaml:"natsUrl"`
	FrontUrl string `yaml:"frontUrl"`
	APIUrl   string `yaml:"apiUrl"`
}


func LoadBuildConfig() (*buildConfig, error) {
	box := packr.NewBox("./../../secrets")
	contents, err := box.Find("staging.yaml")
	if err != nil {
		logrus.Fatal(err)
		return &buildConfig{}, err
	}
	var buildCfg buildConfig
	err = yaml.Unmarshal([]byte(contents), &buildCfg)
	return &buildCfg, nil

}
