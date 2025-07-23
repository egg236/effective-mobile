package app

import (
	"effective-mobile/entities"
	"log/slog"
)

func (a *app) CreateRecord(r *entities.Record) (int, error) {
	id, err := a.db.CreateRecord(r)
	if err != nil {
		slog.Error("failed to create record", "error", err.Error())
		return 0, entities.ErrInternal
	}

	return id, nil
}
