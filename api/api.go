package api

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/pm-ashim-ghosh/ersolv/db"
)

type LogHandler struct{}

func (h LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		showLogs(w, r)

	} else if r.Method == "POST" {
		createLog(w, r)

	} else {
		// Unsupported method
		// TODO: Figure out appropriate JSON reponse.
	}
}

func showLogs(w http.ResponseWriter, r *http.Request) {
	var errorLog db.CodeLog
	var err error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	errorLog, err = db.GetCodeLog()
	if err != nil {
		log.Panic(err.Error())
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(errorLog); err != nil {
		log.Panic(err.Error())
	}
}

func createLog(w http.ResponseWriter, r *http.Request) {
	var errorLog db.CodeLog
	var err error
	var logId struct{
		Id int64
	}

	// Write the POST data to database.

	dec := json.NewDecoder(r.Body)
	if err = dec.Decode(&errorLog); err != nil {
		log.Panic(err.Error())
	}

	// TODO: Validate errorLog

	logId.Id, err = db.CreateCodeLog(errorLog)
	if err != nil {
		log.Panic(err.Error())
	}


	// Reply with logId

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err = enc.Encode(logId); err != nil {
		log.Panic(err.Error())
	}
}

func StartServer() {
	h := LogHandler{}
	http.Handle("/ersolv", h)

	http.ListenAndServe("localhost:4000", nil)
}
