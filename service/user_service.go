package service

import (
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/dto"
)

type UserService struct{}

func NewUserService() *UserService {
    return &UserService{}
}

// CreateUser handles the creation logic for User
func (this *UserService) CreateUser (input *models.User) error {
    return dto.NewUserDto().CreateUser(input)
}

// GetUser handles the retrieval logic for User
func (this *UserService) GetUser(id string) (*models.User, error) {
    return dto.NewUserDto().GetUser(id)
}

// UpdateUser handles the update logic for User
func (this *UserService) UpdateUser(id string, input *models.User) error {
    return dto.NewUserDto().UpdateUser(id, input)
}

// DeleteUser handles the deletion logic for User
func (this *UserService) DeleteUser (id string) error {
    return dto.NewUserDto().DeleteUser(id)
}
