package main

import (
	"dipay-test/config"
	"dipay-test/db"
	"dipay-test/server"
)

func main() {
	config.InitEnv()
	mongo := db.NewMongodb(config.GetMongoEnv())
	mongo.InitDB()
	server.Init(mongo)
}
