package log

import (
	"go.uber.org/zap"
)

var Zap *zap.Logger

// Setup Zap logger.
// @return void
func SetupLogger() {
	var (
		err error
		// zapCfg zap.Config
	)

	// switch config.App.Env {
	// case "production":
	// 	zapCfg = zap.NewProductionConfig()
	// default:
	// 	zapCfg = zap.NewDevelopmentConfig()
	// }

	// zapCfg.OutputPaths = []string{
	// 	"stdout",
	// 	fmt.Sprintf("./tmp/logs/%s-%s.log", config.App.Server.Name, config.App.Env),
	// }

	// Zap, err = zapCfg.Build()

	Zap, err = zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	defer Zap.Sync()
}
