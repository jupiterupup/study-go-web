package store

import (
	"fire-press/api/types"
	"fire-press/apiservice/util/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

const DsnKey = "dsn"

var db *gorm.DB

func init() {
	// dsn更多配置参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := nacos.GetConfIfPresent(DsnKey)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 使用单数表名：避免gorm自动转义的时候加上s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("connect to db failed, error=", err)
	}
	db = database
}

func SaveConfig(config types.Config) {
	db.Create(config)
}

func DeleteConfigById(id int64) int64 {
	return db.Delete(&types.Config{Id: id}).RowsAffected
}
