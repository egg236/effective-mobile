package api

import (
	"effective-mobile/entities"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleCreateRecord(rw http.ResponseWriter, r *http.Request) {
	record, err := getRecordFromRequest(r)
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusBadRequest, false)
		return
	}

	id, err := s.app.CreateRecord(record)
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, id, http.StatusCreated, true)
}

func (s *Server) handleReadRecords(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	serviceName := r.URL.Query().Get("serviceName")
	records, err := s.app.ReadRecords(userID, serviceName)
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, records, http.StatusOK, true)
}

func (s *Server) handleReadRecordsSum(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	serviceName := r.URL.Query().Get("serviceName")

	dateStart := r.URL.Query().Get("dateStart")
	dateEnd := r.URL.Query().Get("dateEnd")

	var sTime, eTime time.Time
	var err error
	if dateStart != "" {
		sTime, err = time.Parse(entities.DateFormat, dateStart)
		if err != nil {
			writeResponse(rw, entities.ErrWrongDateFormat.Error(), http.StatusBadRequest, false)
			return
		}
	}

	if dateEnd != "" {
		eTime, err = time.Parse(entities.DateFormat, dateEnd)
		if err != nil {
			writeResponse(rw, entities.ErrWrongDateFormat.Error(), http.StatusBadRequest, false)
			return
		}
	}

	if dateStart != "" && dateEnd != "" {
		if eTime.Before(sTime) {
			writeResponse(rw, entities.ErrDateWrongRange.Error(), http.StatusBadRequest, false)
			return
		}
	}

	records, err := s.app.ReadRecordsSum(userID, serviceName, dateStart, dateEnd)
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, records, http.StatusOK, true)
}

func (s *Server) handleReadRecordByID(rw http.ResponseWriter, r *http.Request) {
	recordID, ok := mux.Vars(r)["recordID"]
	if !ok {
		writeResponse(rw, entities.ErrEmptyRecordID.Error(), http.StatusBadRequest, false)
		return
	}
	rInt, err := strconv.Atoi(recordID)
	if err != nil || rInt <= 0 {
		writeResponse(rw, entities.ErrWrongRecordID.Error(), http.StatusBadRequest, false)
		return
	}

	record, err := s.app.ReadRecordByID(rInt)
	if errors.Is(err, entities.ErrNotFound) {
		writeResponse(rw, err.Error(), http.StatusNotFound, false)
		return
	}
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, record, http.StatusOK, true)
}

func (s *Server) handleUpdateRecord(rw http.ResponseWriter, r *http.Request) {
	record, err := getRecordFromRequest(r)
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusBadRequest, false)
		return
	}

	err = s.app.UpdateRecord(record)
	if errors.Is(err, entities.ErrNotFound) {
		writeResponse(rw, err.Error(), http.StatusNotFound, false)
		return
	}
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, nil, http.StatusNoContent, true)
}

func (s *Server) handleDeleteRecord(rw http.ResponseWriter, r *http.Request) {
	recordID, ok := mux.Vars(r)["recordID"]
	if !ok {
		writeResponse(rw, entities.ErrEmptyRecordID.Error(), http.StatusBadRequest, false)
		return
	}
	rInt, err := strconv.Atoi(recordID)
	if err != nil || rInt <= 0 {
		writeResponse(rw, entities.ErrWrongRecordID.Error(), http.StatusBadRequest, false)
		return
	}

	err = s.app.DeleteRecord(rInt)
	if errors.Is(err, entities.ErrNotFound) {
		writeResponse(rw, err.Error(), http.StatusNotFound, false)
		return
	}
	if err != nil {
		writeResponse(rw, err.Error(), http.StatusInternalServerError, false)
		return
	}

	writeResponse(rw, nil, http.StatusNoContent, true)
}
