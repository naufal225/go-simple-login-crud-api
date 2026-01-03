package main

import (
	"log"

	"github.com/naufal225/go-simple-login-crud-api/internal/config"
	"github.com/naufal225/go-simple-login-crud-api/internal/db"
)

func main() {
	cfg := config.Load()
	dbConn := db.Connect(cfg)

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal("Error saat konek db:", err)
	}

	defer sqlDB.Close()

	log.Println("🚀 application bootstrapped successfully")
}
