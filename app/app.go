package app

import (
	db2 "effective-mobile/db"
	"effective-mobile/entities"
	"log/slog"
)

type app struct {
	db db2.DB
}

type App interface {
	// CreateRecord Создание записи
	CreateRecord(r *entities.Record) (int, error)
	// DeleteRecord Удаление записи
	DeleteRecord(id int) error
	// ReadRecordByID Чтение записи по ID
	ReadRecordByID(id int) (*entities.Record, error)
	// ReadRecords Чтение записей в системе по фильтрам
	ReadRecords(userID, serviceName string) ([]*entities.Record, error)
	// ReadRecordsSum Сумма цен записей по фильтрам
	ReadRecordsSum(userID, serviceName, dateStart, dateEnd string) (int, error)
	// UpdateRecord Обновление записи по ID
	UpdateRecord(record *entities.Record) error
}

func NewApp() (App, error) {
	db, err := db2.NewDB()
	if err != nil {
		slog.Error("failed to connect to database", "error", err.Error())
		return nil, err
	}

	slog.Info("successfully connected to database")

	return &app{
		db: db,
	}, nil
}
