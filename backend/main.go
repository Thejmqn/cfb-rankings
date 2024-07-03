package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/data/{conference}", requestHandler).Methods(http.MethodPost, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))
	fmt.Println("Server started on port 8080")
	http.ListenAndServe("localhost:8080", router)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	vars := mux.Vars(r)
	weightsForm := r.FormValue("weights")
	var weights Weights
	err := json.Unmarshal([]byte(weightsForm), &weights)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf(`https://api.collegefootballdata.com/teams?conference=%s`, vars["conference"])
	res := getAPIData(url)

	var teams []Team
	err = json.NewDecoder(res.Body).Decode(&teams)
	if err != nil {
		log.Fatal(err)
	}

	var stats []TeamStats
	for _, team := range teams {
		stats = append(stats, getTeamStats(team))

	}

	var teamData []Ranking
	for _, statLine := range stats {
		teamData = append(teamData, Ranking{
			Team:  statLine.Team,
			Score: calculateScore(statLine, weights),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	sendRes := json.NewEncoder(w).Encode(teamData)
	if sendRes != nil {
		log.Fatal(err)
	}
}

func calculateScore(stats TeamStats, weights Weights) float64 {
	return stats.Defense.PPA*float64(weights.PPA) + stats.Offense.PPA*float64(weights.PPA)
}

func getTeamStats(team Team) TeamStats {
	url := fmt.Sprintf(`https://api.collegefootballdata.com/stats/season/advanced?year=2023&team=%s`, team.School)
	res := getAPIData(url)
	var teamData []TeamStats
	err := json.NewDecoder(res.Body).Decode(&teamData)
	if err != nil {
		fmt.Println("DDD")
		log.Fatal(err)
	}
	if len(teamData) != 1 {
		log.Fatal("Failed to get team data.")
	}
	return teamData[0]
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func getAPIData(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", readAPIKey())
	req.Header.Set("accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return res
}

func readAPIKey() string {
	const keyFile = "key.txt"
	key, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatal("Failed to get API key.")
	}
	return "Bearer " + string(key)
}
