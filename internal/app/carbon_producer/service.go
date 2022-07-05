package carbon_producer

import (
	"fmt"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	CarbonProducerRepository repository.CarbonProducer
	UserCarbonProducerRepository repository.UserCarbonProducer
}

type Service interface {
	Find(ctx context.Context,filter *dto.FilterCarbonProducer,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CarbonProducer], error)
	CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, payload *dto.CalculateCarbonProducer) (string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		CarbonProducerRepository: f.CarbonProducerRepository,
		UserCarbonProducerRepository: f.UserCarbonProducerRepository,
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

func (s *service) CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, payload *dto.CalculateCarbonProducer) (string, error) {

	// calculate emition
	var amount float32
	amount = float32(payload.JarakTempuh) * 0.321

	data := model.UserCarbonProducer{UserID :user_id,CarbonProducerID:carbon_producer_id,Amount :amount}
	fmt.Printf("%+v\n", data)
	err := s.UserCarbonProducerRepository.Create(ctx, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	message :=  fmt.Sprintf("%s %f %s", "emisi yang ada hasilkan adalah sejumlah", amount, "KgCO2")

	return message, nil
}
