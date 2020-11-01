package models

import (
	orm "financial/database"
)

//Sector 定义板块结构
type Sector struct {
	ID      int32 `gorm:"pk"`
	Name    string
	Intro   string
	Created int32
}

var sectors []Sector

//SectorList 查询板块列表
func (s Sector) SectorList() (sectors []Sector, err error) {
	if err = orm.Eloquent.Find(&sectors).Error; err != nil {
		return
	}
	return
}
