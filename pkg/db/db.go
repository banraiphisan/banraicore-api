package db

import (
	"fmt"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/pkg/common/exception"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DB struct {
	*gorm.DB
}

func buildDSN(config config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Database,
		config.Port,
	)
}

func NewDB(conf config.DatabaseConfig) (*DB, error) {

	maxPoolOpen := conf.MaxPoolOpen
	maxPoolIdle := conf.MaxPoolIdle
	maxPollLifeTime := conf.MaxPollLifeTime

	// 🔧 Improved Log Level Handling
	var logLevel logger.LogLevel
	switch conf.LogLevel {
	case "INFO":
		logLevel = logger.Info
	case "WARN":
		logLevel = logger.Warn
	case "ERROR":
		logLevel = logger.Error
	default:
		logLevel = logger.Silent
	}

	loggerDB := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := buildDSN(conf)
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{
		Logger: loggerDB,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxPollLifeTime) * time.Millisecond)

	// Auto-migrate models (Uncomment when models are ready)
	// db.AutoMigrate(&entity.User{})

	return &DB{db}, nil
}
