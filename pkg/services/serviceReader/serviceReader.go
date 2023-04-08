package serviceReader

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"microservice-process-file-go/pkg/models"
	"microservice-process-file-go/pkg/utils/parserFile"
	"time"
)

type Handler interface {
	ReadFile(filePath string)
}

type ServiceReader struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (s ServiceReader) ReadFile(filePath string) {
	start := time.Now()

	s.logger.Info("Reading file")
	r := parserFile.ParserFile(filePath)

	var salesModelCollections []models.SalesModel

	s.logger.Info("Converting data to model")
	for _, v := range r {
		saleModel := models.SalesModel{OrderID: v.OrderId, Region: v.Region, UnitSold: v.UnitSold}
		salesModelCollections = append(salesModelCollections, saleModel)
	}

	s.logger.Info("Saving data to database")
	s.db.CreateInBatches(salesModelCollections, 1000)

	elapsed := time.Since(start)
	s.logger.Info("File processed", zap.Int64("duration", elapsed.Milliseconds()))
}

func NewServiceReader(logger *zap.Logger, db *gorm.DB) *ServiceReader {
	return &ServiceReader{
		logger: logger,
		db:     db,
	}
}
