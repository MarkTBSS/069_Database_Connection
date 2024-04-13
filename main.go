package main

import (
	"fmt"

	"github.com/MarkTBSS/069_Database_Connection/config"
	"github.com/MarkTBSS/069_Database_Connection/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db)
}
