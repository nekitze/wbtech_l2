package server

import (
	"dev11/internal/config"
	"dev11/internal/middleware"
	"dev11/internal/service"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Service *service.EventController
}

func NewServer() *Server {
	return &Server{Service: service.NewEventService()}
}

func (s *Server) setupPostHandlers() {
	http.HandleFunc("/create_event", middleware.Logger(s.Service.CreateEvent))
	http.HandleFunc("/update_event", middleware.Logger(s.Service.UpdateEvent))
	http.HandleFunc("/delete_event", middleware.Logger(s.Service.DeleteEvent))
}

func (s *Server) setupGetHandlers() {
	http.HandleFunc("/events_for_day", middleware.Logger(s.Service.GetEventsForDay))
	http.HandleFunc("/events_for_week", middleware.Logger(s.Service.GetEventsForWeek))
	http.HandleFunc("/events_for_month", middleware.Logger(s.Service.GetEventsForMonth))
}

func (s *Server) SetupHandlers() {
	s.setupGetHandlers()
	s.setupPostHandlers()
}

func (s *Server) Up() {
	cfg := config.NewServerConfig()
	addr := ":" + cfg.Port
	fmt.Println("Server listening on", addr)
	log.Println(http.ListenAndServe(addr, nil))
}
