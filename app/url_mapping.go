package app

import (
	"github.com/GrolimundSolutions/2vault/controllers/ping"
	"github.com/GrolimundSolutions/2vault/controllers/vault"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/2vault", vault.Create)
}
