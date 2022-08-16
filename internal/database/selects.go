package database

import (
	"context"

	"test1/models"
)

func (database *Database) SelectDatas(ctx context.Context) (objects models.Objects, err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	rows, err := c.Query(ctx, "select id, data from test")
	if err != nil {
		return
	}

	object := models.Object{}

	for rows.Next() {
		err = rows.Scan(&object.ID, &object.Data)
		if err != nil {
			return
		}

		objects = append(objects, object)
	}

	return
}