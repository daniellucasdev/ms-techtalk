package controller

import (
	"github.com/braiphub/ms-tech-talk/internal/domain/service"
	"github.com/braiphub/ms-tech-talk/internal/infra/anticorruption/frontend"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type SubscriptionController struct {
	service   *service.SubscriptionService
	validator *validator.Validate
}

func NewSubscriptionController(
	service *service.SubscriptionService,
) *SubscriptionController {
	return &SubscriptionController{
		service:   service,
		validator: validator.New(),
	}
}

func (sc *SubscriptionController) CreateSubscription(c echo.Context) error {
	var dto frontend.CreateSubscriptionRequestDTO

	if err := c.Bind(&dto); err != nil {
		return errors.Wrap(err, "bind")
	}

	if err := sc.validator.Struct(dto); err != nil {
		return errors.Wrap(err, "validate")
	}

	subscription := frontend.TranslateCreateSubscriptionRequestToEntity(dto)

	if err := sc.service.Create(c.Request().Context(), &subscription); err != nil {
		return errors.Wrap(err, "create subscription")
	}

	res := frontend.TranslateSubscriptionToCreateSubscriptionResponse(&subscription)

	return c.JSON(http.StatusCreated, res)
}
