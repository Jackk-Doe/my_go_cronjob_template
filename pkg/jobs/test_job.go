package jobs

import (
	"time"

	"github.com/Jackk-Doe/my_go_cronjob_template/pkg/logger"

	"github.com/google/uuid"
)

func TestJob() {
	jobID := uuid.New().String()

	logger.LogInfo(jobID, "START_TEST_JOB...")

	logger.LogInfo(jobID, "Connecting to database...")

	logger.LogInfo(jobID, "END_TEST_JOB...")
}

func TestJobWithArgs(arg1, arg2 string) {
	jobID := uuid.New().String()

	logger.LogInfo(jobID, "START_TEST_JOB_WITH_ARGS...", map[string]interface{}{
		"arg1": arg1,
		"arg2": arg2,
	})

	// Simulate some processing
	time.Sleep(2 * time.Second)

	logger.LogInfo(jobID, "END_TEST_JOB_WITH_ARGS...")

}
