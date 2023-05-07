package config

import (
	"flag"
	"fmt"
)

// var mysql_host *string
var mongouri string

func InitEnv() {
	fmt.Println("-----------------Init Env-----------------")
	flagSet()
	fmt.Println("-----------------Init Env-----------------")
}

func flagSet() {

	flag.StringVar(&mongouri, "MONGO_URI", "", "You Need To Specify MONGO_URI")
	flag.Parse()

	// fmt.Println(flag.Lookup("MONGO_URI").Value.String())
	if flag.Lookup("MONGO_URI").Value.String() == "" {
		panic("You Need To Specify MONGO_URI")
	}

}

func GetMongoEnv() string {
	return mongouri
}
