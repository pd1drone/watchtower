package rest

import (
	"log"
	"net/http"
	"watchtower/database"
)

func (wt *WatchTowerConfiguration) GetWaterLevelRealTimeData(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	waterLevelDataRealtime, err := database.GetWaterLevelRealTimeData(wt.WatchtowerDB)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, waterLevelDataRealtime)
}
