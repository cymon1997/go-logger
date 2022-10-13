package logger

import (
	"context"
	"fmt"

	"github.com/cymon1997/go-logger/internal/utils"
	"github.com/sirupsen/logrus"
)

type Entry interface {
	WithContext(ctx context.Context) Entry
	WithMeta(meta Meta) Entry
	WithError(err error) Entry
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

type entryImpl struct {
	logger *logrus.Logger
	ctx    context.Context
	meta   Meta
	err    error
}

func NewEntry(logger *logrus.Logger) Entry {
	return &entryImpl{
		logger: logger,
	}
}

func (e *entryImpl) WithContext(ctx context.Context) Entry {
	e.ctx = ctx
	return e
}

func (e *entryImpl) WithMeta(meta Meta) Entry {
	e.meta = meta
	return e
}

func (e *entryImpl) WithError(err error) Entry {
	e.err = err
	return e
}

func (e *entryImpl) Debug(args ...interface{}) {
	e.log().Debug(args...)
}

func (e *entryImpl) Debugf(format string, args ...interface{}) {
	e.log().Debugf(format, args...)
}

func (e *entryImpl) Info(args ...interface{}) {
	e.log().Info(args...)
}

func (e *entryImpl) Infof(format string, args ...interface{}) {
	e.log().Infof(format, args...)
}

func (e *entryImpl) Warn(args ...interface{}) {
	e.log().Warn(args...)
}

func (e *entryImpl) Warnf(format string, args ...interface{}) {
	e.log().Warnf(format, args...)
}

func (e *entryImpl) Error(args ...interface{}) {
	e.log().Error(args...)
}

func (e *entryImpl) Errorf(format string, args ...interface{}) {
	e.log().Errorf(format, args...)
}

func (e *entryImpl) Fatal(args ...interface{}) {
	e.log().Fatal(args...)
}

func (e *entryImpl) Fatalf(format string, args ...interface{}) {
	e.log().Fatalf(format, args...)
}

func (e *entryImpl) Panic(args ...interface{}) {
	e.log().Panic(args...)
}

func (e *entryImpl) Panicf(format string, args ...interface{}) {
	e.log().Panicf(format, args...)
}

func (e *entryImpl) log() *logrus.Entry {
	fileName, fnName, line := utils.FuncCaller(defaultSkip)
	entry := e.logger.WithField(callerKey, fmt.Sprintf("%s:%d %s", fileName, line, fnName))
	if e.ctx != nil {
		entry = entry.WithField(string(config.IDKey), utils.GetStrValueFromCtx(e.ctx, config.IDKey))
	}
	if e.meta != nil {
		entry = entry.WithFields(logrus.Fields(e.meta))
	}
	if e.err != nil {
		entry = entry.WithError(e.err)
	}
	return entry
}
