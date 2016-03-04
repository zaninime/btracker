package db

import (
	"github.com/mgutz/logxi/v1"

	"github.com/jmoiron/sqlx"
)

// Mode is the type for representing database connection mode
type Mode int

// Database operation modes
const (
	ModeProd Mode = iota
	ModeDev
)

// Config is the struct containing the database connection parameters
type Config struct {
	DataSource string
	IdleConn   int
	MaxConn    int
	Mode       Mode
}

var (
	// DB is the sqlx database connection
	DB     *sqlx.DB
	logger log.Logger
)

// InitializeModule setup the db module for correct operation
func InitializeModule(dbLogger log.Logger) {
	if dbLogger == nil {
		panic("invalid logger configuration")
	}
	logger = dbLogger
	logger.Info("module ready")
}

// InitializeDatabaseConnection tries a connections to the database, updates
// it's schema and makes it ready for use
func InitializeDatabaseConnection(config Config) error {
	var err error
	DB, err = sqlx.Connect("postgres", config.DataSource)
	if err != nil {
		return err
	}

	DB.SetMaxIdleConns(config.IdleConn)
	DB.SetMaxOpenConns(config.MaxConn)

	// cleanup immediately
	if err = ClearConnections(); err != nil {
		return err
	}
	if err = ClearPeers(3600); err != nil {
		return err
	}

	// start pinging
	startPing()

	// go verify and update schema
	err = checkAndUpdateSchema()
	if err != nil {
		return err
	}

	return prepareStatements()
}
