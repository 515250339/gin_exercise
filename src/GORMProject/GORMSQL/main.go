package main

import (
	"GORMProject/GORMSQL/config"
)

func main() {
	DB := config.Core()
	//config.CreateUser1(DB)
	//config.CreateUser2(DB)
	//config.CreateUser3(DB)
	//config.CreateUser4(DB)
	//config.CreateUser5(DB)
	config.CreateUser6(DB)
}
