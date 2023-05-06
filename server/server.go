package server

import (
	"to-do/db"
)

func Init(mysql *db.MysqlConn) {
	r := NewRouter(mysql)
	r.Run(":3030")
}
