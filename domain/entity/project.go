package entity

// Project struct
type Project struct {
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Sector      string  `json:"sector"`
	Status      string  `json:"status"`
	Company     Company `json:"company"`
}

// Projects array
type Projects []Project
