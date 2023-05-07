package server

import (
	"dipay-test/db"
)

func Init(mongo *db.MongodbCon) {
	r := NewRouter(mongo)
	r.Run(":8080")
}
