package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("Invalid DB_DSN variable.")
	}

	appEnv := os.Getenv("APP_ENV")

	var logLevel logger.LogLevel

	switch appEnv {
	case "development":
		logLevel = logger.Info
	case "staging":
		logLevel = logger.Warn
	case "production":
		logLevel = logger.Error
	default:
		logLevel = logger.Silent // fallback
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	sqlDB, err := db.DB()
	

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}
