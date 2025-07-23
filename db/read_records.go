package db

import (
	"effective-mobile/entities"
	"fmt"
)

func (d *database) ReadRecords(userID, serviceName string) ([]*entities.Record, error) {
	query := `
	SELECT * from records`

	// применяем фильтры
	if userID != "" && serviceName != "" {
		query = query + fmt.Sprintf(" WHERE user_id = '%s' AND service_name = '%s'", userID, serviceName)
	} else {
		if userID != "" {
			query = query + fmt.Sprintf(" WHERE user_id = '%s'", userID)
		}
		if serviceName != "" {
			query = query + fmt.Sprintf(" WHERE service_name = '%s'", serviceName)
		}
	}

	query += " ORDER BY ID ASC"

	rows, err := d.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rec := make([]*entities.Record, 0)
	for rows.Next() {
		var row entities.Record
		err = rows.Scan(&row.ID, &row.ServiceName, &row.Price, &row.UserID, &row.StartDate, &row.EndDate)
		if err != nil {
			return nil, err
		}

		rec = append(rec, &row)
	}

	return rec, nil
}
