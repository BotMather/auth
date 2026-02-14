package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JscorpTech/auth/internal/modules/auth"
	authHttp "github.com/JscorpTech/auth/internal/modules/auth/delivery/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	ctx, _ := context.WithCancel(context.Background())

	router := gin.Default()
	api := router.Group("/api")

	// Auth routes
	authRepository := auth.NewAuthRepository()
	authUsecase := auth.NewAuthUsecase(authRepository)
	authHandler := authHttp.NewAuthHandler(authUsecase)
	authHttp.RegisterAuthRoutes(api, authHandler)

	srv := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		logger.Info("Server ishga tushdi 🚀 :8080")
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-stop
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Serverni o'chirishda xatolik yuz berdi", zap.Error(err))
	}
	logger.Info("Server muvaffaqiyatli to'xtatildi")
}
