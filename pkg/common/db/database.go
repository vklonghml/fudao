package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBDriverName   = "mysql"
	createDBRawSql = "CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;"
)

// func GetDBSource(dataSource string) (string, string, error) {
// 	start := strings.Index(dataSource, "/") + 1
// 	end := strings.Index(dataSource, "?")
// 	dbName := dataSource[start:end]

// }

func CreateDBIfNeeded(dataSource, dbName string) error {
	log.Println("dataSource:", dataSource, "dbName", dbName)
	conn, err := sql.Open(DBDriverName, dataSource)
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		return err
	}
	rawSql := fmt.Sprintf(createDBRawSql, dbName)
	_, err = conn.Exec(rawSql)
	defer conn.Close()
	return err
}
