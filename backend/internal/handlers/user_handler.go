package handlers

import (
	"database/sql"
	"net/http"
)

type Svc struct {
	db *sql.DB
}

func (s *Svc) UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world from my new server!"
	w.Write([]byte(msg))
}
