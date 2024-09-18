package models

type User struct {
	ID   int64  `xorm:"'id' pk autoincr"`
	Name string `xorm:"'name'"`
}

func (m *User) TableName() string {
	return "a_user"
}
