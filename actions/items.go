package actions

import (
	"errors"
	"net/http"

	"github.com/geoffjay/hatedit/models"

	"github.com/gobuffalo/buffalo"
)

type ItemsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v ItemsResource) List(c buffalo.Context) error {
	items := []models.Item{}
	if err := models.DB.All(&items); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string][]models.Item{"items": items}))
}

// Show default implementation.
func (v ItemsResource) Show(c buffalo.Context) error {
	id := c.Param("item_id")
	item := models.Item{}
	if err := models.DB.Find(&item, id); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]models.Item{"item": item}))
}

// Create default implementation.
func (v ItemsResource) Create(c buffalo.Context) error {
	text := c.Params().Get("text")
	list_id := c.Params().Get("list_id")
	list := &models.List{}
	if err := models.DB.Find(list, list_id); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	item := &models.Item{Text: text, List: list}
	if err := models.DB.Create(item); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(actionError(err)))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]*models.Item{"item": item}))
}

// Update default implementation.
func (v ItemsResource) Update(c buffalo.Context) error {
	return c.Render(http.StatusNotFound, r.JSON(actionError(errors.New("resource not implemented"))))
}

// Destroy default implementation.
func (v ItemsResource) Destroy(c buffalo.Context) error {
	return c.Render(http.StatusNotFound, r.JSON(actionError(errors.New("resource not implemented"))))
}
