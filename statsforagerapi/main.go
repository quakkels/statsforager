package main

import (
	"context"
	"fmt"
	"net/http"

	"statsforagerapi/dataaccess"
	"statsforagerapi/middleware"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "stats"
	)

	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	connPool, err := dataaccess.NewConnPool(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	defer connPool.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /thing/{siteKey}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("siteKey")
		var dbversion string
		connPool.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)
		w.Write([]byte("you found me: " + id + "\n\n"))
		w.Write([]byte("<p>db version: " + dbversion + "</p>\n\n"))
	})

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: middlewareStack(router),
	}

	server.ListenAndServe()
}
