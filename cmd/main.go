package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	ginhttp "social_network/internal/http"
	"social_network/internal/http/handlers"
	"social_network/internal/repository"
	"social_network/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := mustGetEnv("DATABASE_URL")

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// repos
	userRepo := repository.NewUserRepository(pool)
	// services
	userSvc := service.NewUserService(userRepo)
	// handlers
	authH := handlers.NewAuthHandler(userSvc)
	userH := handlers.NewUserHandler(userSvc)

	// router
	r := ginhttp.NewRouter(authH, userH)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("listening on :8080")
	log.Fatal(srv.ListenAndServe())
}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing env %s", k)
	}
	return v
}
