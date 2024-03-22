package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mkrs2404/sre-bootcamp/pkg/config"
	"github.com/mkrs2404/sre-bootcamp/pkg/db"
	"github.com/mkrs2404/sre-bootcamp/pkg/handler"
	healthCheckGormImpl "github.com/mkrs2404/sre-bootcamp/pkg/repo/healthcheck/gormimpl"
	studentGormImpl "github.com/mkrs2404/sre-bootcamp/pkg/repo/student/gormimpl"
	"github.com/mkrs2404/sre-bootcamp/pkg/service"
)

func main() {
	cfg := config.Get()

	db, err := db.Connect(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	basePathRoute := engine.Group("/")
	router := engine.Group("/api/v1")

	studentRepo := studentGormImpl.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)
	studentHandler.RegisterRoutes(router)

	healthCheckRepo := healthCheckGormImpl.NewHealthCheckRepository(db)
	healthCheckService := service.NewHealthCheckService(healthCheckRepo)
	healthCheckHandler := handler.NewHealthCheckHandler(healthCheckService)
	healthCheckHandler.RegisterRoutes(basePathRoute)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
