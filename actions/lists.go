package actions

import (
	"errors"
	"net/http"

	"github.com/geoffjay/hatedit/models"

	"github.com/gobuffalo/buffalo"
)

type ListsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v ListsResource) List(c buffalo.Context) error {
	lists := []models.List{}
	if err := models.DB.All(&lists); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string][]models.List{"lists": lists}))
}

// Show default implementation.
func (v ListsResource) Show(c buffalo.Context) error {
	id := c.Param("list_id")
	list := models.List{}
	if err := models.DB.Find(&list, id); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]models.List{"list": list}))
}

// Create default implementation.
func (v ListsResource) Create(c buffalo.Context) error {
	user_id := c.Params().Get("user_id")
	user := &models.User{}
	if err := models.DB.Find(user, user_id); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	list := &models.List{User: user}
	if err := models.DB.Create(list); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]*models.List{"list": list}))
}

// Update default implementation.
func (v ListsResource) Update(c buffalo.Context) error {
	return c.Render(http.StatusNotFound, r.JSON(actionError(errors.New("resource not implemented"))))
}

// Destroy default implementation.
func (v ListsResource) Destroy(c buffalo.Context) error {
	return c.Render(http.StatusNotFound, r.JSON(actionError(errors.New("resource not implemented"))))
}
