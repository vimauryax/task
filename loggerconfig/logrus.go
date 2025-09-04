package loggerconfig

import (
	"fmt"
	"log"
	"os"
	//"regexp"
	"runtime"
	"mv/mvto-do/constants"
	"strings"
	//"sync"
	//"time"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

var Info = func(args ...interface{}) {
	InfoImpl(args...)
}

func InfoImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Info(message)
}

var Warn = func(args ...interface{}) {
	WarnImpl(args...)
}

func WarnImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Warn(message)
}

var Panic = func(args ...interface{}) {
	PanicImpl(args...)
}

func PanicImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Panic(message)
}

func buildMessage(args ...interface{}) string {
	var message string
	for i, arg := range args {
		if i == 0 {
			message = fmt.Sprint(arg)
		} else {
			message += " " + fmt.Sprint(arg)
		}
	}
	return message
}

var logger *logrus.Logger

func LogrusInitialize() {

	env := os.Getenv("GO_ENV")

	logger = logrus.New()
	if env == "local" || env == "debug" {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("InitLogrus Debug logging enabled.")
	}

	// Set logrus configuration
	logger.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "(service-SPACE)", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)

	elasticApmServiceName := os.Getenv(constants.ElasticApmServiceName)
	elasticAPM, err := apm.NewTracer(elasticApmServiceName, constants.ServiceVersion)
	if err != nil {
		log.Printf("Getting Error-%v\n", err)
	}

	var logLevels []logrus.Level

	for _, level := range logrus.AllLevels {
		logLevels = append(logLevels, level)
	}

	logger.AddHook(&apmlogrus.Hook{
		Tracer:    elasticAPM,
		LogLevels: logLevels,
	})
	
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
