package main

import (
	"context"
	"frontend_go/api"
	"frontend_go/notify"
	"frontend_go/store"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("PROFILE_DB")
	if path == "" {
		path = "profile.db"
	}
	db, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	svc := api.NewService(db, notify.NewConsole())
	srv := &http.Server{Addr: ":8080", Handler: api.Routes(svc)}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print(err)
		}
	}()
	<-context.Background().Done()
}
