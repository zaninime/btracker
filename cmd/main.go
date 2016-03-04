package main

import (
	"fmt"
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
	"github.com/mgutz/logxi/v1"
	"github.com/zaninime/btracker/db"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath  = kingpin.Flag("config", "Configuration file path. Default is btracker.toml.").Short('c').PlaceHolder("PATH").Default("btracker.toml").String()
	exampleMode = kingpin.Flag("template", "Prints a configuration file template with comments.").Short('e').Bool()
)

const version = "0.9.0"

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

	dbModes := map[string]db.Mode{
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

	// cron scheduler
	s := gocron.NewScheduler()
	s.Every(2).Minutes().Do(db.ClearConnections)
	s.Every(30).Minutes().Do(db.ClearPeers, 1200)
	go cron(s)

	// setup socket
	localAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", config.Tracker.BindAddress, config.Tracker.Port))
	if err != nil {
		mainLogger.Fatal("couldn't parse bind address", "err", err)
	}
	conn, err := net.ListenUDP("udp4", localAddr)
	if err != nil {
		mainLogger.Fatal("couldn't bind to address", "addr", localAddr, "err", err)
	}
	mainLogger.Info("listening", "addr", localAddr)
	listenUDP(conn)
}

func cron(sched *gocron.Scheduler) {
	<-sched.Start()
}
