package actions

import (
	"net/http"

	"github.com/geoffjay/hatedit/models"

	"github.com/gobuffalo/buffalo"
)

// UsersShow default implementation.
func UsersShow(c buffalo.Context) error {
	id := c.Param("user_id")
	user := models.User{}
	if err := models.DB.Eager().Find(&user, id); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]models.User{"user": user}))
}

// UsersIndex default implementation.
func UsersIndex(c buffalo.Context) error {
	users := []models.User{}
	if err := models.DB.Eager().All(&users); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string][]models.User{"users": users}))
}

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	name := c.Params().Get("name")
	user := &models.User{Name: name}
	if err := models.DB.Create(user); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]*models.User{"user": user}))
}
