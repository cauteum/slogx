package slogx

import "time"

// Operational defaults owned by this package.
const (
	controlHeaderReadTimeout = 5 * time.Second
	controlShutdownTimeout   = 3 * time.Second
	defaultLevelPollInterval = 2 * time.Second
)
