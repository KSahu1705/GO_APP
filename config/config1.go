package config

import (
	"time"
)

type Config struct {
	DB *DBConfig
}
 
type DatabaseConfig struct {
	Name                   string
	Host                   string
	Port                   int
	// Credentials            string // from vault
	// SSL                    string
	Slave                  string
	Master                 string
	MaxOpenConn            int
	MaxIdleConn            int
	ConnMaxLifetimeSeconds time.Duration
}

type DBConfig struct {
	Dialect  string
    Host     string
    Port     int
    User     string
    Password string
    DBname   string
}
 
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect :  "postgres",
			Host:      "localhost",
			Port:      5433,
			User:      "kriti",
			Password:  "nkx01",
			DBname:    "go_dummy",
		},
	}
}

