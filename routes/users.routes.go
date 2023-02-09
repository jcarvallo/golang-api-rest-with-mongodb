package routes

import (
	"github.com/jcarvallo/golang-api-rest-with-mongodb/controllers"
)

func RoutesUser() {
	Routes.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	Routes.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	Routes.HandleFunc("/users", controllers.PostUser).Methods("POST")
	Routes.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	Routes.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

}
