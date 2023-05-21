package logext

import (
	"fmt"
	"github.com/ungame/go-signup/utils"
	"io"
	"log"
	"os"
)

type Logger interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Fatal(format string, v ...any)
	Panic(format string, v ...any)
	Close()
}

var defaultLogger logger

type logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	file  *os.File
}

const (
	debugPrefix = "[DEBUG]\t"
	infoPrefix  = "[INFO] \t"
	warnPrefix  = "[WARN] \t"
	errorPrefix = "[ERROR]\t"

	logFileName = "log.log"
)

func New(name string) (Logger, error) {
	var (
		fileName  = fmt.Sprintf("%s.log", name)
		newLogger = &logger{}
	)
	var err error
	newLogger.file, err = os.Create(normalizeOutput(fileName))
	if err != nil {
		return nil, err
	}
	multiWriter := io.MultiWriter(newLogger.file)
	newLogger.debug = log.New(multiWriter, debugPrefix, log.LstdFlags)
	newLogger.info = log.New(multiWriter, infoPrefix, log.LstdFlags)
	newLogger.warn = log.New(multiWriter, warnPrefix, log.LstdFlags)
	newLogger.error = log.New(multiWriter, errorPrefix, log.LstdFlags)
	return newLogger, nil
}

func init() {
	var err error
	defaultLogger.file, err = os.Create(normalizeOutput(logFileName))
	if err != nil {
		log.Fatalln(err)
	}
	multiWriter := io.MultiWriter(defaultLogger.file)
	defaultLogger.debug = log.New(multiWriter, debugPrefix, log.LstdFlags)
	defaultLogger.info = log.New(multiWriter, infoPrefix, log.LstdFlags)
	defaultLogger.warn = log.New(multiWriter, warnPrefix, log.LstdFlags)
	defaultLogger.error = log.New(multiWriter, errorPrefix, log.LstdFlags)
}

func (l *logger) Close() {
	utils.DefaultCloser.Close(l.file)
}

func (l *logger) Debug(format string, v ...any) {
	l.debug.Printf(withNewLine(format), v...)
}

func (l *logger) Info(format string, v ...any) {
	l.info.Printf(withNewLine(format), v...)
}

func (l *logger) Warn(format string, v ...any) {
	l.warn.Printf(withNewLine(format), v...)
}

func (l *logger) Error(format string, v ...any) {
	l.error.Printf(withNewLine(format), v...)
}

func (l *logger) Fatal(format string, v ...any) {
	log.Fatalf(withNewLine(format), v...)
}

func (l *logger) Panic(format string, v ...any) {
	log.Panicf(withNewLine(format), v...)
}

func withNewLine(s string) string {
	return s + "\n"
}

func Close() {
	if defaultLogger.file == nil {
		return
	}
	err := defaultLogger.file.Close()
	if err != nil {
		log.Println("error on close logger:", err.Error())
	}
}

func Debug(format string, v ...any) {
	defaultLogger.Debug(format, v...)
}

func Info(format string, v ...any) {
	defaultLogger.Info(format, v...)
}

func Warn(format string, v ...any) {
	defaultLogger.Warn(format, v...)
}

func Error(format string, v ...any) {
	defaultLogger.Error(format, v...)
}

func Fatal(format string, v ...any) {
	defaultLogger.Fatal(format, v...)
}

func Panic(format string, v ...any) {
	defaultLogger.Panic(format, v...)
}

type noOpLogger struct{}

func NewNop() Logger {
	return &noOpLogger{}
}

func (n noOpLogger) Debug(string, ...any) {

}

func (n noOpLogger) Info(string, ...any) {

}

func (n noOpLogger) Warn(string, ...any) {

}

func (n noOpLogger) Error(string, ...any) {

}

func (n noOpLogger) Fatal(string, ...any) {

}

func (n noOpLogger) Panic(string, ...any) {

}

func (n noOpLogger) Close() {

}
