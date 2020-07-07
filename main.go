package main

import (
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/api"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/database"
)

func main() {
	//migrations.Migrate()
	//migrations.MigrateTransactions()
	database.InitDatabase()
	api.StartApi()
}
