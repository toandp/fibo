package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/log"
)

var ORM *gorm.DB

// Establishes a DB connection.
func Connect() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.App.Database.Username,
		config.App.Database.Password,
		config.App.Database.Host,
		config.App.Database.Port,
		config.App.Database.DBName,
	)

	ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.App.Database.TablePrefix,
			SingularTable: false,
		},
	})

	if err != nil {
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
		panic(err)
	}

	log.Info("Migration Completed.")
}
