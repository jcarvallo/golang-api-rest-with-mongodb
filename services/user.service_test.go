package services_test

import (
	"testing"
	"time"

	m "github.com/jcarvallo/golang-api-rest-with-mongodb/models"
	"github.com/jcarvallo/golang-api-rest-with-mongodb/services"
)

func TestCreateUser(t *testing.T) {

	user := m.User{
		Name:       "Leanne Graham",
		Email:      "Sincere@april.biz",
		Phone:      "1-770-736-8031 x56442",
		Website:    "hildegard.org",
		Username:   "Bret",
		Created_At: time.Now(),
	}
	_, err := services.CreateUser(user)
	if err != nil {
		t.Error("La prueba de creacion de usuario fallo")
		t.Log(err)
		t.Fail()
	} else {
		t.Log("La prueba de creacion del usuario finalizo con exito!")
	}
}

func TestGetUsers(t *testing.T) {
	_, err := services.GetUsers()
	if err != nil {
		t.Error("La prueba consulta de usuarios fallo")
		t.Log(err)
		t.Fail()
	} else {
		t.Log("La prueba consulta de usuarios finalizo con exito!")
	}

}

func TestGetUser(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
