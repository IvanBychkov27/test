package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func WithSamplingConfig(samplingConfig *zap.SamplingConfig) Option {
	return func(l *Logger) {
		l.samplingConfig = samplingConfig
	}
}

func WithWriteSyncer(ws zapcore.WriteSyncer) Option {
	return func(l *Logger) {
		l.ws = ws
	}
}

func WithFieldsKeys(fieldsKeys FieldsKeys) Option {
	return func(l *Logger) {
		l.fieldsKeys = fieldsKeys
	}
}

func WithDevelopment() Option {
	return func(l *Logger) {
		l.development = true
		l.atomicLevel.SetLevel(zap.DebugLevel)
	}
}
