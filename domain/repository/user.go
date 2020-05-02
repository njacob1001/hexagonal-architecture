package repository

import "github.com/njacob1001/sinapsis-back/domain/entity"

// UserRepository interface
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	DeleteUser(uuid string) error
	UpdateUser(*entity.User) (*entity.User, error)
}
