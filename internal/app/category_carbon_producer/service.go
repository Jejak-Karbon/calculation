package category_carbon_producer

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	CategoryCarbonProducerRepository repository.CategoryCarbonProducer
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CategoryCarbonProducer], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		CategoryCarbonProducerRepository: f.CategoryCarbonProducerRepository,
	}
}

func (s *service) Find(ctx context.Context,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CategoryCarbonProducer], error) {

	CategoryCarbonProducers, info, err := s.CategoryCarbonProducerRepository.Find(ctx,payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.CategoryCarbonProducer])
	result.Datas = CategoryCarbonProducers
	result.PaginationInfo = *info

	return result, nil
}