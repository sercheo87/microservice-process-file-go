package main

import (
	"fmt"
	"go.uber.org/zap"
	"microservice-process-file-go/pkg/configuration/configuration"
	"microservice-process-file-go/pkg/configuration/loggerConfiguration"
	"microservice-process-file-go/pkg/controllers"
	"microservice-process-file-go/pkg/repository/db"
	"microservice-process-file-go/pkg/services/serviceReader"
	"os"
	"os/signal"
	"syscall"
)

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func closeLoggerHandler() func(logger *zap.Logger) {
	return func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {

	configuration.LoadAppConfiguration()
	// configure logging
	logger, _ := loggerConfiguration.ApplyLoggerConfiguration(configuration.Configuration.LogLevel)
	defer closeLoggerHandler()(logger)

	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()
	logger.Info("Logger Configuration applied")

	logger.Info("Starting server",
		zap.String("Microservice", configuration.Configuration.MicroserviceName),
		zap.String("Host", configuration.Configuration.MicroserviceServer),
		zap.String("Port", configuration.Configuration.MicroservicePort))

	var srvCfg controllers.Config
	db := db.GetConnectionClient(logger)
	server := controllers.NewServer(
		logger,
		&srvCfg,
		db,
		serviceReader.NewServiceReader(logger, db),
	)

	stopCh := SetupSignalHandler()

	logger.Debug("Creating routers")
	server.ListenAndServe(stopCh)
}
