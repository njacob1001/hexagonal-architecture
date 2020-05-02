package entity

// Activity struct
type Activity struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Description string `json:"description"`
	Attachments Attachments `json:"attachments"`
}
