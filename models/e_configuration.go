package models

// EConfiguration 系统配置表
type EConfiguration struct {
	ID                  int64  `xorm:"'id' pk autoincr"`      // 主键id
	DailyTranNums       int    `xorm:"daily_tran_nums"`       // 每日转账笔数
	AuthEnergyThreshold string `xorm:"auth_energy_threshold"` // 授权能量阀值
	SingleEnergySupple  string `xorm:"single_energy_supple"`  // 单次补充能量
}

func (m *EConfiguration) TableName() string {
	return "e_configuration"
}
