package main

import (
	"context"
	"project-management-service/internal/config"
	"project-management-service/pkg/Logging"
	"project-management-service/pkg/client/postgresql"
)

func main() {
	logger := Logging.GetLogger()

	cfg := config.GetConfig()

	_, err := postgresql.NewClient(context.TODO(), 3, cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}
}
