package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createMysql(host, user, password string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/github_dev?parseTime=true", user, password, host)
	log.Printf("Connecting to mysql on %s ... \n", host)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	log.Printf("Connected to mysql on %s \n", host)
	return sqlDB, nil
}

func main() {
	log.Println("Starting server ...")
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := createMysql(cfg.MysqlHost, cfg.MysqlUser, cfg.MysqlPassword)
	if err != nil {
		log.Fatal(err)
	}
	client := NewClient(cfg)
	profileRepo := NewProfileRepo(sqlDB)
	handler := NewHandler(cfg, client, profileRepo)
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/auth", handler.Auth)
	http.HandleFunc("/auth/callback", handler.Callback)

	log.Println("listening on port 9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
