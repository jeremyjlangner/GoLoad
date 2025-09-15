package load

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"goload/workerpool"
	"goload/executor"
	"goload/models"
)

func LoadConfigFromJSON(jsonData []byte) ([]models.API, map[int][]models.APILoad, error) {
	var config models.LoadConfig
	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return config.API, config.Schedule, nil
}

// Example usage
func Example() {
	jsonData := []byte(`
{
  "API": [
    {"ID": 1, "URL": "https://pokeapi.co/api/v2/pokemon/1", "Method": "GET", "Headers": ["no-content"], "body": "testing"},
    {"ID": 2, "URL": "https://pokeapi.co/api/v2/pokemon/2", "Method": "GET", "Headers": ["no-content"], "body": "testing"},
    {"ID": 3, "URL": "https://pokeapi.co/api/v2/pokemon/3", "Method": "GET", "Headers": ["no-content"], "body": "testing"},
    {"ID": 4, "URL": "https://pokeapi.co/api/v2/pokemon/4", "Method": "GET", "Headers": ["no-content"], "body": "testing"},
    {"ID": 5, "URL": "https://pokeapi.co/api/v2/pokemon/5", "Method": "GET", "Headers": ["no-content"], "body": "testing"}
  ],
	"Schedule": {
	"1": [{"ID": 2, "Iter": 5}],
	"2": [{"ID": 4, "Iter": 2}, {"ID": 5, "Iter": 1}],
	"3": [{"ID": 1, "Iter": 5}],
	"4": [{"ID": 3, "Iter": 5}, {"ID": 5, "Iter": 5}]
	}
}`)

	apiObjects, schedule, err := LoadConfigFromJSON(jsonData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("API Objects:")
	for _, api := range apiObjects {
		fmt.Printf("%+v\n", api)
	}

	fmt.Println("\nSchedule:")
	for timeOffset, loads := range schedule {
		for _, load := range loads {
			fmt.Printf("Time %d: ID %d Iter %d\n", timeOffset, load.ID, load.Iter)
		}
	}
}

func Task() {
    fmt.Println("Task executed at:", time.Now())
}

func getMaxTime(sched map[int][]models.APILoad) (int) {
	var currMax = -1

	for key, _ := range sched {
		if key > currMax {
			currMax = key
		}
	}

	return currMax
}

func Schedule(jsonData []byte) {
    apis, schedule, err := LoadConfigFromJSON(jsonData)
    if err != nil {
        panic(err)
    }

    apiMap := make(map[int]models.API)
    for _, api := range apis {
        apiMap[api.ID] = api
    }

    start := time.Now()
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for now := range ticker.C {
        elapsed := int(now.Sub(start).Seconds())

        if loads, ok := schedule[elapsed]; ok {
            for _, load := range loads {
                api := apiMap[load.ID]
				job := models.Job{
					ID:   load.ID,
					API:  api,
					Iter: load.Iter,
				}
				go workerpool.Dispatcher(load.Iter, 2, executor.HttpRequest, job)
            }
        }
        if elapsed > getMaxTime(schedule) {
            break
        }
    }
}