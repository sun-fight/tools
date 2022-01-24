package mzap_test

import (
	"github.com/sun-fight/tools/mzap"
	"testing"
)

func Test(t *testing.T) {
	mzap.InitZap(&mzap.ZapConfig{
		Level:         "error",
		Format:        "console",
		Prefix:        "[mzap]",
		Director:      "log",
		LinkName:      "latest_log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	})
	mzap.Glog.Error("test err")
}
