package entity

// Attachment struct
type Attachment struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	MimeType string `json:"mimeType"`
}

// Attachments array
type Attachments []Attachment
