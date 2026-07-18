package models

import "time"

type Track struct {
	Title    string
	Artists  []string
	Album    string
	Duration time.Duration
	ISRC     string
}
