package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"watchtower/database"
)

type DeleteRequest struct {
	ID         int64 `json:"ID"`
	ResidentID int64 `json:"ResidentID"`
}

type DeleteResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type Users struct {
	ID       int64  `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	IsAdmin  bool   `json:"IsAdmin"`
}

type UpdateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type CreateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

func (wt *WatchTowerConfiguration) ReadUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	UserData, err := database.ReadUsers(wt.WatchtowerDB)
	if err != nil {
		fmt.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserData)
}

func (wt *WatchTowerConfiguration) DeleteUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &DeleteRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.DeleteUser(wt.WatchtowerDB, req.ID)
	if err != nil {
		respondJSON(w, 200, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &DeleteResponse{
		Success: true,
		Message: "",
	})
}

func (wt *WatchTowerConfiguration) UpdateUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &Users{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	md5HashPass := MD5HashPassword(req.Password)

	err = database.UpdateUser(wt.WatchtowerDB, req.ID, req.Username, md5HashPass, req.IsAdmin)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &UpdateResponse{
		Success: true,
		Message: "",
	})
}

func (wt *WatchTowerConfiguration) CreateUsers(w http.ResponseWriter, r *http.Request) {

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
		respondJSON(w, 400, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	md5HashPass := MD5HashPassword(req.Password)

	err = database.CreateUser(wt.WatchtowerDB, req.Username, md5HashPass, req.IsAdmin)
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &CreateResponse{
		Success: true,
		Message: "",
	})
}
