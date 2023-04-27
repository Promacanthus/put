package log

import (
	"flag"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger is a zap logger instance used for logging.
	Logger *zap.Logger
	// zapcore.level is a type which is used to represent the severity level of a log statement.
	// It is used in conjunction with the zapcore logging library to set the level of logging that will be generated.
	level *zapcore.Level
	// logEncoding is a string used to store an encoding format for a log file.
	logEncoding string
)

func init() {
	level = zap.LevelFlag("level", zap.InfoLevel, "logging level")
	flag.StringVar(&logEncoding, "encoding", "json", "set encoding format `json` or `console` for log file")
	flag.Parse()

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(*level)
	config.Encoding = logEncoding
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = l
}
