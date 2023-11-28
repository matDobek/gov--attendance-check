package logger

import (
	"fmt"
)

type level int

const (
	debug level = iota
	info
	warn
	err // don't overshadow `error`
	fatal
)

var levelNames = map[level]string{
	debug: "DEBUG",
	info:  "INFO",
	warn:  "WARN",
	err:   "ERROR",
	fatal: "FATAL",
}

func log(l level, msg string, args ...any) {
	fmt.Printf("["+levelNames[l]+"] "+msg+"\n", args...)
}

func Debug(msg string, args ...any) {
	log(debug, msg, args...)
}

func Info(msg string, args ...any) {
	log(info, msg, args...)
}

func Warn(msg string, args ...any) {
	log(warn, msg, args...)
}

func Error(e error) {
	log(err, e.Error())
}

func Fatal(e error) {
	log(fatal, e.Error())
}
