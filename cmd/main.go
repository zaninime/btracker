package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/mgutz/logxi/v1"
	"go.zanini.me/btracker/db"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath  = kingpin.Flag("config", "Configuration file path. Default is btracker.toml.").Short('c').PlaceHolder("PATH").Default("btracker.toml").String()
	exampleMode = kingpin.Flag("template", "Prints a configuration file template with comments.").Short('e').Bool()
)

const version = "0.0.1"

var mainLogger log.Logger

func main() {
	kingpin.Parse()

	// example mode handling
	if *exampleMode {
		fmt.Println(configTemplate)
		return
	}

	// configuration first
	file, err := os.Open(*configPath)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = toml.DecodeReader(file, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse config file: %s\n", err)
		os.Exit(1)
	}

	if err = validateConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid configuration: %s\n", err)
		os.Exit(1)
	}

	// initialize logging
	logLevels := map[string]int{
		"alert":    log.LevelAlert,
		"critical": log.LevelCritical,
		"error":    log.LevelError,
		"warning":  log.LevelWarn,
		"notice":   log.LevelNotice,
		"info":     log.LevelInfo,
		"debug":    log.LevelDebug,
		"trace":    log.LevelTrace,
	}

	mainLogger = log.New("main")
	mainLogger.SetLevel(logLevels[config.Logging.Level])
	mainLogger.Info("btracker is starting", "version", version)

	// initialize database
	dbLogger := log.New("db")
	dbLogger.SetLevel(logLevels[config.Logging.Level])
	db.InitializeModule(dbLogger)

	dbModes := map[string]int{
		"prod": db.ModeProd,
		"dev":  db.ModeDev,
	}
	dbConfig := db.Config{
		config.DB.ConnectionString,
		config.DB.IdleConn,
		config.DB.MaxConn,
		dbModes[config.DB.Mode],
	}
	if err = db.InitializeDatabaseConnection(dbConfig); err != nil {
		mainLogger.Fatal("db initialization failed", "err", err)
	}
}
