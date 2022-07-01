package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	CategoryCarbonProducerRepository repository.CategoryCarbonProducer
	CarbonProducerRepository repository.CarbonProducer
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		CategoryCarbonProducerRepository: repository.NewCategoryCarbonProducer(db),
		CarbonProducerRepository: repository.NewCarbonProducer(db),
	}
}
