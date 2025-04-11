package main

import (
	"context"
	"log"

	"github.com/braiphub/ms-tech-talk/cmd/http/components"
	"github.com/braiphub/ms-tech-talk/internal/api/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	appCtx := context.Background()

	stp, err := components.SetUp(appCtx)
	if err != nil {
		return errors.Wrap(err, "components")
	}

	// domain event listeners
	stp.Container.EventHandler().StartListeners()

	// integration event consumers
	stp.Container.MsProductsAdapter().StartConsumers(appCtx)

	stp.Container.MsOrdersAdapter().StartConsumers(appCtx)


	// api handler
	apiServer := http.NewAPIServer(stp.Container.Logger())
	apiServer.ConfigureRoutes(
		stp.Container.HealthController(),
	)
	apiServer.Start(80) //nolint:mnd

	return nil
}
