package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

const (
	TagsSeparator = "|"
)

type FieldsKeys struct {
	Time       string
	Level      string
	Name       string
	Caller     string
	Message    string
	Stacktrace string
}

type Logger struct {
	z *zap.Logger

	ws             zapcore.WriteSyncer
	fieldsKeys     FieldsKeys
	samplingConfig *zap.SamplingConfig
	atomicLevel    zap.AtomicLevel

	tagsMx  sync.RWMutex
	tags    map[string]struct{}
	tagsAll bool

	development bool
}

type Option func(*Logger)

func New(opts ...Option) *Logger {
	l := &Logger{
		tags: map[string]struct{}{},
		ws:   os.Stderr,
		fieldsKeys: FieldsKeys{
			Time:       "ts",
			Level:      "level",
			Name:       "logger",
			Caller:     "caller",
			Message:    "msg",
			Stacktrace: "stacktrace",
		},
		atomicLevel: zap.NewAtomicLevelAt(zap.InfoLevel),
	}

	for _, o := range opts {
		o(l)
	}

	var encoder zapcore.Encoder

	switch l.development {
	case true:
		encoder = developmentEncoder(l.fieldsKeys)
	default:
		encoder = productionEncoder(l.fieldsKeys)
	}

	core := zapcore.NewCore(encoder, l.ws, l.atomicLevel)

	if l.samplingConfig != nil {
		core = zapcore.NewSampler(core, time.Second, int(l.samplingConfig.Initial), int(l.samplingConfig.Thereafter))
	}

	l.z = zap.New(core)

	return l
}

func productionEncoder(fieldsKeys FieldsKeys) zapcore.Encoder {
	zapEncoder := zapcore.EncoderConfig{
		TimeKey:        fieldsKeys.Time,
		LevelKey:       fieldsKeys.Level,
		NameKey:        fieldsKeys.Name,
		CallerKey:      fieldsKeys.Caller,
		MessageKey:     fieldsKeys.Message,
		StacktraceKey:  fieldsKeys.Stacktrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewJSONEncoder(zapEncoder)
}

func developmentEncoder(fieldsKeys FieldsKeys) zapcore.Encoder {
	zapEncoder := zapcore.EncoderConfig{
		TimeKey:        fieldsKeys.Time,
		LevelKey:       fieldsKeys.Level,
		NameKey:        fieldsKeys.Name,
		CallerKey:      fieldsKeys.Caller,
		MessageKey:     fieldsKeys.Message,
		StacktraceKey:  fieldsKeys.Stacktrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(zapEncoder)
}
