package entity

import "time"

// Meeting struct
type Meeting struct {
	UUID         string        `json:"uuid"`
	Subject      string        `json:"subject"`
	Description  string        `json:"description"`
	Entrepreneur Entrepreneur  `json:"entrepreneurs"`
	Facilitator  Facilitator   `json:"Facilitator"`
	Start        time.Time     `json:"start"`
	End          time.Time     `json:"End"`
	Report       MeetingReport `json:"report"`
}
