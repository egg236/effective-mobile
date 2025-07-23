package db

import (
	"effective-mobile/entities"
)

func (d *database) ReadRecordByID(id int) (*entities.Record, error) {
	query := `
	SELECT * from records
	WHERE id = $1`

	rows, err := d.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, entities.ErrNotFound
	}

	var rec entities.Record
	err = rows.Scan(&rec.ID, &rec.ServiceName, &rec.Price, &rec.UserID, &rec.StartDate, &rec.EndDate)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}
