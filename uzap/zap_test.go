package uzap_test

import (
	"github.com/sun-fight/tools/uzap"
	"testing"
)

func Test(t *testing.T) {
	uzap.InitZap(&uzap.ZapConfig{
		Level:         "error",
		Format:        "console",
		Prefix:        "[uzap]",
		Director:      "log",
		LinkName:      "latest_log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	})
	uzap.Glog.Error("test err")
}
