package models

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite/v4"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	model, err := suite.NewModelWithFixtures(packr.New("app:models:test:fixtures", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ModelSuite{
		Model: model,
	}
	suite.Run(t, as)
}
