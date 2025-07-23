package app

import (
	"effective-mobile/entities"
	"errors"
	"log/slog"
)

func (a *app) ReadRecordByID(id int) (*entities.Record, error) {
	if id <= 0 {
		return nil, entities.ErrNegativeRecordID
	}

	record, err := a.db.ReadRecordByID(id)
	if errors.Is(err, entities.ErrNotFound) {
		slog.Error("record not found by id", "id", id)
		return nil, entities.ErrNotFound
	}
	if err != nil {
		slog.Error("failed to read record by ID", "id", id, "error", err.Error())
		return nil, entities.ErrInternal
	}

	return record, nil
}
