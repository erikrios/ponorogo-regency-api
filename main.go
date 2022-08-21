package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/erikrios/ponorogo-regency-api/config"
	_ "github.com/erikrios/ponorogo-regency-api/docs"
	"github.com/erikrios/ponorogo-regency-api/pb"
	"github.com/erikrios/ponorogo-regency-api/repository"
	"github.com/erikrios/ponorogo-regency-api/rpc"
	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}

	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p", db)
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	provinceRepository := repository.NewProvinceRepositoryImpl(db)
	provinceService := service.NewProvinceServiceImpl(provinceRepository)

	grpcServer := grpc.NewServer()

	provinceServer := rpc.NewProvinceServer(provinceService)

	pb.RegisterProvinceServiceServer(grpcServer, provinceServer)

	reflection.Register(grpcServer)

	log.Printf("Start gRPC server at %s\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln(err.Error())
	}
}
