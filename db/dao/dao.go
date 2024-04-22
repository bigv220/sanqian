package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "sanqiangua"

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}
