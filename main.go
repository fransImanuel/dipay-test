package main

import (
	"to-do/config"
	"to-do/db"
	"to-do/server"
)

func main() {
	config.InitEnv()
	mysql := db.NewMysql(config.GetMysqlEnv())
	mysql.InitDB()
	server.Init(mysql)
}
