package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	store := NewStore(s.db)
	userRoute := NewHandler(store)
	userRoute.RegsterRoutes(subRouter)

	initStorage(s.db)
	return http.ListenAndServe(s.addr, router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Db connected!")
}
