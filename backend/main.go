package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"slices"
	"unicode"

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
	var weights TeamStats
	err := json.Unmarshal([]byte(weightsForm), &weights)
	if err != nil {
		log.Fatal(err)
	}

	stats := getDefaultStats(vars["conference"])

	var teamData []Ranking
	for team, stat := range stats {
		teamData = append(teamData, Ranking{
			Team:  team,
			Score: calculateScore(stat, weights),
		})
	}
	slices.SortFunc(teamData, func(a, b Ranking) int {
		if n := cmp.Compare(b.Score, a.Score); n != 0 {
			return n
		} else {
			return cmp.Compare(b.Team, a.Team)
		}
	})

	w.Header().Set("Content-Type", "application/json")
	sendRes := json.NewEncoder(w).Encode(teamData)
	if sendRes != nil {
		log.Fatal(err)
	}
}

func calculateScore(stats TeamStats, weights TeamStats) float64 {
	statField := reflect.ValueOf(stats)
	weightField := reflect.ValueOf(weights)
	ranking := 0
	for i := 0; i < statField.NumField(); i++ {
		ranking += int(statField.Field(i).Int() * weightField.Field(i).Int())
	}
	return float64(ranking)
}

func getDefaultStats(query string) map[string]TeamStats {
	url := fmt.Sprintf(`https://api.collegefootballdata.com/stats/season?year=2023&&conference=%s`, query)
	res := getAPIData(url)
	var stats []Stat
	err := json.NewDecoder(res.Body).Decode(&stats)
	if err != nil {
		log.Fatal(err)
	}

	teamsUrl := fmt.Sprintf(`https://api.collegefootballdata.com/teams?conference=%s`, query)
	teamsRes := getAPIData(teamsUrl)
	var teams []Team
	err = json.NewDecoder(teamsRes.Body).Decode(&teams)
	if err != nil {
		log.Fatal(err)
	}

	statMap := make(map[string]TeamStats)
	for _, stat := range stats {
		tempTeam := statMap[stat.Team]
		value := reflect.ValueOf(&tempTeam).Elem()
		fieldName := stat.StatName
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(r[0])
		fieldName = string(r)
		field := value.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			switch v := stat.StatValue.(type) {
			case float64:
				field.SetInt(int64(v))
			case string:
				field.SetString(v)
			}
		}
		statMap[stat.Team] = tempTeam
	}
	return statMap
}

func getAdvancedTeamStats(team Team) AdvTeamStats {
	url := fmt.Sprintf(`https://api.collegefootballdata.com/stats/season/advanced?year=2023&team=%s`, team.School)
	res := getAPIData(url)
	var teamData []AdvTeamStats
	err := json.NewDecoder(res.Body).Decode(&teamData)
	if err != nil {
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
