package mzap

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _level zapcore.Level

var Glog *zap.Logger

type ZapConfig struct {
	Level         string // 级别
	Format        string // 输出
	Prefix        string // 日志前缀
	Director      string // 日志文件夹
	LinkName      string // 软链接名称
	ShowLine      bool   // 显示行
	EncodeLevel   string // 编码级
	StacktraceKey string // 栈名
	LogInConsole  bool   // 输出控制台
}

var _zapConfig = &ZapConfig{
	Level:         "info",
	Format:        "console",
	Prefix:        "[mzap]",
	Director:      "log",
	LinkName:      "latest_log",
	ShowLine:      true,
	EncodeLevel:   "LowercaseColorLevelEncoder",
	StacktraceKey: "stacktrace",
	LogInConsole:  true,
}

// mzap.InitZap(&mzap.ZapConfig{
// 	Level:         "error",
// 	Format:        "console",
// 	Prefix:        "[mzap]",
// 	Director:      "log",
// 	LinkName:      "latest_log",
// 	ShowLine:      true,
// 	EncodeLevel:   "LowercaseColorLevelEncoder",
// 	StacktraceKey: "stacktrace",
// 	LogInConsole:  true,
// })
func InitZap(zapConfig *ZapConfig) {
	if zapConfig != nil {
		_zapConfig = zapConfig
	}
	dir := _zapConfig.Director
	if ok, _ := pathExists(dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}

	switch _zapConfig.Level { // 初始化配置文件的Level
	case "debug":
		_level = zap.DebugLevel
	case "info":
		_level = zap.InfoLevel
	case "warn":
		_level = zap.WarnLevel
	case "error":
		_level = zap.ErrorLevel
	case "dpanic":
		_level = zap.DPanicLevel
	case "panic":
		_level = zap.PanicLevel
	case "fatal":
		_level = zap.FatalLevel
	default:
		_level = zap.InfoLevel
	}

	if _level == zap.DebugLevel || _level == zap.ErrorLevel {
		Glog = zap.New(getEncoderCore(), zap.AddStacktrace(_level))
	} else {
		Glog = zap.New(getEncoderCore())
	}
	if _zapConfig.ShowLine {
		Glog = Glog.WithOptions(zap.AddCaller())
	}
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  _zapConfig.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case _zapConfig.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case _zapConfig.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case _zapConfig.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case _zapConfig.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if _zapConfig.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(_zapConfig.Director, "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if _zapConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, _level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(_zapConfig.Prefix + "2006/01/02 - 15:04:05.000"))
}

//pathExists 判断一个文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
