package logger

import "testing"

type level int

const (
	info level = iota
	err
	fatal
)

var levelNames = map[level]string{
	info:  "INFO",
	err:   "ERROR",
	fatal: "FATAL",
}

func log(t *testing.T, l level, msg string, args ...any) {
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
		t.Logf("["+levelNames[l]+"] "+msg, args...)
	})
}

func LogInfo(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, info, msg, args...)
}

func LogError(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, err, msg, args...)
	t.Fail()
}

func LogFatal(t *testing.T, msg string, args ...any) {
	t.Helper()

	log(t, fatal, msg, args...)
	t.FailNow()
}
