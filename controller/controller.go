package controller

import (
	"github.com/runicelf/rpc-server/models"
	"net/http"
)

type Repository interface {
	Add(login string) (string, error)
	Delete(uuid string) error
	Update(user models.User) error
}

type Controller struct {
	Repository Repository
}

func New(r Repository) *Controller {
	return &Controller{Repository: r}
}

func (c *Controller) Add(r *http.Request, args *string, result *string) error {
	uuid, err := c.Repository.Add(*args)
	if err != nil {
		return err
	}

	*result = uuid
	return nil
}

func (c *Controller) Delete(r *http.Request, args *string, result *string) error {
	err := c.Repository.Delete(*args)
	if err != nil {
		return err
	}

	*result = "ok"
	return nil
}

func (c *Controller) Update(r *http.Request, args *models.User, result *string) error {
	err := c.Repository.Update(*args)
	if err != nil {
		return err
	}

	*result = "ok"
	return nil
}
