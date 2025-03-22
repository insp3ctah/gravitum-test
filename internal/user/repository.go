package user

import (
	"gravitum-test/pkg"
	"log"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUser(user *User) error {
	return pkg.DB.Create(user).Error
}

func (r *Repository) GetUser(id uint) (*User, error) {
	user := &User{}
	err := pkg.DB.First(user, id).Error
	if err != nil {
		log.Printf("User %d not found\n", id)
		return nil, err
	}
	return user, err
}

func (r *Repository) UpdateUser(user *User) error {
	return pkg.DB.Save(user).Error
}
