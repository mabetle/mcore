package main

import (
	"github.com/mabetle/mlog"
)

var (
	logger = mlog.GetLogger("main")
)


func main() {
	mlog.SetTraceLevel()

	logger.Trace("Hello")
	logger.Debug("Hello")
	logger.Info("Hello")
	logger.Warn("Hello")
	logger.Error("Hello")
}
