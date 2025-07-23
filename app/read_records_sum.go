package app

import (
	"effective-mobile/entities"
	"log/slog"
	"time"
)

func (a *app) ReadRecordsSum(userID, serviceName, dateStart, dateEnd string) (int, error) {
	records, err := a.db.ReadRecords(userID, serviceName)
	if err != nil {
		slog.Error("failed to read all records", "error", err.Error())
		return 0, entities.ErrInternal
	}

	res := 0

	var dsTime, deTime time.Time

	if dateStart != "" {
		dsTime, err = time.Parse(entities.DateFormat, dateStart)
		if err != nil {
			slog.Error("failed to parse date start", "error", err.Error())
			return 0, entities.ErrInternal
		}
	}
	if dateEnd != "" {
		deTime, err = time.Parse(entities.DateFormat, dateEnd)
		if err != nil {
			slog.Error("failed to parse date end", "error", err.Error())
			return 0, entities.ErrInternal
		}
	}
	for _, record := range records {
		if dateStart != "" {
			rStart, err := time.Parse(entities.DateFormat, record.StartDate)
			if err != nil {
				slog.Error("failed to parse date start", "error", err.Error())
				return 0, entities.ErrInternal
			}
			if rStart.Before(dsTime) {
				continue
			}

		}

		if dateEnd != "" {
			if record.EndDate != "" {
				rEnd, err := time.Parse(entities.DateFormat, record.EndDate)
				if err != nil {
					slog.Error("failed to parse date end", "error", err.Error())
					return 0, entities.ErrInternal
				}
				if rEnd.After(deTime) {
					continue
				}
			}
		}

		res += record.Price
	}

	return res, nil
}
