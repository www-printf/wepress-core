package document_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/www-printf/wepress-core/modules/document/dto"
	"github.com/www-printf/wepress-core/modules/document/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/wrapper"
	tests "github.com/www-printf/wepress-core/tests/mocks"
)

type mockDocumentUsecase struct {
	mock.Mock
}

func (m *mockDocumentUsecase) RequestUpload(ctx context.Context, req *dto.UploadRequestBody, uid string) (*dto.RequestUploadResponseBody, error) {
	args := m.Called(ctx, req, uid)
	return args.Get(0).(*dto.RequestUploadResponseBody), args.Error(1)
}

func (m *mockDocumentUsecase) SaveDocument(ctx context.Context, req *dto.UploadDocumentRequestBody, uid string) (*dto.UploadResponseBody, error) {
	args := m.Called(ctx, req, uid)
	return args.Get(0).(*dto.UploadResponseBody), args.Error(1)
}

func (m *mockDocumentUsecase) DownloadDocument(ctx context.Context, id, uid string) (*dto.DownloadDocumentResponseBody, error) {
	args := m.Called(ctx, id, uid)
	return args.Get(0).(*dto.DownloadDocumentResponseBody), args.Error(1)
}

func (m *mockDocumentUsecase) DownloadDocumentList(ctx context.Context, uid string, page, perPage int) (*dto.DownloadDocumentsResponseBody, error) {
	args := m.Called(ctx, uid, page, perPage)
	return args.Get(0).(*dto.DownloadDocumentsResponseBody), args.Error(1)
}

func TestRequestUpload(t *testing.T) {
	e := echo.New()
	usecase := new(mockDocumentUsecase)
	middlewareMngr := new(tests.MockMiddlewareManager)
	h := &handlers.DocumentHandler{
		documentUC: usecase,
	}

	reqBody := dto.UploadRequestBody{FileName: "test.pdf", FileType: "application/pdf"}
	respBody := dto.RequestUploadResponseBody{PresignedURL: "http://example.com/upload"}
	uid := "1234"

	usecase.On("RequestUpload", mock.Anything, &reqBody, uid).Return(&respBody, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/documents/request-upload", httptest.NewBody(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req = req.WithContext(context.WithValue(req.Context(), "uid", uid))
	
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.RequestUpload(c).Send(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response wrapper.SuccessResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, respBody.PresignedURL, response.Data.(dto.RequestUploadResponseBody).PresignedURL)
	}
}

func TestSaveDocument(t *testing.T) {
	e := echo.New()
	usecase := new(mockDocumentUsecase)
	middlewareMngr := new(tests.MockMiddlewareManager)
	h := &handlers.DocumentHandler{
		documentUC: usecase,
	}

	reqBody := dto.UploadDocumentRequestBody{DocumentID: "doc123"}
	respBody := dto.UploadResponseBody{DocumentID: "doc123", Status: "uploaded"}
	uid := "1234"

	usecase.On("SaveDocument", mock.Anything, &reqBody, uid).Return(&respBody, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/documents/upload", httptest.NewBody(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req = req.WithContext(context.WithValue(req.Context(), "uid", uid))
	
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.SaveDocument(c).Send(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response wrapper.SuccessResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, respBody.DocumentID, response.Data.(dto.UploadResponseBody).DocumentID)
	}
}

func TestDownloadDocument(t *testing.T) {
	e := echo.New()
	usecase := new(mockDocumentUsecase)
	middlewareMngr := new(tests.MockMiddlewareManager)
	h := &handlers.DocumentHandler{
		documentUC: usecase,
	}

	id := "doc123"
	uid := "1234"
	respBody := dto.DownloadDocumentResponseBody{DocumentURL: "http://example.com/download"}

	usecase.On("DownloadDocument", mock.Anything, id, uid).Return(&respBody, nil)
	req := httptest.NewRequest(http.MethodGet, "/documents/download/"+id, nil)
	req = req.WithContext(context.WithValue(req.Context(), "uid", uid))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/documents/download/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, h.DownloadDocument(c).Send(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response wrapper.SuccessResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, respBody.DocumentURL, response.Data.(dto.DownloadDocumentResponseBody).DocumentURL)
	}
}

func TestDownloadDocumentList(t *testing.T) {
	e := echo.New()
	usecase := new(mockDocumentUsecase)
	middlewareMngr := new(tests.MockMiddlewareManager)
	h := &handlers.DocumentHandler{
		documentUC: usecase,
	}

	uid := "1234"
	page := 1
	perPage := 10
	respBody := dto.DownloadDocumentsResponseBody{Documents: []dto.DocumentItem{}}

	usecase.On("DownloadDocumentList", mock.Anything, uid, page, perPage).Return(&respBody, nil)
	req := httptest.NewRequest(http.MethodGet, "/documents/download?page=1&per_page=10", nil)
	req = req.WithContext(context.WithValue(req.Context(), "uid", uid))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.DownloadDocumentList(c).Send(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response wrapper.SuccessResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.NotNil(t, response.Data)
	}
}
