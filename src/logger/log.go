package logger

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
)

var Log *logrus.Logger
var appTag string

//Init 日志初始化
func Init(level string, appName string) {
	Log = NewLogger()

	appTag = appName
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
	//Log.SetFormatter(&logrus.TextFormatter{
	//	ForceColors:               true,
	//	EnvironmentOverrideColors: true,
	//	// FullTimestamp:true,
	//	TimestampFormat: "2006-01-02 15:04:05", //时间格式化
	//	// DisableLevelTruncation:true,
	//})
}

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}
	return logrus.New()

}

func getFileInfo() string {
	_, fileName, codeLine, _ := runtime.Caller(2)
	return fileName + ":" + strconv.Itoa(codeLine)
}

func Debug(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}

	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Debug(msg...)
}

func Debugf(fields map[string]interface{}, format string, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Debugf(format, msg...)
}

func Info(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Info(msg...)
}

func Infof(fields map[string]interface{}, format string, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Infof(format, msg)
}

func Warn(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Warn(msg...)
}

func Warnf(fields map[string]interface{}, format string, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Warnf(format, msg...)
}

func Error(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Error(msg...)
}

func Errorf(fields map[string]interface{}, format string, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Errorf(format, msg...)
}

func Errorln(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Errorln(msg...)
}

func Fatal(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Fatal(msg...)
}

func Panic(fields map[string]interface{}, msg ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = getFileInfo()
	fields["appName"] = appTag
	logrus.WithFields(fields).Panic(msg...)
}
