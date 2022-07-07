package carbon_producer

import (
	_ "fmt"
	"strconv"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	filter := new(dto.FilterCarbonProducer)

	if err := c.Bind(filter); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(filter); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	
	result, err := h.service.Find(c.Request().Context(),filter, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get carbon producer success", &result.PaginationInfo).Send(c)
}

func (h *handler) Calculate(c echo.Context) error {

	payload := new(dto.CalculateCarbonProducer)

	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	id,_:= strconv.Atoi(c.Param("carbon_producer_id"))
	payloadToken := middleware.GetIDFromToken(c)
	
	var user_id uint = uint(payloadToken.(float64))
	var carbon_producer_id uint = uint(id)

	result, err := h.service.CreateUserCarbonProducer(c.Request().Context(), user_id, carbon_producer_id, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)

}