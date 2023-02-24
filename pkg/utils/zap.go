package utils

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	level zapcore.Level
	Conf  Config
)

func Zap(C Config) (logger *zap.Logger) {
	Conf = C
	if ok, _ := PathExists(Conf.Viper.GetString("zap.Director")); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", Conf.Viper.GetString("Director"))
		_ = os.Mkdir(Conf.Viper.GetString("zap.Director"), os.ModePerm)
	}

	switch Conf.Viper.GetString("zap.Level") { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	logger = &zap.Logger{}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if Conf.Viper.GetBool("zap.ShowLine") {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  Conf.Viper.GetString("zap.StacktraceKey"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case Conf.Viper.GetString("zap.EncodeLevel") == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case Conf.Viper.GetString("zap.EncodeLevel") == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case Conf.Viper.GetString("zap.EncodeLevel") == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case Conf.Viper.GetString("zap.EncodeLevel") == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if Conf.Viper.GetString("zap.Format") == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(Conf.Viper.GetString("zap.Prefix") + "2006/01/02 - 15:04:05.000"))
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(Conf.Viper.GetString("zap.Director"), "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if Conf.Viper.GetBool("zap.LogInConsole") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
