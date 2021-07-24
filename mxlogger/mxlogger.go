// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxlogger.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	LoggerMaxSize         = 20         //每个日志文件保存的最大尺寸 单位：M
	LoggerMaxBackups      = 30         //日志文件最多保存多少个备份
	LoggerMaxAge          = 30         //文件最多保存多少天
	LoggerCompress        = true       //是否压缩
	DefaultLoggerFilepath = "/var/log" //默认文件存储路径
)

var levelMap = map[string]zapcore.Level{
	"DEBUG":  zapcore.DebugLevel,
	"INFO":   zapcore.InfoLevel,
	"WARN":   zapcore.WarnLevel,
	"ERROR":  zapcore.ErrorLevel,
	"DPANIC": zapcore.DPanicLevel,
	"PANIC":  zapcore.PanicLevel,
	"FATAL":  zapcore.FatalLevel,
}

//创建日志对象
func NewLoggerDefault(日志目录 string, 日志文件名 string, 日志级别 string, 服务名 string) *zap.Logger {
	if len(日志目录) == 0 {
		日志目录 = DefaultLoggerFilepath
	}

	日志文件路径 := 日志目录 + "/" + 日志文件名

	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   日志文件路径,           // 日志文件路径
		MaxSize:    LoggerMaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: LoggerMaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     LoggerMaxAge,     // 文件最多保存多少天
		Compress:   LoggerCompress,   // 是否压缩
	}
	w := zapcore.AddSync(&hook)

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getLoggerLevel(日志级别))

	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",                            // json时时间键
		LevelKey:       "L",                            // json时日志等级键
		NameKey:        "N",                            // json时日志记录器键
		CallerKey:      "C",                            // json时日志文件信息键
		MessageKey:     "M",                            // json时日志消息键
		StacktraceKey:  "S",                            // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,      // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 友好日志等级名大小写（info INFO）
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 友好日志时日期格式化,ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 日志文件信息（包/文件.go:行号）
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, atomicLevel), // 有好的格式、输出控制台、
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), w, atomicLevel),            // json格式、输出文件、处定义等级规则
	)

	logger := zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", 服务名)))
	logger.Info(服务名, zap.String(日志文件路径, "init success"))
	return logger
}

//创建日志对象
func NewLogger(日志目录 string, 日志文件名 string, 日志级别 string, 日志最大文件大小 int, 日志最大文件数 int,
	日志文件最大保存天天数 int, 日志文件是否压缩 bool, 服务名 string) *zap.Logger {
	if len(日志目录) == 0 {
		日志目录 = DefaultLoggerFilepath
	}
	日志文件路径 := 日志目录 + "/" + 日志文件名

	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   日志文件路径,      // 日志文件路径
		MaxSize:    日志最大文件大小,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 日志最大文件数,     // 日志文件最多保存多少个备份
		MaxAge:     日志文件最大保存天天数, // 文件最多保存多少天
		Compress:   日志文件是否压缩,    // 是否压缩
	}
	w := zapcore.AddSync(&hook)

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getLoggerLevel(日志级别))
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",                            // json时时间键
		LevelKey:       "L",                            // json时日志等级键
		NameKey:        "N",                            // json时日志记录器键
		CallerKey:      "C",                            // json时日志文件信息键
		MessageKey:     "M",                            // json时日志消息键
		StacktraceKey:  "S",                            // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,      // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 友好日志等级名大小写（info INFO）
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 友好日志时日期格式化,ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 日志文件信息（包/文件.go:行号）
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, atomicLevel), // 有好的格式、输出控制台、
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), w, atomicLevel),            // json格式、输出文件、处定义等级规则
	)

	logger := zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", 服务名)))
	logger.Info(服务名, zap.String(日志文件路径, "init success"))
	return logger
}

//获取日志级别
func getLoggerLevel(日志级别 string) zapcore.Level {
	if len(日志级别) == 0 {
		日志级别 = "INFO"
	}
	if level, ok := levelMap[日志级别]; ok {
		return level
	}
	return zapcore.InfoLevel
}
