package models

import (
    "time"
)

// LOAD MODELS
type API struct {
    ID      int      `json:"ID"`
    URL     string   `json:"URL"`
    Method  string   `json:"Method"`
    Headers [][]string `json:"Headers"`
    Body    string   `json:"Body"`
}

type APILoad struct {
    ID   int `json:"ID"`
    Iter int `json:"Iter"`
}

type LoadConfig struct {
	API      []API                    `json:"API"`
	Schedule map[int][]APILoad       `json:"Schedule"`
}

// WORKER POOL MODELS
type Job struct {
    ID   int
    API  API
    Iter int
}

type WorkerFunc func(job Job) (int, []byte, error, time.Duration)

// METRICS MODELS
type EventType int

type Event struct {
	Type EventType
	JobID int
	Latency time.Duration
	Err error
}

type Metrics struct {
	Completed int
	Errors int
	TotalLatency time.Duration
}