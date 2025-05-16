package app

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/nyakovchuk/vsch_youtube_bot/config"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/db"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/logger"

	_ "github.com/mattn/go-sqlite3"
)

var allowedEnvs = []string{"dev", "prod"}

type App struct {
	config *config.Config
	logger *slog.Logger
	db     *sql.DB
}

// Singleton
var (
	instance *App
	once     sync.Once
)

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Logger() *slog.Logger {
	return a.logger
}

func (a *App) DB() *sql.DB {
	return a.db
}

func NewApp(config *config.Config, logger *slog.Logger, db *sql.DB) *App {
	return &App{
		config: config,
		logger: logger,
		db:     db,
	}
}

// GetApp returns singleton
func GetApp() *App {
	once.Do(func() {
		fmt.Print("Loading configuration...")
		configFile := getConfigFile()
		config, err := config.LoadConfig(configFile)
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}
		fmt.Println("DONE")

		fmt.Print("Setting up logging...")
		logger, err := logger.SetupLogger(config, nil)
		if err != nil {
			fmt.Println("Error setting up logging:", err)
			return
		}
		fmt.Println("DONE")

		fmt.Printf("\nDatabase driver: %s\n", strings.ToUpper(config.DbDriver))

		fmt.Print("Connecting to database...")
		db, err := db.ConnectDB(config)
		if err != nil {
			fmt.Println("Error connecting to database:", err)
		}
		fmt.Println("DONE")

		instance = NewApp(config, logger, db)
	})

	if instance == nil {
		log.Fatal("Failed to set up the configuration")
	}

	return instance
}

func getConfigFile() string {
	env := os.Getenv("EXT_ENV_FILE")

	if slices.Contains(allowedEnvs, env) {
		return ".env." + env
	}
	return ".env"
}
