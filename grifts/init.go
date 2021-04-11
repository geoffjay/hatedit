package grifts

import (
	"github.com/geoffjay/hatedit/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
