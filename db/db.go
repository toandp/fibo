package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"moul.io/zapgorm2"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/log"
)

var ORM *gorm.DB

// Establishes a DB connection.
func Connect() {
	var (
		err error
		dsn string
		cfg *gorm.Config
	)

	logger := zapgorm2.New(log.Zap)
	logger.SetAsDefault()

	cfg = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.App.DB.TablePrefix,
			SingularTable: false,
		},
		Logger: logger,
	}

	switch config.App.DB.Driver {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.App.DB.Username,
			config.App.DB.Password,
			config.App.DB.Host,
			config.App.DB.Port,
			config.App.DB.DBName,
		)

		ORM, err = gorm.Open(mysql.Open(dsn), cfg)
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			config.App.DB.Host,
			config.App.DB.Username,
			config.App.DB.Password,
			config.App.DB.DBName,
			config.App.DB.Port,
			config.App.DB.SSLMode,
			config.App.Timezone,
		)

		ORM, err = gorm.Open(postgres.Open(dsn), cfg)
	default:
		ORM, err = gorm.Open(sqlite.Open(config.App.DB.SQLiteFile), cfg)
	}

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	ORM.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(100),
	)
}

// Execute the DB migration.
func Migrate() {
	log.Info("Initiating migration...")

	err := ORM.Migrator().AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Info("Migration Completed.")
}
