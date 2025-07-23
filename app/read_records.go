package app

import (
	"effective-mobile/entities"
	"log/slog"
)

func (a *app) ReadRecords(userID, serviceName string) ([]*entities.Record, error) {
	records, err := a.db.ReadRecords(userID, serviceName)
	if err != nil {
		slog.Error("failed to read all records", "error", err.Error())
		return nil, entities.ErrInternal
	}

	return records, nil
}
