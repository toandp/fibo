package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db/entity"
)

var ORM *gorm.DB

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
		fmt.Println(err)
		os.Exit(2)
	}

	ORM.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(100),
	)
}

func Migrate() {
	log.Println("Initiating migration...")

	err := ORM.Migrator().AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	log.Println("Migration Completed...")
}
