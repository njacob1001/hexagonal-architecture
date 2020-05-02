package entity

import "time"

// Schedule struct
type Schedule struct {
	UUID          string     `json:"uuid"`
	BusyIndicator bool       `json:"busyIndicator"`
	Date          time.Time  `json:"date"`
	TimeTables    TimeTables `json:"timeTables"`
}

// TimeTable struct
type TimeTable struct {
	UUID  string    `json:"time"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// TimeTables array
type TimeTables []TimeTable
