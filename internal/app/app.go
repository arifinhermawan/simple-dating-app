package app

import (
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/app/server"
	"github.com/arifinhermawan/simple-dating-app/internal/app/utils"
)

func NewApplication() {
	infra := server.NewInfra()
	cfg := infra.Config.GetConfig()

	// init db connection
	err := utils.InitDBConn(cfg.Database)
	if err != nil {
		log.Fatalf("[NewApplication] utils.InitDBConn() got error: %+v\n", err)
	}

	// ----------------
	// |init app stack|
	// ----------------

	_ = server.NewService()
	_ = server.NewUseCase()
	_ = server.NewHandler()

}
