package model

// CounterModel 计数器模型
type CounterModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	ShangGua  string    `gorm:"column:shanggua" json:"shanggua"`
	XiaGua    string    `gorm:"column:xiagua" json:"xiagua"`
}
