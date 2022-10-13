package logger

import (
	"context"
)

func WithContext(ctx context.Context) Entry {
	return NewEntry(logger).WithContext(ctx)
}

func WithMeta(meta Meta) Entry {
	return NewEntry(logger).WithMeta(meta)
}

func WithError(err error) Entry {
	return NewEntry(logger).WithError(err)
}

func Debug(args ...interface{}) {
	NewEntry(logger).Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	NewEntry(logger).Debugf(format, args...)
}

func Info(args ...interface{}) {
	NewEntry(logger).Info(args...)
}

func Infof(format string, args ...interface{}) {
	NewEntry(logger).Infof(format, args...)
}

func Warn(args ...interface{}) {
	NewEntry(logger).Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	NewEntry(logger).Warnf(format, args...)
}

func Error(args ...interface{}) {
	NewEntry(logger).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	NewEntry(logger).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	NewEntry(logger).Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	NewEntry(logger).Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	NewEntry(logger).Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	NewEntry(logger).Panicf(format, args...)
}
