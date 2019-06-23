package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//全局变量，数据库实例
var db *sql.DB


const (
	host     = "postgres_net"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

//初始化连接数据库
func Initdb() {
	var err error
	//打开数据库
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", dbInfo)

	err = db.Ping()
	checkError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)

	fmt.Println("DB connected!")
}

//关闭数据库
func Close() {
	db.Close()
}

//错误处理
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
