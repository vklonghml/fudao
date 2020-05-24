package store

import (
	"log"

	"github.com/astaxie/beego/orm"

	"fudao/cmd/options"
	"fudao/pkg/common/db"
)

func RegisterDB(dbSource, dbName string) {
	dbTimeout := options.GetOptions().DBTimeout
	dataSource := fmt.Sprintf("%s%s?charset=utf8&%s", dbSource, dbName, dbTimeout)
	orm.RegisterDriver(db.DBDriverName, orm.DBMySQL)
	orm.RegisterDataBase("default", db.DBDriverName, dataSource, 30)
}

func InitDB() error {
	// create table
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Println("failed to create tables with err:", err)
		return err
	}
}
