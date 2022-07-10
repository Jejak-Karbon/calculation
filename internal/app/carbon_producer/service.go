package carbon_producer

import (
	"fmt"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
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
	CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, category_carbon_producer_id uint ,  payload *dto.CalculateCarbonProducer) (string, error)
	FindByID(ctx context.Context, carbon_producer_id uint) (*model.CarbonProducer, error)
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

func (s *service) FindByID(ctx context.Context, carbon_producer_id uint) (*model.CarbonProducer, error) {

	data, err := s.CarbonProducerRepository.FindByID(ctx, carbon_producer_id)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) CreateUserCarbonProducer(ctx context.Context, user_id uint, carbon_producer_id uint, category_carbon_producer_id uint , payload *dto.CalculateCarbonProducer) (string, error) {

	// calculate emition
	var koef float32
	var amount float32

	// transportation
	if category_carbon_producer_id == 1{

		if carbon_producer_id == 1{
			koef = 0.025	
		}else if carbon_producer_id == 2{
			koef = 0.1179
		}else if carbon_producer_id == 4{
			koef = 0.1689
		}

		if payload.BahanBakar == "Bensin"{
			amount = 2.6 
		}else if payload.BahanBakar == "Solar"{
			amount = 2.2 
		}

		amount = amount * float32(payload.JarakTempuh) * koef

	}else if category_carbon_producer_id == 2{
		amount = float32(payload.JumlahWatt) * float32(payload.LamaPenggunaan) / float32(1000) * float32(0.725)
	}

	data2 := model.UserCarbonProducer{UserID :user_id,CarbonProducerID:carbon_producer_id,Amount :amount}
	err2 := s.UserCarbonProducerRepository.Create(ctx, data2)
	if err2 != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err2)
	}

	message :=  fmt.Sprintf("%s %f %s", "emisi yang ada hasilkan adalah sejumlah", amount, "KgCO2")

	return message, nil
}
