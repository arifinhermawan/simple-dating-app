package app

import (
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/app/server"
	"github.com/arifinhermawan/simple-dating-app/internal/app/utils"
	"github.com/arifinhermawan/simple-dating-app/internal/repository/pgsql"
)

func NewApplication() {
	infra := server.NewInfra()
	cfg := infra.Config.GetConfig()

	// init db connection
	db, err := utils.InitDBConn(cfg.Database)
	if err != nil {
		log.Fatalf("[NewApplication] utils.InitDBConn() got error: %+v\n", err)
	}
	defer db.Close()

	// ----------------
	// |init app stack|
	// ----------------
	repoDB := pgsql.NewRepository(infra, db)
	services := server.NewService(repoDB, infra)
	usecases := server.NewUseCase(services)
	handlers := server.NewHandler(usecases, infra)

	// register handler
	utils.HandleRequest(handlers)
}
