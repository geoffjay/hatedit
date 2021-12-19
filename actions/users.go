package actions

import (
	"fmt"
	"net/http"

	"github.com/geoffjay/hatedit/models"
	"github.com/sirupsen/logrus"

	"github.com/gobuffalo/buffalo"
)

func userError(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

// UsersShow default implementation.
func UsersShow(c buffalo.Context) error {
	users := []models.User{}
	name := c.Params().Get("name")
	logrus.Infof("checking for %s", name)
	stmt := fmt.Sprintf("name = '%s'", name)
	query := models.DB.Where(stmt)
	if err := query.All(&users); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(userError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]models.User{"user": users[0]}))
}

// UsersIndex default implementation.
func UsersIndex(c buffalo.Context) error {
	users := []models.User{}
	if err := models.DB.All(&users); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(userError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string][]models.User{"users": users}))
}

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	name := c.Params().Get("name")
	user := &models.User{Name: name}
	if err := models.DB.Create(user); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(userError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"user": user.String()}))
}
