package app

import (
	"effective-mobile/entities"
	"errors"
	"log/slog"
)

func (a *app) UpdateRecord(r *entities.Record) error {
	if r.ID <= 0 {
		return entities.ErrNegativeRecordID
	}

	err := a.db.UpdateRecord(r)
	if errors.Is(err, entities.ErrNotFound) {
		slog.Error("record not found by id", "id", r.ID)
		return entities.ErrNotFound
	}
	if err != nil {
		slog.Error("failed to update record", "id", r.ID, "error", err.Error())
		return entities.ErrInternal
	}

	return nil
}
