package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

type AriesLoggerFields map[string]interface{}

type AriesLoggerEntry struct {
	context  context.Context
	stdEntry *logrus.Entry
	Data     AriesLoggerFields
	err      string
}

func newAriesLoggerEntry(ctx context.Context, stdEntry *logrus.Entry) *AriesLoggerEntry {
	return &AriesLoggerEntry{context: ctx, stdEntry: stdEntry, Data: make(AriesLoggerFields, 6), err: ""}
}

func NewloggerEntry() *AriesLoggerEntry {
	ctx := context.TODO()
	return newAriesLoggerEntry(ctx, logrus.StandardLogger().WithContext(ctx))
}
