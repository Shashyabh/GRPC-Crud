package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shashyabh/mygo/repository"
	"github.com/shashyabh/mygo/server"
	"github.com/shashyabh/mygo/service"
)

func main() {
	db, err :=repository.NewPostgresRepository()
	if err != nil {
		log.Println("Error in connecting to Db ",err.Error())
	}
	defer db.Close()
	log.Println("Server started ")
	ser := service.NewService(db)
	//log.Fatal(server.ListenGRPC(ser,6000))
	const grpcAddress = "localhost:50051"
	go func() {
		if err := server.ListenGRPC(ser, grpcAddress); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	log.Println("gRPC server running on ", grpcAddress)

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Shutting down gracefully...")
}