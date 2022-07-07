package user_carbon_producer

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	UserCarbonProducerRepository repository.UserCarbonProducer
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest, id uint) (*dto.SearchGetResponse[model.UserCarbonProducer], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserCarbonProducerRepository: f.UserCarbonProducerRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest, id uint) (*dto.SearchGetResponse[model.UserCarbonProducer], error) {

	UserCarbonProducer, info, err := s.UserCarbonProducerRepository.Find(ctx, payload, id, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// var datas []dto.UserResponse

	// for _, user := range users {

	// 	datas = append(datas, dto.UserResponse{
	// 		ID:    user.ID,
	// 		Name:  user.Name,
	// 		Email: user.Email,
	// 	})

	// }

	result := new(dto.SearchGetResponse[model.UserCarbonProducer])
	result.Datas = UserCarbonProducer
	result.PaginationInfo = *info

	return result, nil
}

