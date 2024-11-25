package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type PrinterHandler struct {
	printerUC usecases.PrinterUsecase
}

func NewPrinterHandler(g *echo.Group, printerUC usecases.PrinterUsecase, middlewareMngr middlewares.MiddlewareManager) {
	h := &PrinterHandler{
		printerUC: printerUC,
	}

	api := g.Group("printers")
	api.Use(middlewareMngr.Auth())

	api.POST("/add", wrapper.Wrap(h.AddPrinter)).Name = "printer:add-printer"
	api.GET("/list", wrapper.Wrap(h.ListPrinter)).Name = "printer:list-printer"
	api.GET("/view-detail/:id", wrapper.Wrap(h.GetPrinter)).Name = "printer:view-detail"
	api.GET("/view-status/:id", wrapper.Wrap(h.ViewStatus)).Name = "printer:view-status"
}

// @Summary Add Printer
// @Description Add New Printer
// @Tags printers
// @Accept json
// @Produce json
// @Param request body dto.AddPrinterRequestBody true "Add Printer Request Body"
// @Success      201  {object}  wrapper.SuccessResponse{data=dto.PrinterResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /printers/add [post]
func (h *PrinterHandler) AddPrinter(c echo.Context) wrapper.Response {
	req := &dto.AddPrinterRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, err := h.printerUC.AddPrinter(c.Request().Context(), req)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: resp, Status: http.StatusCreated}
}

// @Summary List Printer
// @Description List All Printers
// @Tags printers
// @Accept json
// @Produce json
// @Param cluster_id query string false "Cluster ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.ListPrinterResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /printers/list [get]
func (h *PrinterHandler) ListPrinter(c echo.Context) wrapper.Response {
	clusterIDStr := c.QueryParam("cluster_id")
	if clusterIDStr == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	clusterID, err := strconv.ParseUint(clusterIDStr, 10, 32)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, erro := h.printerUC.ListPrinter(c.Request().Context(), uint(clusterID))
	if erro != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}

// @Summary View Printer Detail
// @Description View Printer Detail
// @Tags printers
// @Produce json
// @Param id path string true "Printer ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.PrinterResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /printers/view-detail/{id} [get]
func (h *PrinterHandler) GetPrinter(c echo.Context) wrapper.Response {
	idStr := c.Param("id")
	if idStr == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, erro := h.printerUC.GetPrinter(c.Request().Context(), uint(id))
	if erro != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}

// @Summary View Printer Status
// @Description View Printer Status
// @Tags printers
// @Produce json
// @Param id path string true "Printer ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.PrinterStatusResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /printers/view-status/{id} [get]
func (h *PrinterHandler) ViewStatus(c echo.Context) wrapper.Response {
	idStr := c.Param("id")
	if idStr == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, erro := h.printerUC.ViewStatus(c.Request().Context(), uint(id))
	if erro != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}
