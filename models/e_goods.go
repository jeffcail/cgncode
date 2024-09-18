package models

type Goods struct {
	ID        int64  `xorm:"'id' pk autoincr"`
	GoodsName string `xorm:"'goods_name'"`
}

func (m *Goods) TableName() string {
	return "a_goods"
}
