package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/transactions"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/useraccounts"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/users"
	"github.com/gorilla/mux"
)

type Register struct {
	Username string
	Email    string
	Password string
}

type TransactionBody struct {
	UserId uint
	From   uint
	To     uint
	Amount int
}

type Login struct {
	Username string
	Password string
}

func rearbody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "Login" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		// Error
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body := rearbody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := rearbody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	apiResponse(register, w)
}

func getuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func getMyTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]
	auth := r.Header.Get("Authorization")

	transactions := transactions.GetMyTransactions(userId, auth)
	apiResponse(transactions, w)
}

func transaction(w http.ResponseWriter, r *http.Request) {
	body := rearbody(r)
	auth := r.Header.Get("Authorization")
	var formattedBody TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	operation := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, auth)
	apiResponse(operation, w)

}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandle)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/transaction/{userID}", getMyTransactions).Methods("GET")
	router.HandleFunc("/user/{id}", getuser).Methods("GET")
	fmt.Println("App iniciada en el puerto 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
