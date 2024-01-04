package rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (wt *WatchTowerConfiguration) check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Routes() {
	// Check if the log file exists
	if _, err := os.Stat("/root/watchtower/http_logs.log"); os.IsNotExist(err) {
		// Create a log file
		logFile, err := os.Create("/root/watchtower/http_logs.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
	}

	// Open the log file
	logFile, err := os.OpenFile("/root/watchtower/http_logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// // Create a new logger using the log file
	logger := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(logFile, "", log.LstdFlags), // Use the log file as the output
		NoColor: true,
	})

	log.Print("Starting Watchtower Backend Service.....")
	r := chi.NewRouter()

	newWatchtower, err := New()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(logger)
	r.Use(func(next http.Handler) http.Handler {
		return logPayloads(logFile, next)
	})

	// Check route if service is running
	r.Get("/check", newWatchtower.check)

	// Login route
	r.Post("/login", newWatchtower.Login)

	// Register route
	r.Post("/register", newWatchtower.RegisterUser)

	// add water level history
	r.Post("/history", newWatchtower.AddWaterLevelHistory)

	// get water level history
	r.Get("/getHistory", newWatchtower.GetWaterLevelHistory)

	// get water level history
	r.Get("/waterlvl", newWatchtower.GetWaterLevelRealTimeData)

	// CRUD USERS
	r.Get("/users", newWatchtower.ReadUsers)
	r.Post("/users", newWatchtower.CreateUsers)
	r.Delete("/users", newWatchtower.DeleteUsers)
	r.Put("/users", newWatchtower.UpdateUsers)

	log.Fatal(http.ListenAndServe("0.0.0.0:8090", r))
}

func logPayloads(logFile *os.File, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Record start time of request processing
		startTime := time.Now()

		// Call the next middleware/handler in the chain and capture response
		responseRecorder := httptest.NewRecorder()
		next.ServeHTTP(responseRecorder, r)

		// Record end time of request processing
		endTime := time.Now()

		// Read request payload
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Restore request body after reading
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// Log the request and response payloads
		log.SetOutput(logFile)
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		if len(requestBody) > 0 {
			log.Printf("Request Payload: %s\n", string(requestBody))
		} else {
			log.Println("Request Payload: [Empty]")
		}

		// Read response payload
		responseBody := responseRecorder.Body.Bytes()

		if len(responseBody) > 0 {
			log.Printf("Request Response: %d %s\nPayload: %s\n", responseRecorder.Code, http.StatusText(responseRecorder.Code), string(responseBody))
		} else {
			log.Printf("Request Response: %d %s\nPayload: [Empty]\n", responseRecorder.Code, http.StatusText(responseRecorder.Code))
		}
		log.Printf("Duration: %v\n", endTime.Sub(startTime))

		// Write response back to original response writer
		for k, v := range responseRecorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(responseRecorder.Code)
		w.Write(responseRecorder.Body.Bytes())

		log.Printf("\n")
	})
}
