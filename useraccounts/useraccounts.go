package useraccounts

import (
	"fmt"

	"github.com/Paulo-Ariel-Pareja/bank-backend-go/database"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/interfaces"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/transactions"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)

	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not fund"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You don't have permission for this operation"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is to small"}
		}
		updatedAccount := updateAccount(from, int(fromAccount.Balance)-amount)
		updateAccount(to, int(toAccount.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "Transactions completed"}
		response["data"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
