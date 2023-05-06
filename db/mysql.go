package db

import (
	"fmt"
	"log"
	"to-do/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConn struct {
	mysqlAddres string
	db          *gorm.DB
}

func NewMysql(mysqladdr string) *MysqlConn {
	return &MysqlConn{
		mysqlAddres: mysqladdr,
		// db: ,
	}
}

func (m *MysqlConn) InitDB() {
	mysqlAddr := m.mysqlAddres
	db, err := gorm.Open(mysql.Open(mysqlAddr), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("InitDB Address: ", mysqlAddr)
	fmt.Println("Connected to mysql succesfully")

	m.db = db
	m.Migrate()
	fmt.Println("Table migrated succesfully")
}

func (m *MysqlConn) Migrate() {
	// var table interface{}
	// if err := m.db.Raw("show databases").Scan(&table).Error; err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(table)
	// if err := m.db.Debug().AutoMigrate(&models.Test{}); err != nil {
	if err := m.db.Debug().AutoMigrate(&models.Todo{}, &models.Activities{}); err != nil {
		log.Println("Error when migrating table")
		log.Println(err)
		panic(err)
	}
}

func (m *MysqlConn) GetDBInstance() *gorm.DB {
	return m.db
}
