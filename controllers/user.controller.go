package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcarvallo/golang-api-rest-with-mongodb/models"
	"github.com/jcarvallo/golang-api-rest-with-mongodb/services"
	"github.com/jcarvallo/golang-api-rest-with-mongodb/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, error := services.GetUsers()
	if error != nil {
		utils.RespondWithError(error, w)
	}
	json.NewEncoder(w).Encode(&users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	user, error := services.GetUser(params["id"])

	if error != nil {
		utils.RespondWithError(error, w)
	}

	json.NewEncoder(w).Encode(&user)

}
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result, error := services.CreateUser(user)
	if error != nil {
		utils.RespondWithError(error, w)
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&user)
	result, error := services.UpdateUser(user, params["id"])
	if error != nil {
		utils.RespondWithError(error, w)
	}
	json.NewEncoder(w).Encode(result)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	error := services.DeleteUser(params["id"])

	if error != nil {
		utils.RespondWithError(error, w)
	}
	utils.RespondNotContent("User delete", w)
}
