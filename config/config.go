package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// AppConfig is a struct holding the application settings.
type AppConfig struct {
	Debug    bool `env:"APP_DEBUG" env-default:"false"`
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig is a struct holding the server settings.
type ServerConfig struct {
	Name        string `env:"APP_NAME" env-default:"fibo"`
	Host        string `env:"APP_HOST" env-default:"localhost"`
	Port        string `env:"APP_PORT" env-default:"8080"`
	ProxyHeader string `mapstructure:"PROXY_HEADER" env:"PROXY_HEADER" env-default:"*"`
	UploadSize  int    `mapstructure:"UPLOAD_SIZE" env:"UPLOAD_SIZE" env-default:"400"`
}

// DatabaseConfig is a struct holding the database settings.
type DatabaseConfig struct {
	Driver      string `env:"DB_DRIVER" env-default:"mysql"`
	Host        string `env:"DB_HOST" env-default:"127.0.0.1"`
	Username    string `env:"DB_USER" env-default:"root"`
	Password    string `env:"DB_PASS" env-default:""`
	DBName      string `env:"DB_NAME" env-default:"fibo_dev"`
	Port        int    `env:"DB_PORT" env-default:"3306"`
	TablePrefix string `env:"DB_TABLE_PREFIX" env-default:"tbl_"`
}

type JWTConfig struct {
	Expire int64  `env:"JWT_EXPIRE" env-default:"3600"`
	Secret string `env:"JWT_SECRET" env-default:"1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426"`
}

var App AppConfig

func LoadConfigFromEnv() {
	var err error

	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	err = cleanenv.ReadEnv(&App)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
