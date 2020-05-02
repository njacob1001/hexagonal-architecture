package entity

// Person struct
type Person struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// Faculty struct
type Faculty struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Dean Person `json:"dean"`
}

// Career struct
type Career struct {
	UUID     string  `json:"uuid"`
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Director Person  `json:"director"`
	Faculty  Faculty `json:"faculty"`
}
