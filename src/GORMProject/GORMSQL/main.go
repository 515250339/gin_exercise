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
	//config.CreateUser6(DB)
	//config.CreateUser7(DB)
	//config.CreateUser8(DB)
	//config.CreateUser9(DB)
	//config.UpsertFunc(DB)

	//config.QueryTest1(DB)
	//config.QueryTest2(DB)
	//config.QueryTest3(DB)
	//config.QueryALlTest(DB)
	//config.FilterTest(DB)
	//config.FilterMapAndStruct(DB)
	//config.FilterTest2(DB)
	//config.FilterNot(DB)
	//config.FilterOr(DB)
	//config.SelectField(DB)
	config.OrderData(DB)
}
