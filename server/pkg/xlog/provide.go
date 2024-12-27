package xlog

import (
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// ProvideZapLogger
// https://github.com/uber-go/zap
func ProvideZapLogger(conf *viper.Viper, disk string, debug bool) (*zap.Logger, error) {
	//var err error
	// lp := conf.GetString("local.log_path")
	//fn := conf.GetString("logger." + disk + ".name")
	//if fn == "" {
	//	return nil, errors.New("log disk(" + disk + ") not exist")
	//}

	//filename := path.Join(lp, fn)
	//if strings.Index(filename, ".") == 0 {
	//	filename, err = filepath.Abs(filename)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	filename := conf.GetString("logger." + disk + ".filename")
	if filename == "" {
		return nil, errors.New("log disk(" + disk + ") not exist")
	}

	fl := fileLogger(
		filename,
		conf.GetInt("logger."+disk+".max_size"),
		conf.GetInt("logger."+disk+".max_backups"),
		conf.GetInt("logger."+disk+".max_age"),
		conf.GetBool("logger."+disk+".compress"),
	)

	var ws []zapcore.WriteSyncer
	ws = append(ws, zapcore.AddSync(fl))
	//debug := conf.GetBool("app.debug")
	if debug {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		newEncoder(conf.GetString("logger."+disk+".encoding")),
		zapcore.NewMultiWriteSyncer(
			ws...,
		), // Print to console and file
		parseLevel(conf.GetString("logger."+disk+".level")),
	)

	//if conf.GetString("app.env") != "prod" {
	//	return zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)), nil
	//}
	return zap.New(
		core,
		// zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCaller(),
	), nil
}

func fileLogger(filename string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,   // Log file path
		MaxSize:    maxSize,    // Maximum size unit for each log file: M
		MaxBackups: maxBackups, // The maximum number of backups that can be saved for log files
		MaxAge:     maxAge,     // Maximum number of days the file can be saved
		Compress:   compress,   // Compression or not
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
	//enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

func parseLevel(lv string) zapcore.Level {
	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}
	return level
}

func newEncoder(encoding string) zapcore.Encoder {
	if encoding == "json" {
		return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			// EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}

	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}
