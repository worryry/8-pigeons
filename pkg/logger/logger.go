package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/worryry/8-pigeons/pkg/setting"
	"log"
	"os"
	"time"
)

const (
	DefaultPath = "./logs"

	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
	FatalLevel = "fatal"
	TraceLevel = "trace"

	filePrefix    = "app-"
	JsonFormatter = "json"
	TextFormatter = "text"

	TimestampFormat = "2006-01-02_15-04-05"
)
const logMsgLen = 10000

var logFileHandle *os.File

var logLevelMap map[string]logrus.Level

func Start() {
	target := setting.GetString("log.target")
	if target == "console" {
		loggerInit(JsonFormatter, setting.GetString("log.level"), "")
	} else {
		loggerInit(JsonFormatter, setting.GetString("log.level"), GetLogFileName())
	}
}

func loggerInit(format string, level string, path string) {
	_, err := os.Stat(DefaultPath)
	if err != nil {
		err := os.MkdirAll(DefaultPath, os.ModePerm)
		if err != nil {
			log.Println("创建日志目录失败：", err)
			return
		}
	}
	initLoggerFormatter(format)
	initLoggerLevel(level)
	if len(path) == 0 {
		initConsole()
	} else {
		initLoggerPath(path)
	}
	initLogLevelMap()
}
func initLoggerFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(&FastJsonFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
func initLoggerPath(path string) {
	if logFileHandle != nil {
		if err := logFileHandle.Close(); err != nil {
			log.Println("文件句柄关闭失败")
		}
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("日志文件打开失败：", err)
	}
	logrus.SetOutput(file)
	logFileHandle = file
}

func initLoggerLevel(level string) {
	switch level {
	case DebugLevel:
		logrus.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		logrus.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		logrus.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		logrus.SetLevel(logrus.ErrorLevel)
	case PanicLevel:
		logrus.SetLevel(logrus.PanicLevel)
	case FatalLevel:
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func initConsole() {
	logrus.SetOutput(os.Stdout)
}

func initLogLevelMap() {
	logLevelMap = make(map[string]logrus.Level, 7)
	logLevelMap[DebugLevel] = logrus.DebugLevel
	logLevelMap[InfoLevel] = logrus.InfoLevel
	logLevelMap[WarnLevel] = logrus.WarnLevel
	logLevelMap[ErrorLevel] = logrus.ErrorLevel
	logLevelMap[FatalLevel] = logrus.FatalLevel
	logLevelMap[TraceLevel] = logrus.TraceLevel
	logLevelMap[PanicLevel] = logrus.PanicLevel
}

func printLog(level string, isFormat bool, format string, msg ...interface{}) {
	var lv logrus.Level
	if v, ok := logLevelMap[level]; ok {
		lv = v
	} else {
		lv = logrus.DebugLevel
	}
	logDefault(lv, isFormat, format, msg...)
}
func Debug(msg ...interface{}) {
	printLog(DebugLevel, false, "", msg...)
}

func Info(msg ...interface{}) {
	printLog(InfoLevel, false, "", msg...)
}

func Warn(msg ...interface{}) {
	printLog(WarnLevel, false, "", msg...)
}

func Error(msg ...interface{}) {
	printLog(ErrorLevel, false, "", msg...)
}

func Fatal(msg ...interface{}) {
	printLog(FatalLevel, false, "", msg...)
}

func Panic(msg ...interface{}) {
	printLog(PanicLevel, false, "", msg...)
}

func Debugf(format string, args ...interface{}) {
	printLog(DebugLevel, true, format, args...)
}

func Infof(format string, args ...interface{}) {
	printLog(InfoLevel, true, format, args...)
}

func Warnf(format string, args ...interface{}) {
	printLog(WarnLevel, true, format, args...)
}

func Errorf(format string, args ...interface{}) {
	printLog(ErrorLevel, true, format, args...)
}

func Panicf(format string, args ...interface{}) {
	printLog(PanicLevel, true, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	printLog(FatalLevel, true, format, args...)
}

func logDefault(level logrus.Level, isFormat bool, format string, msg ...interface{}) {
	lg := logrus.StandardLogger()
	if isFormat {
		lg.Logf(level, format, msg...)
	} else {
		lg.Log(level, msg...)
	}
}

func GetLogFileName() string {
	logDir := os.Getenv("BAYMAX_APP_LOG_DIR")
	if len(logDir) == 0 {
		logDir = DefaultPath
	}

	return logDir + "/" + filePrefix + time.Now().Format(TimestampFormat) + ".log"
}
