package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

func (l *Logger) SetLevel(level zapcore.Level) {
	l.atomicLevel.SetLevel(level)
}

func (l *Logger) tagged(tags string, level zapcore.Level, msg string, fields ...zap.Field) {
	if l.tagsAll {
		l.z.Check(level, msg).Write(fields...)
		return
	}

	modeOnlyAllTags := false
	if strings.HasPrefix(tags, "!") {
		tags = tags[1:]
		modeOnlyAllTags = true
	}

	l.tagsMx.RLock()
	defer l.tagsMx.RUnlock()

	for _, tag := range strings.Split(tags, TagsSeparator) {
		_, ok := l.tags[tag]
		if !ok && modeOnlyAllTags {
			return
		}
		if ok && !modeOnlyAllTags {
			l.z.Check(level, msg).Write(fields...)
			return
		}
	}

	if modeOnlyAllTags {
		l.z.Check(level, msg).Write(fields...)
	}
}

func (l *Logger) TaggedError(tags string, msg string, fields ...zap.Field) {
	l.tagged(tags, zap.ErrorLevel, msg, fields...)
}

func (l *Logger) TaggedWarn(tags string, msg string, fields ...zap.Field) {
	l.tagged(tags, zap.WarnLevel, msg, fields...)
}

func (l *Logger) TaggedInfo(tags string, msg string, fields ...zap.Field) {
	l.tagged(tags, zap.InfoLevel, msg, fields...)
}

func (l *Logger) TaggedDebug(tags string, msg string, fields ...zap.Field) {
	l.tagged(tags, zap.DebugLevel, msg, fields...)
}

func (l *Logger) Zap() *zap.Logger {
	return l.z
}

func (l *Logger) TagOn(name string) {
	l.tagsMx.Lock()
	defer l.tagsMx.Unlock()

	l.tags[name] = struct{}{}
}

func (l *Logger) TagOff(name string) {
	l.tagsMx.Lock()
	defer l.tagsMx.Unlock()

	l.tagsAll = false

	delete(l.tags, name)
}

func (l *Logger) TagsOnAll() {
	l.tagsMx.Lock()
	defer l.tagsMx.Unlock()

	l.tagsAll = true
}

func (l *Logger) TagsOffAll() {
	l.tagsMx.Lock()
	defer l.tagsMx.Unlock()

	l.tagsAll = false

	for name := range l.tags {
		delete(l.tags, name)
	}
}

func (l *Logger) Tags() []string {
	l.tagsMx.Lock()
	defer l.tagsMx.Unlock()

	tags := make([]string, 0)

	for name := range l.tags {
		tags = append(tags, name)
	}

	return tags
}
