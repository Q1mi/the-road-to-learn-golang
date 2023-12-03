package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志的相关代码

// FileLogger 是一个文件日志结构体
type FileLogger struct {
	level       Level
	filePath    string
	fileName    string
	file        *os.File
	errFile     *os.File
	maxFileSize int64
}

// NewFileLogger 是FileLogger的构造函数
func NewFileLogger(levelStr, filePath, fileName string) (f *FileLogger) {
	level := parseLevel(levelStr)
	fl := &FileLogger{
		level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: 10 * 1024 * 1024,
	}
	fl.initLogFile()
	return fl
}

// 初始化文件句柄的方法（打开文件等待写日志）
func (f *FileLogger) initLogFile() {
	// logFileName := fmt.Sprintf("%s/%s", f.filePath, f.fileName)
	logFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", logFileName))
	}
	f.file = fileObj
	// 打开会创建一个专门记录错误日志的文件
	errFileName := fmt.Sprintf("%s.err", logFileName)
	errFileObj, err := os.OpenFile(errFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", errFileName))
	}
	f.errFile = errFileObj
}

// SetLevel 是设置日志级别的方法
func (f *FileLogger) SetLevel(leverStr string) {
	level := parseLevel(leverStr)
	f.level = level
}

// Debug 记录debug日志的方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info 记录info日志的方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn 记录Warning日志的方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error 记录error日志的方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	f.log(ErrorLevel, format, args...)
}

// Fatal 记录fatal日志的方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	f.log(FatalLevel, format, args...)
}

// 真正记录日志的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	// 如果日志级别不够就啥也不干
	if f.level > level {
		return
	}
	// 拼接日志信息 [time][level][fileName:lineNo][funcName] msg
	now := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := getLevelStr(level)
	fileName, lineNo, funcName := getCallerInfo(3)
	msg := fmt.Sprintf(format, args...)
	logMsg := fmt.Sprintf("[%s][%s][%s:%d][%s] %s", now, levelStr, fileName, lineNo, funcName, msg)

	// 对f.file做是否拆分的检测
	if f.checkSplit(f.file) {
		f.file = f.splitFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg)
	if level >= ErrorLevel {
		// 对f.errFile做是否拆分的检测
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

func (f *FileLogger) checkSplit(file *os.File) bool {
	statInfo, err := file.Stat()
	if err != nil {
		return false
	}
	fileSize := statInfo.Size()
	return fileSize >= f.maxFileSize
}

// 将日志文件按文件大小切分
func (f *FileLogger) splitFile(file *os.File) *os.File {
	fileName := file.Name()
	now := time.Now()
	backupName := fmt.Sprintf("%s_%s", fileName, now.Format("20060102150405"))
	file.Close()                    // 关闭原来的文件
	os.Rename(fileName, backupName) // 备份文件
	newFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("创建日志文件%s失败", fileName))
	}
	return newFile
}

// Close 是关闭日志文件句柄的方法
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
