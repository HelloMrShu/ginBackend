package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

//DB object
var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/finance?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
