package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log logger
)

type bookstoreLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}


// var (
// 	log *zap.Logger
// )

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"}, 
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), 
		Encoding:  "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level", 
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder, 			
		},

	}
	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// func GetLogger() *zap.Logger {
// 	return log
// }

func GetLogger() bookstoreLogger {
	return log
}


func Info(msg string, tags...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string,  err error, tags...zap.Field) {
	
	tags = append(tags, zap.NamedError("error", err))	
	log.log.Error(msg, tags...)
	log.log.Sync()
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}
