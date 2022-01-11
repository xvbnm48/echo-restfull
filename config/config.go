package config

import "github.com/tkanos/gonfig"

type configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func GetConfig() configuration {
	conf := configuration{}

	gonfig.GetConf("config/config.json", &conf)
	return conf
}
