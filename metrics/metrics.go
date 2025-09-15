package metrics

import (
	"fmt"
	"goload/models"
)

const (
	EventCompleted models.EventType = iota
	EventError
	EventLatency
)

func Collector(events <-chan models.Event, done chan<- models.Metrics) {
	m := models.Metrics{}
	for e := range events {
		switch e.Type {
		case EventCompleted:
			m.Completed++
		case EventError:
			m.Errors++
		}
		m.TotalLatency += e.Latency
		record(e)
	}
	done <- m
}

func record(e models.Event) {
	fmt.Printf("Event JobID: %d EventType: %v Latency: %v Error: %v\n", e.JobID, e.Type, e.Latency, e.Err)
}