package grifts

import (
	"errors"
	"fmt"
	"github.com/geoffjay/hatedit/models"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		return nil
	})

	grift.Desc("create:user", "Creates a new user")
	grift.Add("create:user", func(c *grift.Context) error {
		if len(c.Args) < 1 {
			return errors.New("a <user> argument is required")
		}
		u := &models.User{Name: c.Args[0]}
		if err := models.DB.Create(u); err != nil {
			return err
		}
		return nil
	})

	grift.Desc("create:list", "Creates a new list for the user")
	grift.Add("create:list", func(c *grift.Context) error {
		if len(c.Args) < 1 {
			return errors.New("a <user> argument is required")
		}
		// get the user
		users := []models.User{}
		stmt := fmt.Sprintf("name = '%s'", c.Args[0])
		query := models.DB.Where(stmt)
		if err := query.All(&users); err != nil {
			return err
		}
		// create the list
		l := &models.List{
			User: &users[0],
		}
		if err := models.DB.Create(l); err != nil {
			return err
		}
		return nil
	})

	grift.Desc("create:item", "Creates a new list item for the user")
	grift.Add("create:item", func(c *grift.Context) error {
		if len(c.Args) < 2 {
			return errors.New("a <user> argument is required")
		}
		// get the user
		var users []models.User
		stmt := fmt.Sprintf("name = '%s'", c.Args[0])
		query := models.DB.Where(stmt)
		if err := query.Eager().All(&users); err != nil {
			return err
		}
		// create the item
		i := &models.Item{
			Text: c.Args[1],
			List: &users[0].List,
		}
		if err := models.DB.Create(i); err != nil {
			return err
		}
		return nil
	})
})
