package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"microservice-process-file-go/pkg/configuration/configuration"
	"microservice-process-file-go/pkg/services/serviceReader"
	"net/http"
)

type Config struct {
}

type Server struct {
	Router        *mux.Router
	Logger        *zap.Logger
	Configuration *Config
	Handler       http.Handler
	Db            *gorm.DB
	ServiceReader *serviceReader.ServiceReader
}

func NewServer(
	logger *zap.Logger,
	configuration *Config,
	db *gorm.DB,
	serviceReader *serviceReader.ServiceReader,
) *Server {
	return &Server{
		Router:        mux.NewRouter(),
		Logger:        logger,
		Configuration: configuration,
		Db:            db,
		ServiceReader: serviceReader,
	}
}

func (s *Server) ListenAndServe(stopCh <-chan struct{}) {

	s.registerHandlers()
	s.Handler = s.Router
	s.Logger.Info("Starting server...")

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("%s:%s", configuration.Configuration.MicroserviceServer, configuration.Configuration.MicroservicePort), s.Handler); err != nil {
			s.Logger.Fatal("Error starting server", zap.Error(err))
		}
	}()
	s.Logger.Info("Server started")

	// wait for SIGTERM or SIGINT
	<-stopCh
	s.Logger.Info("Stopping server...")
}

func (s *Server) registerHandlers() {
	s.Router.HandleFunc("/load", s.LoadHandler).Methods("GET")
}
