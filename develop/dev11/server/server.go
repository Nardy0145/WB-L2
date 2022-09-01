package server

import (
	"log"
	"net/http"
	"server/middleware"
	"server/server/handlers"
)

// InitHandlers ../
func InitHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", handlers.CreateEvent)
	mux.HandleFunc("/update_event", handlers.UpdateEvent)
	mux.HandleFunc("/delete_event", handlers.DeleteEvent)
	mux.HandleFunc("/events_for_day", handlers.EventsForDay)
	mux.HandleFunc("/events_for_week", handlers.EventsForWeek)
	mux.HandleFunc("/events_for_month", handlers.EventsForMonth)
}

// RunServer ../
func RunServer() {
	mux := http.NewServeMux()
	InitHandlers(mux)
	handler := middleware.Logging(mux)
	err := http.ListenAndServe("localhost:8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
