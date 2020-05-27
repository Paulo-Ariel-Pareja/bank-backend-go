package users

import (
	"time"

	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username, pass string) map[string]interface{} {
	// Connect db
	db := helpers.ConnectDB()
	defer db.Close()

	user := &interfaces.User{}
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	// Verifica pass
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "User not found"}
	}

	// Busca cuenta para el usuario
	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	// Parte de la respuesta
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// Token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	// Response para el request
	var response = map[string]interface{}{"message": "Login"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
