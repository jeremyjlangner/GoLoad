package main

import (
	"goload/load"
)

func main() {
	jsonData := []byte(`
		{
		"API": [
			{"ID": 1, "URL": "https://pokeapi.co/api/v2/pokemon/1", "Method": "GET", "Headers": [["Content-Type", "application/json"]], "Body": "testing"},
			{"ID": 2, "URL": "https://pokeapi.co/api/v2/pokemon/2", "Method": "GET", "Headers": [["Content-Type", "application/json"]], "Body": "testing"},
			{"ID": 3, "URL": "https://pokeapi.co/api/v2/pokemon/3", "Method": "GET", "Headers": [["Content-Type", "application/json"]], "Body": "testing"},
			{"ID": 4, "URL": "https://pokeapi.co/api/v2/pokemon/4", "Method": "GET", "Headers": [["Content-Type", "application/json"]], "Body": "testing"},
			{"ID": 5, "URL": "https://pokeapi.co/api/v2/pokemon/5", "Method": "GET", "Headers": [["Content-Type", "application/json"]], "Body": "testing"}
		],
			"Schedule": {
			"1": [{"ID": 2, "Iter": 5}],
			"2": [{"ID": 4, "Iter": 2}, {"ID": 5, "Iter": 1}],
			"3": [{"ID": 1, "Iter": 3}],
			"4": [{"ID": 3, "Iter": 2}, {"ID": 5, "Iter": 2}]
			}
		}`)

	load.Schedule(jsonData)
	//load.Example()
}