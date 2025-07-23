package db

import (
	"effective-mobile/entities"
)

func (d *database) UpdateRecord(rec *entities.Record) error {
	query := `
	UPDATE records
	SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
	WHERE id = $6
	`

	res, err := d.Exec(query, rec.ServiceName, rec.Price, rec.UserID, rec.StartDate, rec.EndDate, rec.ID)
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
