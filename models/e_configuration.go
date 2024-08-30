package models

type User struct {
	ID   int64 `xorm:"'id' pk autoincr"`
	Name string
}

func (m *User) TableName() string {
	return "a_user"
}
