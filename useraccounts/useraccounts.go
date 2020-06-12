package useraccounts

import "github.com/Paulo-Ariel-Pareja/bank-backend-go/interfaces"

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	defer db.Close()

	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)

}
