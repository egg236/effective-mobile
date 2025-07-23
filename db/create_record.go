package db

import (
	"effective-mobile/entities"
	"errors"
)

func (d *database) CreateRecord(rec *entities.Record) (int, error) {
	query := `
	INSERT INTO records
	(service_name, price, user_id, start_date, end_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	rows, err := d.Query(query, rec.ServiceName, rec.Price, rec.UserID, rec.StartDate, rec.EndDate)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, errors.New("error while preparing rows for scan")
	}

	var id int
	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
