package main

import (
	"fmt"

	mysqldb "absolutetech/farm_report/persistance/mysql"

	"log"
)

func main() {
	fmt.Println("Farm Report Generation")
	flg := getFlags()
	dbClient, closeDBClient, err := mysqldb.NewIgrowDBClient(flg.environment, "farm_reprt")
	if err != nil {
		print(closeDBClient)
	}

	var value int
	err = dbClient.QueryRow("select 1 from DUAL").Scan(&value)
	if err != nil {
		log.Println("query failed:", err)
		return
	}

	log.Println("value:", value)
}
