package logger

import (
	"GoWild/common/loader"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var (
	apiLogger         *logrus.Logger
	errlogger         *logrus.Logger
	infoLogger        *logrus.Logger
	sqlLogger         *logrus.Logger
	searchLogger      *logrus.Logger
	acrLogger         *logrus.Logger
	greenLogger       *logrus.Logger
	oppoMigrateLogger *logrus.Logger
)

func init() {
	logPath, err := loader.LoadAppConfig().GetString("log", "path")
	if err != nil {
		panic(err)
	}

	apiLogger = createLogger(logPath, "api")
	errlogger = createLogger(logPath, "error")
	errlogger.SetReportCaller(true)
	infoLogger = createLogger(logPath, "info")
	sqlLogger = createLogger(logPath, "sql")
	searchLogger = createLogger(logPath, "search")
	acrLogger = createLogger(logPath, "acrcloud")
	greenLogger = createLogger(logPath, "aliGreen")
	oppoMigrateLogger = createLogger(logPath, "oppoMigrate")

	//feishuWebhook, _ := loader.LoadAppConfig().GetString("feishu", "webhook")
	//if len(feishuWebhook) > 0 && !loader.IsDebugMode() {
	//	feishuHook, err := feishu.NewFeiShuHook(feishuWebhook, nil)
	//	if err != nil {
	//		panic(err)
	//		return
	//	}
	//	errlogger.AddHook(feishuHook)
	//}
	//
	//dingdingWebhook, _ := loader.LoadAppConfig().GetString("dingding", "webhook")
	//dingdingSignKey, _ := loader.LoadAppConfig().GetString("dingding", "sign")
	//if len(dingdingWebhook) > 0 && len(dingdingSignKey) > 0 && !loader.IsDebugMode() {
	//	dingdingHook, err := dingding.NewDingDingHook(dingdingWebhook, dingdingSignKey, nil)
	//	if err != nil {
	//		panic(err)
	//		return
	//	}
	//	errlogger.AddHook(dingdingHook)
	//}
}

func ApiLogger() *logrus.Logger {
	return apiLogger
}

func ErrLogger() *logrus.Logger {
	return errlogger
}

func InfoLogger() *logrus.Logger {
	return infoLogger
}

func SQLLogger() *logrus.Logger {
	return sqlLogger
}

func SearchLogger() *logrus.Logger {
	return searchLogger
}

func AcrLogger() *logrus.Logger {
	return acrLogger
}

func GreenLogger() *logrus.Logger {
	return greenLogger
}

func OppoMigrateLogger() *logrus.Logger {
	return oppoMigrateLogger
}

func createLogger(logPath, fileName string) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	writeToFile := createLogOutputFile(logPath, fileName)
	if loader.IsDebugMode() {
		logger.SetOutput(io.MultiWriter(writeToFile, os.Stdout))
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetOutput(writeToFile)
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func createLogOutputFile(logFilePath, logFileName string) io.Writer {
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		_ = os.MkdirAll(path.Dir(fileName), 0777)
	}
	w, err := rotatelogs.New(
		fileName+"-%Y%m%d.log",
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(7))
	if err != nil {
		panic(err)
	}
	return w
}
