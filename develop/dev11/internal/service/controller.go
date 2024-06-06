package service

import (
	"dev11/internal/dto"
	"sync"
)

type EventController struct {
	Storage map[string]dto.Event
	Mu      sync.RWMutex
}

func NewEventService() *EventController {
	return &EventController{Storage: make(map[string]dto.Event)}
}
