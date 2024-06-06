package service

import (
	"dev11/internal/dto"
	"dev11/internal/util"
	"errors"
	"net/http"
	"time"
)

func (s *EventController) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userId, date, err := validateAndParseQuery(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var eventsForDay []dto.Event
	s.Mu.RLock()
	for _, v := range s.Storage {
		vDate, _ := time.Parse("2006-01-02", v.Date)
		if v.UserID == userId && vDate.Equal(date) {
			eventsForDay = append(eventsForDay, v)
		}
	}
	s.Mu.RUnlock()

	resp, err := util.EventsToJson(eventsForDay...)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func (s *EventController) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userId, date, err := validateAndParseQuery(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var eventsForWeek []dto.Event
	s.Mu.RLock()
	for _, v := range s.Storage {
		vDate, _ := time.Parse("2006-01-02", v.Date)
		if v.UserID == userId && vDate.Equal(date) || (vDate.After(date) && vDate.Before(date.Add(time.Hour*24*7))) {
			eventsForWeek = append(eventsForWeek, v)
		}
	}
	s.Mu.RUnlock()

	resp, err := util.EventsToJson(eventsForWeek...)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func (s *EventController) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	userId, date, err := validateAndParseQuery(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var eventsForWeek []dto.Event

	s.Mu.RLock()
	for _, v := range s.Storage {
		vDate, _ := time.Parse("2006-01-02", v.Date)
		if v.UserID == userId && vDate.Month() == date.Month() {
			eventsForWeek = append(eventsForWeek, v)
		}
	}
	s.Mu.RUnlock()

	resp, err := util.EventsToJson(eventsForWeek...)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func validateAndParseQuery(r *http.Request) (userId string, date time.Time, err error) {
	if r.Method != http.MethodGet {
		return "", time.Time{}, errors.New("unsupported method")
	}

	userId = r.FormValue("user_id")
	dateParam := r.FormValue("date")

	if len(userId) == 0 || len(dateParam) == 0 {
		return "", time.Time{}, errors.New("wrong query params")
	}

	date, err = time.Parse("2006-01-02", dateParam)
	if err != nil {
		return "", time.Time{}, err
	}

	return userId, date, nil
}
