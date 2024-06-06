package service

import (
	"dev11/internal/dto"
	"dev11/internal/util"
	"errors"
	"io"
	"net/http"
)

func (s *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := validateAndParseRequest(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	s.Mu.RLock()
	_, exists := s.Storage[event.Name]
	s.Mu.RUnlock()
	if exists {
		w.WriteHeader(409)
		return
	}

	s.Mu.Lock()
	s.Storage[event.Name] = event
	s.Mu.Unlock()
	w.WriteHeader(201)
	resp, _ := util.EventToJson(event)
	_, _ = w.Write(resp)
}

func (s *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := validateAndParseRequest(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	s.Mu.RLock()
	_, exists := s.Storage[event.Name]
	s.Mu.RUnlock()
	if !exists {
		w.WriteHeader(404)
		return
	}

	s.Mu.Lock()
	s.Storage[event.Name] = event
	s.Mu.Unlock()
	w.WriteHeader(204)
	resp, _ := util.EventToJson(event)
	_, _ = w.Write(resp)
}

func (s *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventName := r.FormValue("name")

	if len(eventName) == 0 {
		w.WriteHeader(400)
		return
	}

	s.Mu.RLock()
	_, exists := s.Storage[eventName]
	s.Mu.RUnlock()
	if !exists {
		w.WriteHeader(404)
		return
	}

	s.Mu.Lock()
	delete(s.Storage, eventName)
	s.Mu.Unlock()
	w.WriteHeader(200)
}

func validateAndParseRequest(r *http.Request) (event dto.Event, err error) {
	if r.Method != http.MethodPost {
		return dto.Event{}, errors.New("unsupported method")
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return dto.Event{}, err
	}
	defer func() { _ = r.Body.Close() }()

	event, err = util.JsonToEvent(data)
	if err != nil {
		return dto.Event{}, err
	}

	return event, nil
}
