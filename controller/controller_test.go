package controller

import (
	"github.com/runicelf/rpc-server/models"
	"testing"
)

type SuccessTestRepository struct{}

func (SuccessTestRepository) Add(login string) (string, error) {
	return "uuid", nil
}
func (SuccessTestRepository) Delete(uuid string) error {
	return nil
}
func (SuccessTestRepository) Update(user models.User) error {
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

	args := models.User{UUID: "uuid", Login: "login"}
	var result string
	err := controller.Update(nil, &args, &result)
	if err != nil {
		t.Fatal("Expected: nil, Received: ", err)
	}
	if result != "ok" {
		t.Fatal("Expected: \"ok\", Received: ", result)
	}
}

func TestController_Delete(t *testing.T) {
	controller := Controller{Repository: SuccessTestRepository{}}

	args := "uuid"
	var result string
	err := controller.Delete(nil, &args, &result)
	if err != nil {
		t.Fatal("Expected: nil, Received: ", err)
	}
	if result != "ok" {
		t.Fatal("Expected: \"ok\", Received: ", result)
	}
}
