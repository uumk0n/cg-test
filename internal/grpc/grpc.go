package grpc

import (
	"cg-test/config"
	"cg-test/github.com/uumkon/cg-test/proto/user"
	"cg-test/github.com/uumkon/cg-test/proto/weather"
	"cg-test/internal/services"
	"cg-test/internal/storage/repo"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.DbName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repo.NewUserRepo(db)
	actionsRepo := repo.NewActionsRepo(db)

	userServiceServer := services.NewUserServiceServer(userRepo)
	userListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	userServer := grpc.NewServer()
	user.RegisterUserServiceServer(userServer, userServiceServer)

	weatherServiceServer := services.NewWeatherServiceServer(userRepo, actionsRepo, &cfg)
	weatherListener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	weatherServer := grpc.NewServer()
	weather.RegisterWeatherServiceServer(weatherServer, weatherServiceServer)

	go func() {
		if err := userServer.Serve(userListener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	go func() {
		if err := weatherServer.Serve(weatherListener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Println("gRPC servers started successfully")

	select {}
}
