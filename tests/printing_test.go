package printer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/usecases/mocks"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

func setup() (*echo.Echo, *PrinterHandler, *mocks.PrinterUsecase) {
	e := echo.New()
	mockUsecase := new(mocks.PrinterUsecase)
	handler := &PrinterHandler{
		printerUC: mockUsecase,
	}
	return e, handler, mockUsecase
}

func TestAddPrinter(t *testing.T) {
	e, handler, mockUsecase := setup()

	reqBody := dto.AddPrinterRequestBody{
		Name: "Printer1",
		Type: "Laser",
	}
	reqJSON, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/printers/add", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	mockUsecase.On("AddPrinter", c.Request().Context(), &reqBody).Return(&dto.PrinterResponseBody{
		ID:   1,
		Name: reqBody.Name,
		Type: reqBody.Type,
	}, nil)

	if assert.NoError(t, handler.AddPrinter(c).Handle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var resp wrapper.SuccessResponse
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp)) {
			assert.Equal(t, "Printer1", resp.Data.(map[string]interface{})["name"])
		}
	}
}

func TestListPrinter(t *testing.T) {
	e, handler, mockUsecase := setup()

	clusterID := 1
	req := httptest.NewRequest(http.MethodGet, "/printers/list?cluster_id="+strconv.Itoa(clusterID), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	mockUsecase.On("ListPrinter", c.Request().Context(), uint(clusterID)).Return(&dto.ListPrinterResponseBody{
		Printers: []dto.PrinterResponseBody{
			{ID: 1, Name: "Printer1", Type: "Laser"},
		},
	}, nil)

	if assert.NoError(t, handler.ListPrinter(c).Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp wrapper.SuccessResponse
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp)) {
			assert.Len(t, resp.Data.(map[string]interface{})["printers"], 1)
		}
	}
}

func TestGetPrinter(t *testing.T) {
	e, handler, mockUsecase := setup()

	id := 1
	req := httptest.NewRequest(http.MethodGet, "/printers/info/"+strconv.Itoa(id), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	mockUsecase.On("GetPrinter", c.Request().Context(), uint(id)).Return(&dto.PrinterResponseBody{
		ID:   uint(id),
		Name: "Printer1",
		Type: "Laser",
	}, nil)

	if assert.NoError(t, handler.GetPrinter(c).Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp wrapper.SuccessResponse
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp)) {
			assert.Equal(t, "Printer1", resp.Data.(map[string]interface{})["name"])
		}
	}
}

func TestCancelPrintJob(t *testing.T) {
	e, handler, mockUsecase := setup()

	jobID := "123"
	req := httptest.NewRequest(http.MethodPost, "/print-jobs/cancel/"+jobID, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(jobID)

	mockUsecase.On("CancelPrintJob", c.Request().Context(), jobID).Return(nil)

	if assert.NoError(t, handler.CancelPrintJob(c).Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp wrapper.SuccessResponse
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp)) {
			assert.Nil(t, resp.Data)
		}
	}
}
