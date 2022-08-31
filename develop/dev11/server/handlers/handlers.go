package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/models"
	"strconv"
	"time"
)

var Events map[int]models.Event

func init() {
	Events = make(map[int]models.Event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var event models.Event
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &event)
		if err != nil {
			fmt.Println(err)
		}
		id, _ := strconv.Atoi(event.Result.Id)
		Events[id] = event
		text, _ := getJson("success", 0)
		_, _ = w.Write(text)
	} else {
		w.WriteHeader(500)
		text, _ := getJson("only post method!", 1)
		_, err := w.Write(text)
		if err != nil {
			log.Println(err)
		}
	}
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var event models.Event
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &event)
		if err != nil {
			fmt.Println(err)
		}
		id, _ := strconv.Atoi(event.Result.Id)
		if id > len(Events) {
			w.WriteHeader(400)
			text, _ := getJson("too big id!", 1)
			_, _ = w.Write(text)
			return
		}
		var tick bool
		for k, _ := range Events {
			if k == id {
				Events[k] = event
				tick = true
				break
			}
		}
		if tick {
			text, _ := getJson("success", 0)
			_, _ = w.Write(text)
		} else {
			w.WriteHeader(400)
			text, _ := getJson("no event with such id!", 1)
			_, _ = w.Write(text)
		}
	} else {
		w.WriteHeader(500)
		text, _ := getJson("only post method!", 1)
		_, err := w.Write(text)
		if err != nil {
			log.Println(err)
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var event models.Event
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &event)
		if err != nil {
			fmt.Println(err)
		}
		id, _ := strconv.Atoi(event.Result.Id)
		if id > len(Events) {
			w.WriteHeader(400)
			text, _ := getJson("too big id!", 1)
			_, err := w.Write(text)
			if err != nil {
				log.Println(err)
			}
			return
		}
		delete(Events, id)
		text, _ := getJson("success", 0)
		_, _ = w.Write(text)
	} else {
		w.WriteHeader(500)
		text, _ := getJson("only post method!", 1)
		_, _ = w.Write(text)
	}
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		resp, err := getEvents(1)
		if err != nil {
			w.WriteHeader(503)
			text, _ := getJson("get events went wrong", 1)
			_, _ = w.Write(text)
		}
		_, _ = w.Write(resp)
	}
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		resp, err := getEvents(7)
		if err != nil {
			w.WriteHeader(503)
			text, _ := getJson("get events went wrong", 1)
			_, _ = w.Write(text)
		}
		_, _ = w.Write(resp)
	}
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		resp, err := getEvents(30)
		if err != nil {
			w.WriteHeader(503)
			text, _ := getJson("get events went wrong", 1)
			_, _ = w.Write(text)
		}
		_, _ = w.Write(resp)
	}
}

func getEvents(days int) ([]byte, error) {
	ets := models.Events{}
	curTime := time.Now()
	var untilTime time.Time
	switch days {
	case 1:
		untilTime = curTime.Add(24 * time.Hour)
	case 7:
		untilTime = curTime.Add(168 * time.Hour)
	case 30:
		untilTime = curTime.Add(720 * time.Hour)
	}
	//fmt.Println(formattedTime)
	//fmt.Println(curParsedTime)
	//fmt.Println(untilTime)
	for k, v := range Events {
		parsedTime, _ := time.Parse("2006-01-02", v.Result.Datetime)
		fmt.Println(parsedTime)
		if parsedTime.Before(untilTime) {
			ets = append(ets, Events[k])
		}
	}
	return json.Marshal(ets)
}

func getJson(str string, code int) ([]byte, error) {
	switch code {
	case 0:
		res := models.Success{}
		res.Result.Result = str
		return json.Marshal(res)
	case 1:
		err := models.Err{}
		err.Error.Error = str
		return json.Marshal(err)
	}
	return nil, nil
}
