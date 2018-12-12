package grifts

import (
	"github.com/gobuffalo/buffalo"
	"scrum_vote_api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
