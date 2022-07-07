package carbon_producer

import (
	"context"
	"fmt"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	CarbonProducerRepository     repository.CarbonProducer
	UserCarbonProducerRepository repository.UserCarbonProducer
}

type Service interface {
	Find(ctx context.Context, filter *dto.FilterCarbonProducer, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CarbonProducer], error)
	CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, payload *dto.CalculateCarbonProducer) (string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		CarbonProducerRepository:     f.CarbonProducerRepository,
		UserCarbonProducerRepository: f.UserCarbonProducerRepository,
	}
}

func (s *service) Find(ctx context.Context, filter *dto.FilterCarbonProducer, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CarbonProducer], error) {

	CarbonProducers, info, err := s.CarbonProducerRepository.Find(ctx, filter, payload, &payload.Pagination)

	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.CarbonProducer])
	result.Datas = CarbonProducers
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, payload *dto.CalculateCarbonProducer) (string, error) {

	// get category carbon producer id
	data, err := s.CarbonProducerRepository.FindByID(ctx, carbon_producer_id)
	if err != nil {
		if err == constant.RecordNotFound {
			return "", res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// calculate emition
	var amount float32

	// transportation
	if data.CategoryCarbonProducerID == 1 {
		amount = float32(payload.JarakTempuh) * 0.321
	} else if data.CategoryCarbonProducerID == 2 {
		amount = float32(payload.JarakTempuh) * 0.421
	}

	data2 := model.UserCarbonProducer{UserID: user_id, CarbonProducerID: carbon_producer_id, Amount: amount}
	err2 := s.UserCarbonProducerRepository.Create(ctx, data2)
	if err2 != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err2)
	}

	message := fmt.Sprintf("%s %f %s", "emisi yang ada hasilkan adalah sejumlah", amount, "KgCO2")

	return message, nil
}
