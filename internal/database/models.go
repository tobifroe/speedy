package database

import "time"

type TestResult struct {
	// Download duration in seconds
	Duration float64
	// Download speed in Megabytes per second
	DownSpeed float64
	Target    string
	CreatedAt time.Time
}

type Config struct {
	Schedule string
}
