package main

import (
	"context"
	"errors"

	"github.com/cymon1997/go-logger"
)

func main() {
	// Define IDKey
	var (
		Key logger.IDKey = "context_id"
	)

	type Input struct {
		A string `json:"a"`
		B int    `json:"b"`
	}

	// Init client
	logger.Init(logger.Config{
		Format:     logger.TextFormat,
		Level:      logger.InfoLevel,
		IsBeautify: false,
		IDKey:      Key,
	})

	// Set IDKey
	ctx := context.WithValue(context.Background(), Key, "50bc28f9-d9bb-41f4-83b3-3329797fb6a5")

	// Debug level
	logger.WithContext(ctx).WithMeta(logger.Meta{
		"input": Input{
			A: "param",
			B: 123,
		},
		"result": "data",
	}).Debug("execution result")

	// Info level
	logger.WithContext(ctx).Infof("execution success: %d results", 123)

	// Warn level
	err := errors.New("safe to ignore")
	logger.WithContext(ctx).WithError(err).Warn("executed with warn")

	// Error level
	err = errors.New("some errors")
	logger.WithContext(ctx).WithError(err).Error("error execution")

	// Fatal level
	err = errors.New("unknown failure")
	logger.WithContext(ctx).WithError(err).Fatal("exception occurred")

	// Panic level
	err = errors.New("exception with trace")
	logger.WithContext(ctx).WithError(err).Panic("application crash!")
}
