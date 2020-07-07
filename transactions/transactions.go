package transactions

import (
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/database"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/interfaces"
)

func CreateTransaction(From, To uint, Amount int) {
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	database.DB.Create(&transaction)
}

func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction {
	transactions := []interfaces.ResponseTransaction{}

	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(interfaces.Transaction{From: id}).Or(interfaces.Transaction{To: id}).Scan(&transactions)

	return transactions
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	// Validate JWT
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		// Find and return transactions
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ? ", id).Scan(&accounts)

		transactions := []interfaces.ResponseTransaction{}
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "Operation complete"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
