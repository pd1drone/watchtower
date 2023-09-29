package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"watchtower/database"
)

type WaterLevelData struct {
	WaterLevel string `json:"WaterLevel"`
}

type AddWaterLevelResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

func (wt *WatchTowerConfiguration) GetWaterLevelHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	waterLvlHistory, err := database.GetWaterLevelHistory(wt.WatchtowerDB)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, waterLvlHistory)
}

func (wt *WatchTowerConfiguration) AddWaterLevelHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &WaterLevelData{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, &AddWaterLevelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	f, err := strconv.ParseFloat(req.WaterLevel, 64)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, &AddWaterLevelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.AddWaterLevelHistory(wt.WatchtowerDB, f)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, &AddWaterLevelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &AddWaterLevelResponse{
		Success: true,
		Message: "",
	})
}
