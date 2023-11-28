package logger

import "testing"

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

func log(t *testing.T, level Level, msg string, args ...any) {
	// Clump together all the logs for given test case
	// Usefull for parallel tests, to switch
	// from:
	//		[info] test_1
	//		[info] test_2.1
	//		[info] test_3
	//		[info] test_2.2
	// to:
	//		[info] test_1
	//		[info] test_3
	//		[info] test_2.1
	//		[info] test_2.2
	t.Cleanup(func() {
		t.Logf("["+levelNames[level]+"] "+msg, args...)
	})
}

func LogInfo(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, INFO, msg, args...)
}

func LogError(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, ERROR, msg, args...)
	t.Fail()
}

func LogFatal(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, FATAL, msg, args...)
	t.FailNow()
}
