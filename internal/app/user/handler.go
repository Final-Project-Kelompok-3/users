package user

import (
	"strconv"

	res "github.com/Final-Project-Kelompok-3/authentications/pkg/util/response"

	"github.com/Final-Project-Kelompok-3/authentications/internal/dto"
	"github.com/Final-Project-Kelompok-3/authentications/internal/factory"
	_ "github.com/go-playground/validator/v10"
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


func (h *handler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	
	token, err := h.service.Login(c.Request().Context(), email, password)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := map[string]interface{}{"token":token}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindAll(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get users success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetByID(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Create(c echo.Context) error {

	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.Create(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	
	payload := new(dto.UpdateUserRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	strID := c.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	result, err := h.service.Update(c.Request().Context(), uint(ID), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Delete(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.Delete(c.Request().Context(), payload.ID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}