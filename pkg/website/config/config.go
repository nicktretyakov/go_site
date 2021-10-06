package config

import (
	"/pkg/fo_errors"
	"/pkg/tools/helpers"
)

type Config struct {
	HTTPAddr           string
	DBConnectionString string
	LogPath            string
	LogLevel           string
	Debug              bool
}

// NewWebsiteConfig instantiate config from environment
func NewWebsiteConfig() (*Config, error) {
	cfg := Config{}

	helpers.EnvToString(&cfg.HTTPAddr, "FO_WEBSITE_HTTP_ADDR", "127.0.0.1:5000")
	helpers.EnvToString(&cfg.DBConnectionString, "FO_WEBSITE_DB_CONNECTION", "")
	helpers.EnvToString(&cfg.LogPath, "FO_WEBSITE_LOGPATH", "./debug.log")
	helpers.EnvToString(&cfg.LogLevel, "FO_WEBSITE_LOGLEVEL", "debug")
	helpers.EnvToBool(&cfg.Debug, "FO_WEBSITE_DEBUG", false)

	if cfg.DBConnectionString == "" {
		return nil, fo_errors.ErrDBConnection
	}

	if cfg.HTTPAddr == "" {
		return nil, fo_errors.ErrHttpAddrWebsite
	}

	return &cfg, nil
}
