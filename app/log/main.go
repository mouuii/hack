package main

import "github/mouuii/hack/pkg/log"

func main() {
	logger := log.DefaultLogger
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.Caller2(3))
	_ = logger.Log(log.LevelInfo, "key1", "value1")
}
