package carbon_producer

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	CarbonProducerRepository repository.CarbonProducer
}

type Service interface {
	Find(ctx context.Context,filter *dto.FilterCarbonProducer,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CarbonProducer], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		CarbonProducerRepository: f.CarbonProducerRepository,
	}
}

func (s *service) Find(ctx context.Context,filter *dto.FilterCarbonProducer,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CarbonProducer], error) {

	CarbonProducers, info, err := s.CarbonProducerRepository.Find(ctx,filter,payload, &payload.Pagination)
	
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.CarbonProducer])
	result.Datas = CarbonProducers
	result.PaginationInfo = *info

	return result, nil
}
