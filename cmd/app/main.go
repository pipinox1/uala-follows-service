package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	_ "go.uber.org/automaxprocs/maxprocs"
	httpServer "net/http"
	"os"
	"os/signal"
	"syscall"
	"uala-followers-service/cmd/http"
	"uala-followers-service/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", err))
	}

	level, err := zerolog.ParseLevel(zerolog.InfoLevel.String())
	if err != nil {
		panic(fmt.Errorf("invalid LogLevel value retrieved from environment: %w", err))
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	_ = zerolog.New(os.Stdout).
		Level(level).
		With().
		Timestamp().
		Str("service", cfg.ServiceName).
		Logger()

	dependencies, err := config.BuildDependencies(*cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error building dependencies: %w", err))
	}

	go func() {
		router := http.SetupRouterAndRoutes(cfg, dependencies)
		go func() {
			fmt.Println("starting server on port: ", cfg.Port)
			if err := httpServer.ListenAndServe(
				fmt.Sprintf(":%s", cfg.Port),
				router,
			); err != nil {
				fmt.Println("error starting server")
				panic(err)
			}
		}()

		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

		<-exit
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
