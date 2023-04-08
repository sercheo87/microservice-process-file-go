package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"microservice-process-file-go/pkg/configuration/configuration"
	"microservice-process-file-go/pkg/models"
)

func GetConnectionClient(logger *zap.Logger) *gorm.DB {
	logger.Info("Initializing database connection client")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Guayaquil",
		configuration.Configuration.Database.Host,
		configuration.Configuration.Database.Username,
		configuration.Configuration.Database.Password,
		configuration.Configuration.Database.Dbname,
		configuration.Configuration.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	})
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
	}

	// Migrate the schema to ensure that the table exists
	err = db.AutoMigrate(&models.SalesModel{})
	if err != nil {
		panic(err)
	}

	logger.Info("Connected to database")
	return db
}
