package migrations

import (
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/interfaces"
)

func createAccount() {
	db := helpers.ConnectDB()
	defer db.Close()

	users := &[2]interfaces.User{
		{Username: "Paulo", Email: "correo1@correo.com"},
		{Username: "Ariel", Email: "correo2@correo.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + "account"), Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	defer db.Close()

	db.AutoMigrate(User, Account)

	createAccount()
}
