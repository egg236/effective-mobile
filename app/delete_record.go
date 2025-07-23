package app

import (
	"effective-mobile/entities"
	"errors"
	"log/slog"
)

func (a *app) DeleteRecord(id int) error {
	if id <= 0 {
		return entities.ErrNegativeRecordID
	}

	err := a.db.DeleteRecord(id)
	if errors.Is(err, entities.ErrNotFound) {
		slog.Error("record not found by id", "id", id)
		return entities.ErrNotFound
	}
	if err != nil {
		slog.Error("failed to delete record", "id", id, "error", err.Error())
		return entities.ErrInternal
	}

	return nil
}
