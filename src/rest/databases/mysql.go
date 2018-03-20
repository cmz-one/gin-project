package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:qwe757575@tcp(101.236.25.58:3306)/rest?parseTime=true")
	if err != nil{
		log.Fatalln(err.Error())
	}
	SqlDB.SetMaxIdleConns(20)//空闲连接
	SqlDB.SetMaxOpenConns(20)//最大打开连接
	if err := SqlDB.Ping(); err != nil{
		log.Fatalln(err.Error())
	}
}