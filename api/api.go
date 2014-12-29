package api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/pm-ashim-ghosh/ersolv/db"
)

type LogHandler struct{}

func (h LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errorLog db.CodeLog
	var err error

	if r.Method == "GET" {
		errorLog, err = db.GetCodeLog()
		if err != nil {
			fmt.Println(err)
		}

		out := fmt.Sprintf("%v\n", errorLog)
		if _, err = w.Write([]byte(out)); err != nil {
			log.Panic(err.Error())
		}


	} else if r.Method == "POST" {

	} else {
		// Unsupported method
	}
}

func StartServer() {
	h := LogHandler{}
	http.Handle("/ersolv", h)

	http.ListenAndServe("localhost:4000", nil)
}
