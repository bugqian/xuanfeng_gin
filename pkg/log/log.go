package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config .
type Config struct {
	Level       string `yaml:"level"`        // 日志等级 debug info warn error dpanic panic fatal
	AddCaller   bool   `yaml:"add-caller"`   // 开启行号
	WriteStdout bool   `yaml:"write-stdout"` // 是否写入到控制台
	Filename    string `yaml:"filename"`     // 日志文件路径，为空则不写入文件
	MaxSize     int    `yaml:"max-size"`     // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups  int    `yaml:"max-backups"`  // 日志文件最多保存多少个备份
	MaxAge      int    `yaml:"max-age"`      // 文件最多保存多少天
	Compress    bool   `yaml:"compress"`     // 是否压缩
}

var L *zap.Logger

// Init .
func Init(conf *Config) {
	hook := lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	switch conf.Level {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	case "dpanic":
		atomicLevel.SetLevel(zap.DPanicLevel)
	case "panic":
		atomicLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		atomicLevel.SetLevel(zap.FatalLevel)
	default:
		atomicLevel.SetLevel(zap.WarnLevel)
	}

	zapWS := []zapcore.WriteSyncer{}
	if conf.Filename != "" {
		zapWS = append(zapWS, zapcore.AddSync(&hook))
	}
	if conf.WriteStdout {
		zapWS = append(zapWS, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(zapWS...),
		atomicLevel,
	)

	hostName, _ := os.Hostname()
	zapOption := []zap.Option{
		zap.Fields(zap.String("HostName", hostName)), // 设置初始化字段
	}

	if conf.AddCaller {
		zapOption = append(zapOption, zap.AddCaller())
	}

	if conf.Level == "debug" {
		// 开启开发模式
		zapOption = append(zapOption, zap.Development())
	}
	L = zap.New(core, zapOption...)
}
