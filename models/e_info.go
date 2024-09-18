package models

type Info struct {
	ID     int64  `xorm:"'id' pk autoincr"`
	IdCard string `xorm:"'id_card'"`
}

func (m *Info) TableName() string {
	return "e_info"
}
