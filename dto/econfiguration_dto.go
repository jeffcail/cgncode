package dto

import (
    "errors"
    "github.com/jeffcail/cgncode/models"
    "github.com/jeffcail/cgncode/vm"
)

type EConfigurationDto struct{}

func NewEConfigurationDto() *EConfigurationDto {
    return &EConfigurationDto{}
}

// CreateEConfiguration handles database insertion for EConfiguration
func (this *EConfigurationDto) CreateEConfiguration(bean *models.EConfiguration) error {
    affected, err := vm.Db.Insert(bean)
    if err != nil {
        return err
    }
    if affected == 0 {
        return errors.New("insert affected 0 rows")
    }

    return nil
}

// GetEConfiguration handles database retrieval for EConfiguration
func (this *EConfigurationDto) GetEConfiguration(id string) (*models.EConfiguration, error) {
    var bean = new(models.EConfiguration)
    has, err := vm.Db.Id(id).Get(bean)
    if err != nil {
        return bean, err
    }
    if !has {
        return bean, errors.New("data is not found")
    }
    return bean, nil
}

// UpdateEConfiguration handles database update for EConfiguration
func (this *EConfigurationDto) UpdateEConfiguration(id string, bean *models.EConfiguration) error {
    _, err := vm.Db.Id(bean.ID).AllCols().Update(bean)
    if err != nil {
        return err
    }
    return nil
}

// DeleteEConfiguration handles database deletion for EConfiguration
func (this *EConfigurationDto) DeleteEConfiguration(id string) error {
    affected, err := vm.Db.Id(id).Delete(&models.EConfiguration{})
    if err != nil {
        return err
    }

    if affected == 0 {
        return errors.New("delete affected 0 rows")
    }
    return errors.New("not implemented")
}