package config

import (
	"errors"
	"os"
)

type Config struct {
	// DB config
	dBHost string
	dBPort string
	dBUser string
	dBPass string
	dBName string

	// Server config
	port string
}

func (c *Config) DBHost() string {
	return c.dBHost
}

func (c *Config) DBPort() string {
	return c.dBPort
}

func (c *Config) DBUser() string {
	return c.dBUser
}

func (c *Config) DBPass() string {
	return c.dBPass
}

func (c *Config) DBName() string {
	return c.dBName
}

func (c *Config) Port() string {
	return c.port
}

var Cfg *Config

func LoadConfig() error {
	Cfg = &Config{}
	var ok bool
	Cfg.dBHost, ok = os.LookupEnv("DB_HOST")
	if !ok {
		return errors.New("DB_HOST environment variable not set")
	}

	Cfg.dBPort, ok = os.LookupEnv("DB_PORT")
	if !ok {
		return errors.New("DB_PORT environment variable not set")
	}

	Cfg.dBUser, ok = os.LookupEnv("DB_USER")
	if !ok {
		return errors.New("DB_USER environment variable not set")
	}

	Cfg.dBPass, ok = os.LookupEnv("DB_PASS")
	if !ok {
		return errors.New("DB_PASS environment variable not set")
	}

	Cfg.dBName, ok = os.LookupEnv("DB_NAME")
	if !ok {
		return errors.New("DB_NAME environment variable not set")
	}

	Cfg.port, ok = os.LookupEnv("PORT")
	if !ok {
		return errors.New("PORT environment variable not set")
	}

	return nil
}
