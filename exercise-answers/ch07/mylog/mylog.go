package mylog

import (
	"path"
	"runtime"
	"strings"
)

// mylogger 是一个封装好的日志库

// Level 是自定义的用来表示日志级别的类型
type Level uint16

// Logger 是日志实例接口
type Logger interface {
	SetLevel(string)
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Close()
}

// 定义常用的5个日志级别
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// 根据日志级别返回对应的日志级别字符串
func getLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}
	return "Unknow"
}

// 解析字符串格式日志级别
func parseLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	}
	return DebugLevel
}

// getCallerInfo 获取当前调用信息
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(fileName)
	return
}
