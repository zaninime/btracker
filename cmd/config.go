package main

import (
	"fmt"
	"net"
)

var config = ConfigRoot{}

type ConfigRoot struct {
	Tracker TrackerConfig  `toml:"tracker"`
	DB      DatabaseConfig `toml:"database"`
	Logging LoggingConfig  `toml:"logging"`
}

type TrackerConfig struct {
	BindAddress   string `toml:"bind_address"`
	Port          int    `toml:"port"`
	AllowScraping bool   `toml:"allow_scraping"`
	Interval      int    `toml:"interval"`
}

type DatabaseConfig struct {
	ConnectionString string `toml:"connection_string"`
	IdleConn         int    `toml:"idle_connections"`
	MaxConn          int    `toml:"max_connections"`
	Mode             string `toml:"mode"`
}

type LoggingConfig struct {
	Level string `toml:"level"`
}

func validateConfig() error {
	// tracker config
	if bindAddress := net.ParseIP(config.Tracker.BindAddress); bindAddress == nil {
		return fmt.Errorf("tracker: invalid bind address: %s", config.Tracker.BindAddress)
	}

	if config.Tracker.Port < 1 || 65535 < config.Tracker.Port {
		return fmt.Errorf("tracker: UDP bind port out of range")
	}

	if config.Tracker.Interval < 60 || 3600*2 < config.Tracker.Interval {
		return fmt.Errorf("tracker: announce interval out of range (allowed %d-%d)", 60, 3600*2)
	}

	// database config
	if len(config.DB.ConnectionString) == 0 {
		return fmt.Errorf("database: empty connection string")
	}

	dbMode := map[string]bool{
		"prod": true,
		"dev":  true,
	}
	if !dbMode[config.DB.Mode] {
		return fmt.Errorf("database: invalid connection mode")
	}

	if config.DB.IdleConn < 1 {
		return fmt.Errorf("database: idle connections must be at least 1")
	}

	if config.DB.MaxConn < config.DB.IdleConn {
		return fmt.Errorf("database: # of max active connections must be greater or equal than # of idle connections")
	}

	// logging
	logLevel := map[string]bool{
		"alert":    true,
		"critical": true,
		"error":    true,
		"warning":  true,
		"notice":   true,
		"info":     true,
		"debug":    true,
		"trace":    true,
	}
	if !logLevel[config.Logging.Level] {
		return fmt.Errorf("logging: invalid level")
	}
	return nil
}

const configTemplate = `[tracker]
bind_address = "0.0.0.0"  # bind address for udp socket
port = 1234               # udp port
allow_scraping = true     # enables or disables scraping functionality
interval = 900            # interval between announces sent to the clients (sec)

[database]
# Supported parameters:
# * dbname - The name of the database to connect to
# * user - The user to sign in as
# * password - The user's password
# * host - The host to connect to. Values that start with / are for unix domain
#          sockets. (default is localhost)
# * port - The port to bind to. (default is 5432)
# * sslmode - Whether or not to use SSL (default is require, this is not the
#             default for libpq)
# * connect_timeout - Maximum wait for connection, in seconds.
#                     Zero or not specified means wait indefinitely.
# * sslcert - Cert file location. The file must contain PEM encoded data.
# * sslkey - Key file location. The file must contain PEM encoded data.
# * sslrootcert - The location of the root certificate file. The file must
#                 contain PEM encoded data.
connection_string = "postgres://btracker@localhost/btracker?sslmode=disable"
idle_connections = 2          # keep at most # idle connections
max_connections = 5           # max concurrent connections to db
mode = "prod|dev"             # set running mode. dev is slower (query debug)

[logging]
level = "info"      # one of: alert, critical, error, warning, notice, info,
#                             debug, trace`
