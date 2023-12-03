package mylog

import (
	"fmt"
	"os"
	"time"
)

// 往终端里写日志的相关代码

// ConsoleLogger 是一个文件日志结构体
type ConsoleLogger struct {
	level Level
}

// NewConsoleLogger 是ConsoleLogger的构造函数
func NewConsoleLogger(levelStr string) (c *ConsoleLogger) {
	level := parseLevel(levelStr)
	cl := &ConsoleLogger{
		level: level,
	}
	return cl
}

// SetLevel 是设置日志级别的方法
func (c *ConsoleLogger) SetLevel(leverStr string) {
	level := parseLevel(leverStr)
	c.level = level
}

// Debug 记录debug日志的方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info 记录info日志的方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warn 记录Warning日志的方法
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error 记录error日志的方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	c.log(ErrorLevel, format, args...)
}

// Fatal 记录fatal日志的方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	c.log(FatalLevel, format, args...)
}

// 真正记录日志的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	if c.level > level {
		return
	}
	// 拼接日志信息 [time][level][fileName:lineNo][funcName] msg
	now := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := getLevelStr(level)
	fileName, lineNo, funcName := getCallerInfo(3)
	msg := fmt.Sprintf(format, args...)
	logMsg := fmt.Sprintf("[%s][%s][%s:%d][%s] %s", now, levelStr, fileName, lineNo, funcName, msg)
	// 对f.file做是否拆分的检测
	fmt.Fprintln(os.Stdout, logMsg)
}

// Close 是关闭日志文件句柄的方法
func (c *ConsoleLogger) Close() {

}
