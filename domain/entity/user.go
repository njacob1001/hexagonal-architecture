package entity

import "github.com/google/uuid"

// User struct
type User struct {
	UUID           string `json:"UUID"`
	Names          string `json:"names"`
	LastNames      string `json:"lastNames"`
	Identification string `json:"identification"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Photo          string `json:"photo"`
	Address        string `json:"address"`
}

// Administrator struct
type Administrator struct {
	User User   `json:"user"`
	UUID string `json:"uuid"`
}

// Administrators array
type Administrators []Administrator

// Entrepreneur struct
type Entrepreneur struct {
	User       User   `json:"user"`
	UUID       string `json:"uuid"`
	Code       string `json:"code"`
	Profession string `json:"profession"`
	Career     Career `json:"career"`
}

// Entrepreneurs array
type Entrepreneurs []Entrepreneur

// Facilitator struct
type Facilitator struct {
	User       User   `json:"user"`
	UUID       string `jsobn:"uuid"`
	Profession string `json:"profession"`
}

// Facilitators array
type Facilitators []Facilitator

// NewUser function
func NewUser(names string, lastNames string, id string, phone string, email string, pass string, photo string, address string) *User {
	return &User{
		UUID:           uuid.New().String(),
		Names:          names,
		LastNames:      lastNames,
		Identification: id,
		Phone:          phone,
		Email:          email,
		Password:       pass,
		Photo:          photo,
		Address:        address,
	}
}

// NewEntrepreteur function
func NewEntrepreteur(user User, code string, profession string, career Career) *Entrepreneur {
	return &Entrepreneur{
		UUID:       uuid.New().String(),
		User:       user,
		Profession: profession,
		Career:     career,
	}
}

// NewFacilitator function
func NewFacilitator(user User, profession string) *Facilitator {
	return &Facilitator{
		User:       user,
		UUID:       uuid.New().String(),
		Profession: profession,
	}
}

// NewAdministrator function
func NewAdministrator(user User) *Administrator {
	return &Administrator{
		UUID: uuid.New().String(),
		User: user,
	}
}
