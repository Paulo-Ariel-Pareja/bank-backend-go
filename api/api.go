package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Paulo-Ariel-Pareja/bank-backend-go/helpers"
	"github.com/Paulo-Ariel-Pareja/bank-backend-go/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)
	if login["message"] == "Login" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		// Error
		resp := ErrResponse{Message: "Usuario o contraseña incorrecto"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App iniciada en el puerto 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
