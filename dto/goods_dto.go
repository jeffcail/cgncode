package dto

import (
    "errors"
    "github.com/jeffcail/cgncode/models"
    "github.com/jeffcail/cgncode/vm"
)

type GoodsDto struct{}

func NewGoodsDto() *GoodsDto {
    return &GoodsDto{}
}

// CreateGoods handles database insertion for Goods
func (this *GoodsDto) CreateGoods(bean *models.Goods) error {
    affected, err := vm.Db.Insert(bean)
    if err != nil {
        return err
    }
    if affected == 0 {
        return errors.New("insert affected 0 rows")
    }

    return nil
}

// GetGoods handles database retrieval for Goods
func (this *GoodsDto) GetGoods(id string) (*models.Goods, error) {
    var bean = new(models.Goods)
    has, err := vm.Db.Id(id).Get(bean)
    if err != nil {
        return bean, err
    }
    if !has {
        return bean, errors.New("data is not found")
    }
    return bean, nil
}

// UpdateGoods handles database update for Goods
func (this *GoodsDto) UpdateGoods(id string, bean *models.Goods) error {
    _, err := vm.Db.Id(bean.ID).AllCols().Update(bean)
    if err != nil {
        return err
    }
    return nil
}

// DeleteGoods handles database deletion for Goods
func (this *GoodsDto) DeleteGoods(id string) error {
    affected, err := vm.Db.Id(id).Delete(&models.Goods{})
    if err != nil {
        return err
    }

    if affected == 0 {
        return errors.New("delete affected 0 rows")
    }
    return errors.New("not implemented")
}