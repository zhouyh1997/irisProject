package controller

import (
	"context"

	"CmsProject/service"
)

type OrderController struct {
	Ctx     context.Context
	Service service.OrderService
}
