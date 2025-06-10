package logger

import (
	"log/slog"
	"os"
)

/*
Example of logging in JSON format:

{
	"level": "info",
	"msg": "[START] CRON Job: Checking health claim which real document is not yet upload and created_date is already 14 days or more...",
	"time": "2021-08-25T09:00:00+07:00",
	"id": "f7b3b3b4-0b3b-4b3b-8b3b-0b3b3b3b3b3b",
	"datas": {},
}

*/

var myLogger *slog.Logger

func Init() {
	myLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{

		// ReplaceAttr is a function that can be used to modify the attributes of a log entry before it is serialized.
		// Modify the level key value from INFO to info, DEBUG to debug, WARN to warn, ERROR to error (lowercase)
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			if a.Key == slog.LevelKey {
				a.Key = "level"

				switch a.Value.Any().(slog.Level) {
				case slog.LevelDebug:
					a.Value = slog.StringValue("debug")
				case slog.LevelInfo:
					a.Value = slog.StringValue("info")
				case slog.LevelWarn:
					a.Value = slog.StringValue("warn")
				case slog.LevelError:
					a.Value = slog.StringValue("error")
				}
			}

			return a
		},
	}))
}

// Display log entry (INFO) in JSON format
func LogInfo(id, msg string, datas ...interface{}) {
	if len(datas) > 0 {
		myLogger.Info(
			msg,
			slog.String("id", id),
			slog.Any("datas", datas[0]),
		)
	} else {
		myLogger.Info(
			msg,
			slog.String("id", id),
		)
	}
}

// Display log entry (DEBUG) in JSON format
func LogDebug(id, msg string, datas ...interface{}) {
	if len(datas) > 0 {
		myLogger.Debug(
			msg,
			slog.String("id", id),
			slog.Any("datas", datas[0]),
		)
	} else {
		myLogger.Debug(
			msg,
			slog.String("id", id),
		)
	}
}

// Display log entry (WARN) in JSON format
func LogWarn(id, msg string, datas ...interface{}) {
	if len(datas) > 0 {
		myLogger.Warn(
			msg,
			slog.String("id", id),
			slog.Any("datas", datas[0]),
		)
	} else {
		myLogger.Warn(
			msg,
			slog.String("id", id),
		)
	}
}

// Display log entry (ERROR) in JSON format
func LogError(id, msg string, datas ...interface{}) {
	if len(datas) > 0 {
		myLogger.Error(
			msg,
			slog.String("id", id),
			slog.Any("datas", datas[0]),
		)
	} else {
		myLogger.Error(
			msg,
			slog.String("id", id),
		)
	}
}

// Display log entry (FATAL) in JSON format, and exit the program
func LogFatal(id, msg string, datas ...interface{}) {
	if len(datas) > 0 {
		myLogger.Error(
			msg,
			slog.String("id", id),
			slog.Any("datas", datas[0]),
		)
	} else {
		myLogger.Error(
			msg,
			slog.String("id", id),
		)
	}

	os.Exit(1) // Exit the program with a non-zero status code
}
