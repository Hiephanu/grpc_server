package main

import (
	"fmt"
	"log"
	"net"
	"simple-server/db"
	api "simple-server/grpcapi"
	"simple-server/server"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 "root",
		Passwd:               "#doanvanhiep",
		Addr:                 "127.0.1:3307",
		DBName:               "go_db",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Db connect error!")
	}
	server := server.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		fmt.Println("Error")
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	api.RegisterUserServiceServer(s, &api.GrpcApi{})

	log.Printf("Server is listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Printf("Server run on port %v", 8080)
}
