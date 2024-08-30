package dto

import (
    "errors"
    "github.com/jeffcail/cgncode/models"
    "github.com/jeffcail/cgncode/vm"
)

type UserDto struct{}

func NewUserDto() *UserDto {
    return &UserDto{}
}

// CreateUser handles database insertion for User
func (this *UserDto) CreateUser(bean *models.User) error {
    affected, err := vm.Db.Insert(bean)
    if err != nil {
        return err
    }
    if affected == 0 {
        return errors.New("insert affected 0 rows")
    }

    return nil
}

// GetUser handles database retrieval for User
func (this *UserDto) GetUser(id string) (*models.User, error) {
    var bean = new(models.User)
    has, err := vm.Db.Id(id).Get(bean)
    if err != nil {
        return bean, err
    }
    if !has {
        return bean, errors.New("data is not found")
    }
    return bean, nil
}

// UpdateUser handles database update for User
func (this *UserDto) UpdateUser(id string, bean *models.User) error {
    _, err := vm.Db.Id(bean.ID).AllCols().Update(bean)
    if err != nil {
        return err
    }
    return nil
}

// DeleteUser handles database deletion for User
func (this *UserDto) DeleteUser(id string) error {
    affected, err := vm.Db.Id(id).Delete(&models.User{})
    if err != nil {
        return err
    }

    if affected == 0 {
        return errors.New("delete affected 0 rows")
    }
    return errors.New("not implemented")
}