package database

import (
	"github.com/spf13/viper"
	"net/url"
)

func init() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.username", "limo")
	viper.SetDefault("db.password", "limo")
	viper.SetDefault("db.database", "limo")
}

type Config struct {
	Enabled      bool
	Hostname     string
	Port         string
	Username     string
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
}

// BuildDsn builds out a Postgres DSN from the current config.
func (client *Config) BuildDsn() (dsn string) {
	dsn += "host=" + url.QueryEscape(client.Hostname)
	if client.Username != "" {
		dsn += " user=" + url.QueryEscape(client.Username)
	}
	if client.Password != "" {
		dsn += " password=" + url.QueryEscape(client.Password)
	}
	if client.Database != "" {
		dsn += " dbname=" + url.QueryEscape(client.Database)
	}
	dsn += " port=" + url.QueryEscape(client.Port)
	return dsn
}

// NewConfig creates a new config object with defaults from Viper.
func NewConfig() Config {
	return Config{
		Hostname:     viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		Database:     viper.GetString("db.database"),
		MaxIdleConns: viper.GetInt("db.max-idle"),
		MaxOpenConns: viper.GetInt("db.max-open"),
	}
}
