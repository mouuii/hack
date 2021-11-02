package log

import "testing"

func TestInfo(t *testing.T) {
	logger := DefaultLogger
	logger = With(logger, "ts", DefaultTimestamp, "caller", Caller(3))
	_ = logger.Log(LevelInfo, "key1", "value1")
}
