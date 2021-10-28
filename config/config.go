package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// AppConfig is a struct holding the application settings.
type AppConfig struct {
	Env      string `env:"APP_ENV" env-default:"development"`
	Debug    bool   `env:"APP_DEBUG" env-default:"false"`
	Timezone string `env:"APP_TZ" env-default:"Asia/Ho_Chi_Minh"`
	Server   ServerConfig
	DB       DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
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
	SSLMode     string `env:"DB_SSL_MODE" env-default:"disable"`
	SQLiteFile  string `env:"DB_SQLITE_FILE" env-default:"sqlite.db"`
}

// JWTConfig is a struct holding the JWT settings.
type JWTConfig struct {
	Expire int64  `env:"JWT_EXPIRE" env-default:"3600"`
	Secret string `env:"JWT_SECRET" env-default:"1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426"`
}

// LogConfig is a struct holding the JWT settings.
type LogConfig struct {
	FilePath string `env:"LOG_FILE_FORMAT" env-default:"./tmp/logs/%s-%s.log"`
}

var App AppConfig

// Read configuration from the environment variables.
func LoadConfigFromEnv() {
	var err error

	// Load env variables
	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Bind configuration
	err = cleanenv.ReadEnv(&App)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
