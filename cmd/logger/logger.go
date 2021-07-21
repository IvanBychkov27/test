package main

import (
	"gitlab.com/adscompass/logger"
	"go.uber.org/zap"
)

type TaggedLogger interface {
	Zap() *zap.Logger
	TaggedInfo(tags string, msg string, fields ...zap.Field)
}

func main() {
	var loggerOptions []logger.Option
	loggerOptions = append(loggerOptions, logger.WithDevelopment())
	l := logger.New(loggerOptions...)
	l.Zap().Info("logger")

	l.TagOn("tags.1")
	l.TagOn("tags.2")
	//l.TagsOnAll()

	l.TaggedInfo("tags.1", "tags 1", zap.Int("int", 10))
	//l.TagOff("tags.1")
	l.TaggedInfo("tags.1", "tags 1", zap.Int("int", 20))

	//l.TagsOffAll()

	l.TaggedInfo("tags.2", "tags 2", zap.Int("int", 100))
	//l.TagOff("tags.2")
	l.TaggedInfo("tags.2", "tags 2", zap.Int("int", 200))

}
