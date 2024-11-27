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

	printer := g.Group("printers")
	printer.Use(middlewareMngr.Auth())

	printer.POST("/add", wrapper.Wrap(h.AddPrinter)).Name = "printer:add-printer"
	printer.GET("/list", wrapper.Wrap(h.ListPrinter)).Name = "printer:list-printer"
	printer.GET("/info/:id", wrapper.Wrap(h.GetPrinter)).Name = "printer:view-info"
	printer.GET("/monitor/:id", wrapper.Wrap(h.ViewPrinterStatus)).Name = "printer:view-status"

	cluster := g.Group("clusters")
	cluster.GET("/list", wrapper.Wrap(h.ListCluster)).Name = "printer:list-cluster"

	job := g.Group("print-jobs")
	job.POST("/submit", wrapper.Wrap(h.SubmitPrintJob)).Name = "printer:submit-printjob"
	job.POST("/cancel/:id", wrapper.Wrap(h.CancelPrintJob)).Name = "printer:cancel-printjob"
	job.GET("/list/:printerid", wrapper.Wrap(h.ListPrintJob)).Name = "printer:list-printjob"
	job.GET("/monitor/:id", wrapper.Wrap(h.ViewJobStatus)).Name = "printer:view-printjob-status"
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
// @Router       /printers/info/{id} [get]
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
// @Router       /printers/monitor/{id} [get]
func (h *PrinterHandler) ViewPrinterStatus(c echo.Context) wrapper.Response {
	idStr := c.Param("id")
	if idStr == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, erro := h.printerUC.ViewPrinterStatus(c.Request().Context(), uint(id))
	if erro != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}

// @Summary List All Clusters
// @Description List All Clusters
// @Tags clusters
// @Produce json
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.ListClusterResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /clusters/list [get]
func (h *PrinterHandler) ListCluster(c echo.Context) wrapper.Response {
	resp, err := h.printerUC.ListCluster(c.Request().Context())
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}

// @Summary Submit Print Job
// @Description Submit Print Job
// @Tags print jobs
// @Accept json
// @Produce json
// @Param request body dto.SubmitPrintJobRequestBody true "Submit Print Job Request Body"
// @Success      201  {object}  wrapper.SuccessResponse{data=dto.PrintJobResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /print-jobs/submit [post]
func (h *PrinterHandler) SubmitPrintJob(c echo.Context) wrapper.Response {
	req := &dto.SubmitPrintJobRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}
	resp, err := h.printerUC.SubmitPrintJob(c.Request().Context(), req)
	if err != nil {
		return wrapper.Response{Error: err}
	}
	return wrapper.Response{Data: resp, Status: http.StatusCreated}
}

// @Summary Cancel Print Job
// @Description Cancel Print Job
// @Tags print jobs
// @Produce json
// @Param id path string true "Print Job ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=nil}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /print-jobs/cancel/{id} [post]
func (h *PrinterHandler) CancelPrintJob(c echo.Context) wrapper.Response {
	id := c.Param("id")
	if id == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	err := h.printerUC.CancelPrintJob(c.Request().Context(), id)
	if err != nil {
		return wrapper.Response{Error: err}
	}
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}

// @Summary List Print Jobs
// @Description List Print Jobs of a Printer
// @Tags print jobs
// @Produce json
// @Param printerid path string true "Printer ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.ListPrintJobResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /print-jobs/list [get]
func (h *PrinterHandler) ListPrintJob(c echo.Context) wrapper.Response {
	printerIDStr := c.Param("printerid")
	if printerIDStr == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	printerID, err := strconv.ParseUint(printerIDStr, 10, 32)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, erro := h.printerUC.ListPrintJobs(c.Request().Context(), uint(printerID))
	if erro != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}

// @Summary View Print Job Status
// @Description View Print Job Status
// @Tags print jobs
// @Produce json
// @Param id path string true "Print Job ID"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.PrintJobResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /print-jobs/monitor/{id} [get]
func (h *PrinterHandler) ViewJobStatus(c echo.Context) wrapper.Response {
	id := c.Param("id")
	if id == "" {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	resp, err := h.printerUC.ViewJobStatus(c.Request().Context(), id)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: resp, Status: http.StatusOK}
}
