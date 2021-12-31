package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/henrjan/microservice/configs"
	pb "github.com/henrjan/microservice/internal/proto"
	"github.com/henrjan/microservice/pkg/driver"
	"github.com/henrjan/microservice/pkg/handler/v2"
	"github.com/henrjan/microservice/pkg/repository"
	"github.com/henrjan/microservice/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db           = initDB()
	accessRepo   = repository.NewAccessRepository(db)
	accessSrv    = service.NewAccessService(accessRepo)
	movieDriver  = driver.NewMovieDriver()
	movieSrv     = service.NewMovieService(movieDriver)
	movieHandler = handler.NewMovieHandler(movieSrv, accessSrv)
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterSearchMovieServer(s, movieHandler)

	// Serve gRPC server on Port 8080
	go s.Serve(lis)

	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx1,
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("Failed to dial server:", err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel2()
	gwmux := runtime.NewServeMux()
	err = pb.RegisterSearchMovieHandler(ctx2, gwmux, conn)
	if err != nil {
		fmt.Println("Failed to register gateway:", err)
	}

	// Serving gRPC-Gateway on port :8081
	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: gwmux,
	}
	gwServer.ListenAndServe()
}

func initDB() *gorm.DB {
	dsn := configs.GetMySqlDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
