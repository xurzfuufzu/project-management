package app

import (
	"context"
	"fmt"
	"net/http"
	"project-management-service/pkg/Logging"
	"time"

	"github.com/go-chi/chi/v5"
	"project-management-service/config"
	"project-management-service/internal/repository"
	"project-management-service/internal/service"
	"project-management-service/internal/transport/http/v1"
	"project-management-service/pkg/client/postgresql"
)

func Run() {
	cfg := config.GetConfig()
	logger := Logging.GetLogger()

	db, err := postgresql.NewClient(context.Background(), 3, cfg.DB)
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}

	//db.close()

	repos := repository.NewRepositories(db)
	userService := service.NewUserService(repos.Users)
	userHandler := v1.NewUserHandler(userService)

	projectService := service.NewProjectService(repos.Projects)
	ProjectHandler := v1.NewProjectHandler(projectService)

	router := chi.NewRouter()
	userHandler.RegisterRoutes(router)
	ProjectHandler.RegisterRoutes(router)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Infof("Starting server on %s", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Server failed: ", err)
	}
}
