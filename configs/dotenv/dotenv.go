package dotenv

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/Jackk-Doe/my_go_cronjob_template/pkg/logger"
)

func Init() {
	// If in development mode, load local .env file
	mode := os.Getenv("MODE")

	if mode == "dev" {
		logger.LogInfo("-", "LOADING_ENV_FILE")
		if err := godotenv.Load(); err != nil {
			logger.LogFatal("-", "load .env file error in DEV mode", err.Error())
		}
	}
}
