package entities

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

const (
	DateFormat = "01-2006"
)

// main entity
type Record struct {
	ID          int64  `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date,omitempty"`
}

func (r *Record) Scan(src any) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("value is not a []byte")
	}

	return json.Unmarshal(b, r)
}

func (r *Record) Validate() error {
	if r.UserID == "" {
		return ErrEmptyUserID
	}
	if _, err := uuid.Parse(r.UserID); err != nil {
		return ErrWrongUserID
	}

	if r.Price < 0 {
		return ErrNegativePrice
	}

	var sTime, eTime time.Time
	var err error
	if sTime, err = time.Parse(DateFormat, r.StartDate); err != nil {
		return ErrWrongDateFormat
	}

	if r.EndDate != "" {
		if eTime, err = time.Parse(DateFormat, r.EndDate); err != nil {
			return ErrWrongDateFormat
		}
		if eTime.Before(sTime) {
			return ErrDateWrongRange
		}
	}

	if r.ServiceName == "" {
		return ErrEmptyServiceName
	}

	return nil
}

// server response format
type Response struct {
	IsSuccess bool `json:"is_success"`
	Result    any  `json:"result"`
}

// errors
var (
	ErrNotFound = errors.New("error not found")

	ErrInternal = errors.New("error internal server")

	ErrWrongRecordID    = errors.New("wrong recordID")
	ErrEmptyRecordID    = errors.New("record id cant be zero string")
	ErrNegativeRecordID = errors.New("record id must be positive")
	ErrDateWrongRange   = errors.New("wrong date range")
	ErrWrongDateFormat  = errors.New("wrong date format")
	ErrEmptyServiceName = errors.New("serviceName cant be zero string")
	ErrNegativePrice    = errors.New("price cant be less than zero")
	ErrWrongUserID      = errors.New("userID must be UUID")
	ErrEmptyUserID      = errors.New("userID cant be zero string")
)
