package main

import (
	"fmt"
	"log"
	"simple-server/db"
	"simple-server/server"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
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

	fmt.Printf("Server run on port %v", 8080)
}
