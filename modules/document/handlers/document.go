package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/www-printf/wepress-core/modules/document/dto"
	"github.com/www-printf/wepress-core/modules/document/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type DocumentHandler struct {
	documentUC usecases.DocumentUsecase
}

func NewDocumentHandler(g *echo.Group, documentUC usecases.DocumentUsecase, middlewareMngr middlewares.MiddlewareManager) {
	h := &DocumentHandler{
		documentUC: documentUC,
	}

	api := g.Group("documents")
	api.Use(middlewareMngr.Auth())

	api.POST("/upload", wrapper.Wrap(h.SaveDocument)).Name = "document:save-document"
	api.POST("/request-upload", wrapper.Wrap(h.RequestUpload)).Name = "document:generate-signed-url-for-upload"
	api.GET("/download", wrapper.Wrap(h.DownloadDocumentList)).Name = "document:download-document-list"
	api.GET("/download/:id", wrapper.Wrap(h.DownloadDocument)).Name = "document:download-document"
}

// @Summary Request Upload Document
// @Description Generate Presigned URL For Upload Document
// @Tags documents
// @Accept json
// @Produce json
// @Param request body dto.UploadRequestBody true "Presigned URL Request Body"
// @Success      201  {object}  wrapper.SuccessResponse{data=dto.RequestUploadResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents/request-upload [post]
func (h *DocumentHandler) RequestUpload(c echo.Context) wrapper.Response {
	req := &dto.UploadRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	uid := c.Get("uid").(string)
	resp, err := h.documentUC.RequestUpload(c.Request().Context(), req, uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: resp, Status: http.StatusCreated}
}

// @Summary Save Document
// @Description Save Document
// @Tags documents
// @Accept json
// @Produce json
// @Param request body dto.UploadDocumentRequestBody true "Upload Document Request Body"
// @Success      201  {object}	wrapper.SuccessResponse{data=dto.UploadResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      403  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents/upload [post]
func (h *DocumentHandler) SaveDocument(c echo.Context) wrapper.Response {
	req := &dto.UploadDocumentRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	uid := c.Get("uid").(string)
	doc, err := h.documentUC.SaveDocument(c.Request().Context(), req, uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: doc, Status: http.StatusCreated}
}

// @Summary Download Document
// @Description Download Document
// @Tags documents
// @Produce json
// @Param id path string true "Document ID"
// @Success      200  {object}	wrapper.SuccessResponse{data=dto.DownloadDocumentResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      403  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents/download/{id} [get]
func (h *DocumentHandler) DownloadDocument(c echo.Context) wrapper.Response {
	id := c.Param("id")
	uid := c.Get("uid").(string)

	doc, err := h.documentUC.DownloadDocument(c.Request().Context(), id, uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: doc, Status: http.StatusOK}
}

// @Summary Download Document List
// @Description Download Document List
// @Tags documents
// @Produce json
// @Param page query string false "Page Number"
// @Param per_page query string false "Documents Per Page"
// @Success      200  {object}	wrapper.SuccessResponse{data=dto.DownloadDocumentsResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      403  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents/download [get]
func (h *DocumentHandler) DownloadDocumentList(c echo.Context) wrapper.Response {
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("per_page")
	uid := c.Get("uid").(string)

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = constants.DefaultPerPage
	}

	docs, erro := h.documentUC.DownloadDocumentList(c.Request().Context(), uid, page, perPage)
	if err != nil {
		return wrapper.Response{Error: erro}
	}

	return wrapper.Response{Data: docs, Status: http.StatusOK}
}
