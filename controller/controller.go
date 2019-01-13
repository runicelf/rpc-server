package controller

import (
	"github.com/runicelf/rpc-server/models"
	"net/http"
)

type Repository interface {
	Add(login string) (string, error)
	Get(uuid string) (models.DBModelUser, error)
	Update(user models.RequestModelUser) error
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

func (c *Controller) Get(r *http.Request, args *string, result *models.DBModelUser) error {
	user, err := c.Repository.Get(*args)
	if err != nil {
		return err
	}

	*result = user
	return nil
}

func (c *Controller) Update(r *http.Request, args *models.RequestModelUser, result *string) error {
	err := c.Repository.Update(*args)
	if err != nil {
		return err
	}

	*result = "ok"
	return nil
}
