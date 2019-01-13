package controller

import (
	"github.com/runicelf/rpc-server/models"
	"reflect"
	"testing"
	"time"
)

type SuccessTestRepository struct{}

func (SuccessTestRepository) Add(login string) (string, error) {
	return "uuid", nil
}

func (SuccessTestRepository) Get(uuid string) (models.DBModelUser, error) {
	return models.DBModelUser{UUID: "uuid", Login: "login", Date: time.Time{}}, nil
}
func (SuccessTestRepository) Update(user models.RequestModelUser) error {
	return nil
}

func TestController_Add(t *testing.T) {
	controller := Controller{Repository: SuccessTestRepository{}}

	args := "uuid"
	var result string
	err := controller.Add(nil, &args, &result)
	if err != nil {
		t.Fatal("Expected: nil, Received: ", err)
	}
	if result != args {
		t.Fatal("Expected: ", args, ", Received: ", result)
	}
}

func TestController_Update(t *testing.T) {
	controller := Controller{Repository: SuccessTestRepository{}}

	args := models.RequestModelUser{UUID: "uuid", Login: "login"}
	var result string
	err := controller.Update(nil, &args, &result)
	if err != nil {
		t.Fatal("Expected: nil, Received: ", err)
	}
	if result != "ok" {
		t.Fatal("Expected: \"ok\", Received: ", result)
	}
}

func TestController_Get(t *testing.T) {
	controller := Controller{Repository: SuccessTestRepository{}}

	args := "uuid"
	expectedOutput := models.DBModelUser{UUID: "uuid", Login: "login", Date: time.Time{}}
	var result models.DBModelUser
	err := controller.Get(nil, &args, &result)
	if err != nil {
		t.Fatal("Expected: nil, Received: ", err)
	}
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Fatal("Expected: expectedOutput, Received: ", result)
	}
}
