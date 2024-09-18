package service

import (
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/dto"
)

type GoodsService struct{}

func NewGoodsService() *GoodsService {
    return &GoodsService{}
}

// CreateGoods handles the creation logic for Goods
func (this *GoodsService) CreateGoods (input *models.Goods) error {
    return dto.NewGoodsDto().CreateGoods(input)
}

// GetGoods handles the retrieval logic for Goods
func (this *GoodsService) GetGoods(id string) (*models.Goods, error) {
    return dto.NewGoodsDto().GetGoods(id)
}

// UpdateGoods handles the update logic for Goods
func (this *GoodsService) UpdateGoods(id string, input *models.Goods) error {
    return dto.NewGoodsDto().UpdateGoods(id, input)
}

// DeleteGoods handles the deletion logic for Goods
func (this *GoodsService) DeleteGoods (id string) error {
    return dto.NewGoodsDto().DeleteGoods(id)
}
