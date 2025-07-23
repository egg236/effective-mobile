package db

import (
	"effective-mobile/entities"
)

func (d *database) DeleteRecord(id int) error {
	query := `
	DELETE FROM records
	WHERE id = $1
	`

	res, err := d.Exec(query, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return entities.ErrNotFound
	}

	return nil
}
