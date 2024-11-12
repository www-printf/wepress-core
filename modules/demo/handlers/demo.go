package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/www-printf/wepress-core/modules/demo/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type DemoHandler struct {
	DemoUC usecases.DemoUsecase
}

func NewDemoHandler(g *echo.Group, demoUsecase usecases.DemoUsecase) {
	handler := &DemoHandler{
		DemoUC: demoUsecase,
	}
	api := g.Group("demo")

	api.GET("", wrapper.Wrap(handler.Get)).Name = "get:demo"
}

// GetDemo godoc
// @Summary Get demo
// @Description Get demo
// @Tags demo
// @Accept json
// @Produce json
// @Success      200  {object}  wrapper.SuccessResponse{data=domains.Demo}
// @Security     Bearer
// @Router       /demo [get]
func (h *DemoHandler) Get(c echo.Context) wrapper.Response {
	ctx := c.Request().Context()
	demo, err := h.DemoUC.Get(ctx)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPInternal}
	}

	return wrapper.Response{Data: demo, Status: http.StatusOK}
}
