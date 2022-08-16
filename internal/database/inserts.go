package database

import (
	"context"
	"test1/models"
)

func (database *Database) InsertObject(ctx context.Context, object *models.Object) (err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	err = c.QueryRow(
		ctx, "insert into test (data) values ($1) returning id",
		object.Data,
	).Scan(
		&object.ID,
	)

	return
}