package renderer

import (
	"test1/internal/database"
)

type (
	Renderer struct {
		Database *database.Database
	}
)

func NewRenderer(d *database.Database) (rr *Renderer) {
	rr = &Renderer{
		Database: d,
	}

	return
}
