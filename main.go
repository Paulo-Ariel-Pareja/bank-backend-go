package main

import (
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/api"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/migrations"
)

func main() {
	migrations.Migrate()
	api.StartApi()
}
