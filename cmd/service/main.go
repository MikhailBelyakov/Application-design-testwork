package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	initInfra "application-design-test-master/internal/infra/init"
)

func main() {
	var err error

	container := initInfra.NewContainer()

	container.Common.Router.Route("/orders", func(r chi.Router) {
		r.Post("/", container.Api.OrderCreateApi.Handle)
	})

	container.Common.Logger.LogInfo("Server listening on localhost:8080")
	err = http.ListenAndServe(":8080", container.Common.Router)
	if errors.Is(err, http.ErrServerClosed) {
		container.Common.Logger.LogInfo("Server closed")
	} else if err != nil {
		container.Common.Logger.LogInfo("Server failed: %s", err)
		os.Exit(1)
	}
}
