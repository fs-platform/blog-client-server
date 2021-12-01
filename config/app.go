package config

import "client_server/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"env": config.Env("APP_ENV", "production"),
	})
}
