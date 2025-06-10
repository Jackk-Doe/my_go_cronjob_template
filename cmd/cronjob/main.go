package main

import (
	"os"
	"time"

	"github.com/Jackk-Doe/my_go_cronjob_template/configs/dotenv"
	"github.com/Jackk-Doe/my_go_cronjob_template/pkg/database"
	cronJobs "github.com/Jackk-Doe/my_go_cronjob_template/pkg/jobs"
	"github.com/Jackk-Doe/my_go_cronjob_template/pkg/logger"

	"github.com/go-co-op/gocron"
	_ "go.uber.org/automaxprocs"
)

func runCronJobs() {

	// CRON job instance, set timezone to Asia/Bangkok
	bnkLocation, _ := time.LoadLocation("Asia/Bangkok")
	jobs := gocron.NewScheduler(bnkLocation)

	/// Example of a cronjob
	/// Run : every day at 00:01
	jobs.Every(1).Day().At("00:01").Do(cronJobs.TestJob)

	/// Example of a cronjob with arguments
	/// Run : every day at 12:00
	jobs.Every(1).Day().At("12:00").Do(cronJobs.TestJob, "Hello", "World")

	jobs.StartBlocking()
}

func main() {

	// Initialize all configurations from here...
	logger.Init()
	database.Init()
	dotenv.Init()

	logger.LogInfo("-", "FINISHED_INITIALIZING_CONFIGURATIONS")

	logger.LogInfo("-", "STARTING_CRON_JOBS_APP...")

	// NOTE: MODE has 'dev', 'uat', 'test', 'build', 'prod' or '' (empty) values
	mode := os.Getenv("MODE")
	if mode == "" {
		// If MODE is not set, panic
		logger.LogFatal("-", "MODE environment variable is not set")
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	appDatas := map[string]interface{}{
		"APP_NAME":     os.Getenv("APP_NAME"),
		"APP_VERSION":  os.Getenv("APP_VERSION"),
		"APP_BUILD_AT": os.Getenv("BUILD_DATE"),
		"APP_RUN_AT":   currentTime,
		"MODE":         mode,
	}

	logger.LogInfo("-", "APP_INFO", appDatas)

	runCronJobs()

	logger.LogInfo("-", "STOPPING_CRON_JOBS_APP...")

	database.Close() // Close database connection
}
