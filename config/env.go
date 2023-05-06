package config

import (
	"flag"
	"fmt"
)

// var mysql_host *string
var mysql_host string
var mysql_port string
var mysql_user string
var mysql_pw string
var mysql_dbname string

func InitEnv() {
	fmt.Println("-----------------Init Env-----------------")
	flagSet()
	fmt.Println("-----------------Init Env-----------------")
}

func flagSet() {

	flag.StringVar(&mysql_host, "MYSQL_HOST", "nil mysql host", "You Need To Specify MYSQL_HOST")
	flag.StringVar(&mysql_user, "MYSQL_USER", "nil mysql user", "You Need To Specify MYSQL_USER")
	flag.StringVar(&mysql_port, "MYSQL_PORT", "3306", "You Need To Specify MYSQL_PORT")
	flag.StringVar(&mysql_pw, "MYSQL_PASSWORD", "nil mysql password", "You Need To Specify MYSQL_PASSWORD")
	flag.StringVar(&mysql_dbname, "MYSQL_DBNAME", "nil dbname", "You Need To Specify MYSQL_DBNAME")
	flag.Parse()

	if flag.Lookup("MYSQL_HOST") == nil {
		panic("You Need To Specify MYSQL_HOST")
	}
	if flag.Lookup("MYSQL_USER") == nil {
		panic("You Need To Specify MYSQL_USER")
	}
	if flag.Lookup("MYSQL_PASSWORD") == nil {
		panic("You Need To Specify MYSQL_PASSWORD")
	}
	if flag.Lookup("MYSQL_DBNAME") == nil {
		panic("You Need To Specify MYSQL_DBNAME")
	}

	fmt.Println(mysql_host)
	fmt.Println(mysql_port)
	fmt.Println(mysql_user)
	fmt.Println(mysql_pw)
	fmt.Println(mysql_dbname)
	// panic(1)
}

func GetMysqlEnv() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql_user, mysql_pw, mysql_host, mysql_port, mysql_dbname)
}
