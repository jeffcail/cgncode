package dto

import (
	"errors"
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/vm"
)

type InfoDto struct{}

func NewInfoDto() *InfoDto {
	return &InfoDto{}
}

// CreateInfo handles database insertion for Info
func (this *InfoDto) CreateInfo(bean *models.Info) error {
	affected, err := vm.Db.Insert(bean)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("insert affected 0 rows")
	}

	return nil
}

// GetInfo handles database retrieval for Info
func (this *InfoDto) GetInfo(id string) (*models.Info, error) {
	var bean = new(models.Info)
	has, err := vm.Db.Id(id).Get(bean)
	if err != nil {
		return bean, err
	}
	if !has {
		return bean, errors.New("data is not found")
	}
	return bean, nil
}

// UpdateInfo handles database update for Info
func (this *InfoDto) UpdateInfo(id string, bean *models.Info) error {
	_, err := vm.Db.Id(bean.ID).AllCols().Update(bean)
	if err != nil {
		return err
	}
	return nil
}

// DeleteInfo handles database deletion for Info
func (this *InfoDto) DeleteInfo(id string) error {
	affected, err := vm.Db.Id(id).Delete(&models.Info{})
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("delete affected 0 rows")
	}
	return errors.New("not implemented")
}
