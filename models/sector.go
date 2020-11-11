package models

import (
	. "financial/database"
	"time"
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
func (s Sector) SectorList() (sectors []Sector, err error) {
	if err = DB.Find(&sectors).Error; err != nil {
		return
	}
	return
}

//SectorSave 保存板块
func (s *Sector) SectorSave() error {
	sector := &Sector{
		Name:    s.Name,
		Intro:   s.Intro,
		Created: time.Now().Unix(),
	}

	err := DB.Create(&sector).Error
	return err
}

//SectorDelete 删除板块
func (s *Sector) SectorDelete() error {
	return DB.Delete(&Sector{ID: s.ID}).Error
}
