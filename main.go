package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	api "github.com/viardant/leetcode-solutions/webapi"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	router := api.Handler(log)
	api.Route(router)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	shutdownSign := make(chan os.Signal, 1)
	signal.Notify(shutdownSign, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)

	go func() {
		log.Infof("API Server listetning on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		log.WithError(err).Fatal("an error has occurred")
	case <-shutdownSign:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.WithError(err).Fatal("server forced to stop")
		}
		log.Info("server stopped gracefully")
	}
}
