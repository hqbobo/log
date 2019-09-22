package log

import (
	"fmt"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
	"io"
	"runtime"
	"strings"
	"time"
)

var WarnLevels = []log.Level{
	log.WarnLevel,
	log.ErrorLevel,
	log.PanicLevel,
	log.AlertLevel,
	log.FatalLevel,
}

var InFoLevels = []log.Level{
	log.InfoLevel,
	log.NoticeLevel,
	log.WarnLevel,
	log.ErrorLevel,
	log.PanicLevel,
	log.AlertLevel,
	log.FatalLevel,
}

var DebugLevels = []log.Level{
	log.DebugLevel,
	log.InfoLevel,
	log.NoticeLevel,
	log.WarnLevel,
	log.ErrorLevel,
	log.PanicLevel,
	log.AlertLevel,
	log.FatalLevel,
}

var ErrorLevels = []log.Level{
	log.ErrorLevel,
	log.PanicLevel,
	log.AlertLevel,
	log.FatalLevel,
}

var AllLevels = log.AllLevels

var logHandle *console.Console
var udpHandle *UdpHandler
var code bool
//初始本地日志
func InitLog(levels ...log.Level) {
	logHandle = console.New(false)
	logHandle.SetDisplayColor(false)
	logHandle.SetTimestampFormat(time.StampMilli)
	log.AddHandler(logHandle, levels...)
	code = false
}


//初始化网络日志
func InitNetLog(name, ip string, port int, levels ...log.Level) {
	udpHandle = NewUdpHandler()
	udpHandle.SetSvr(name, ip, port)
	udpHandle.SetTimestampFormat(time.StampMilli)
	log.AddHandler(udpHandle, levels...)

}
func SetCode(c bool) {
	code = c
}
var gpath string
func SetPathFilter(path string){
	gpath = path
}

func GetNetWirter() io.Writer {
	if udpHandle != nil {
		return udpHandle.writer
	}
	return nil
}

func format(v ...interface{}) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	s := fmt.Sprint(v...)
	if !code {
		return s
	}
	if gpath != "" {
		//过滤路径
		pos := strings.Index(file, gpath)
		if pos != -1 {
			pos += len(gpath)
		} else {
			pos = 0
		}
		return fmt.Sprintf("%s:%d %s", file[pos+1:], line, s)
	}
	return s
}
func logerformat(v ...interface{}) string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	s := fmt.Sprint(v...)
	if !code {
		return s
	}
	if gpath != "" {
		//过滤路径
		pos := strings.Index(file, gpath)
		if pos != -1 {
			pos += len(gpath)
		} else {
			pos = 0
		}
		return fmt.Sprintf("%s:%d %s", file[pos+1:], line, s)
	}
	return s
}


func Debug(v ...interface{}) { log.Debug(format(v...)) }
func Debuglogger(v ...interface{}) { log.Debug(logerformat(v...)) }
func Debugf(s string, v ...interface{}) { log.Debugf(s, format(v...)) }

func Info(v ...interface{}) { log.Info(format(v...)) }
func Infologger(v ...interface{}) { log.Info(logerformat(v...)) }

func Infof(s string, v ...interface{}) { log.Infof(s, format(v...)) }

func Notice(v ...interface{}) { log.Notice(format(v...)) }
func Noticelogger(v ...interface{}) { log.Notice(logerformat(v...)) }

func Warn(v ...interface{}) { log.Warn(format(v...)) }
func Warnlogger(v ...interface{}) { log.Warn(logerformat(v...)) }

func Warnf(s string, v ...interface{}) { log.Warnf(s, format(v...)) }

func Error(v ...interface{}) { log.Error(format(v...)) }
func Errorlogger(v ...interface{}) { log.Error(logerformat(v...)) }

func Errorf(s string, v ...interface{}) { log.Errorf(s, format(v...)) }

func Panic(v ...interface{}) { log.Panic(format(v...)) }
func Paniclogger(v ...interface{}) { log.Panic(logerformat(v...)) }

func Alert(v ...interface{}) { log.Alert(format(v...)) }
func Alertlogger(v ...interface{}) { log.Alert(logerformat(v...)) }

func Fatal(v ...interface{}) { log.Fatal(format(v...)) }
func Fatallogger(v ...interface{}) { log.Fatal(logerformat(v...)) }

