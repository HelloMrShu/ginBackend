package models

import (
	. "financial/database"
)

//Sector 定义板块结构
type Sector struct {
	ID      int `gorm:"pk"`
	Name    string
	Intro   string
	Created int64
}

var sectors []Sector

//SectorList 查询板块列表
func (s Sector) SectorList(page, pageSize int) (sectors []Sector, err error) {
	offset := (page - 1) * pageSize
	if err = DB.Offset(offset).Limit(pageSize).Find(&sectors).Error; err != nil {
		return
	}
	return
}

//SectorSave 保存板块
func (s *Sector) SectorSave() error {

	sector := &Sector{
		ID: s.ID,
	}

	DB.First(&sector)

	var err error
	if sector != nil {
		err = DB.Model(&sector).Updates(Sector{Name: s.Name, Intro: s.Intro}).Error
	} else {
		err = DB.Create(&sector).Error

	}
	return err
}

//SectorDelete 删除板块
func (s *Sector) SectorDelete() error {
	return DB.Delete(&Sector{ID: s.ID}).Error
}

//SectorCount 统计表中数据量
func (s Sector) SectorCount() int {
	DB.Select("id").Find(&sectors)
	return len(sectors)
}
