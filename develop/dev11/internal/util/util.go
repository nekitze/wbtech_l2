package util

import (
	"dev11/internal/dto"
	"encoding/json"
	"errors"
	"time"
)

func EventToJson(event dto.Event) (data []byte, err error) {
	data, err = json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return
}

func EventsToJson(events ...dto.Event) (data []byte, err error) {
	data, err = json.Marshal(events)
	if err != nil {
		return nil, err
	}
	return
}

func JsonToEvent(data []byte) (event dto.Event, err error) {
	err = json.Unmarshal(data, &event)
	if err != nil {
		return dto.Event{}, err
	}
	if len(event.Name) == 0 {
		return dto.Event{}, errors.New("empty event name")
	}
	if !IsValidDate(event.Date) {
		return dto.Event{}, errors.New("wrong date")
	}
	return
}

func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
