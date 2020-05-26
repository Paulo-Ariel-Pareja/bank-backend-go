package migrations

import (
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgress sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccount() {
	db := connectDB()
	defer db.Close()

	users := [2]User{
		{Username: "Paulo", Email: "correo1@correo.com"},
		{Username: "Ariel", Email: "correo2@correo.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + "account"), Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	db := connectDB()
	defer db.Close()

	db.AutoMigrate(&User{}, &Account{})

	createAccount()
}
