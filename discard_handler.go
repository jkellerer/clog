package clog

import "errors"

var (
	// errorDiscarded is sent when using the Discard handler
	errorDiscarded = errors.New("this message is not going anywhere")
)

// DiscardHandler forgets any log message
type DiscardHandler struct{}

// LogEntry discards the LogEntry
func (l *DiscardHandler) LogEntry(LogEntry) error {
	return errorDiscarded
}

// Verify interface
var (
	_ Handler = &DiscardHandler{}
)
