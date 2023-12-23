package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"watchtower/database"
)

type RegisterUserRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (wt *WatchTowerConfiguration) RegisterUser(w http.ResponseWriter, r *http.Request) {

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

	req := &Users{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &RegisterResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.RegisterUser(wt.WatchtowerDB, req.Username, req.Password)
	if err != nil {
		respondJSON(w, 200, &RegisterResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &RegisterResponse{
		Success: true,
		Message: "",
	})
}
