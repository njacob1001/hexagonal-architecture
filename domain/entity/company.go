package entity

// Company struct
type Company struct {
	UUID            string `json:"uuid"`
	NIT             string `json:"string"`
	Creation        string `json:"creation"`
	Employees       uint16 `json:"employees"`
	Chanels         string `json:"chanels"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Sector          string `json:"sector"`
	MainActivity    string `json:"mainActivity"`
	CustomerSegment string `json:"CustomerSegment"`
}
