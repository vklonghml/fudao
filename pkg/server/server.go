package server

import (
	"log"
	"net/http"
	"fmt"

	"fudao/cmd/options"
	"fudao/pkg/common/db"
	"fudao/pkg/router"
	"fudao/pkg/spider"
	"fudao/pkg/store"
)

func StartServer(opts *options.Options) error {
	err := db.CreateDBIfNeeded(opts.DBSource, opts.DBName)
	if err != nil {
		log.Println("failed to create database with err:", err)
		return err
	}
	// register db
	store.RegisterDB(opts.DBSource, opts.DBName)
	// init db
	err = store.InitDB()
	if err != nil {
		log.Println("failed to init database with err:", err)
		return err
	}

	go spider.DownCourse()

	router.Register()
	log.Println("start listening on add: %s, port: %s", opts.Address, opts.Port)
	addr := fmt.Sprintf("%s:%s", opts.Address, opts.Port)
	log.Println(addr)
	log.Fatal(http.ListenAndServe(":12345", nil))
	return nil
}
