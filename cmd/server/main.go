package main

import (
	"fmt"
	"github.com/alfinkly/api-gateway/internal/adapters/database"
	"github.com/alfinkly/api-gateway/internal/adapters/http"
	"github.com/alfinkly/api-gateway/internal/config"
	"github.com/alfinkly/api-gateway/internal/repository/sqlc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		panic("У-у-у-у, не нашёл конфигов (┬┬﹏┬┬)")
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		fmt.Println(err)
		panic("А-а-а-а база мертва-а-а (┬┬﹏┬┬)")
	}

	queries := sqlc.New(db)

	http.StartServer(queries, cfg.HTTPAddr)
}
