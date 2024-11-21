package handlers

import (
	"net/http"

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

	api.POST("/", wrapper.Wrap(h.SaveDocument)).Name = "document:save-document"
	api.POST("/generate-presigned-url", wrapper.Wrap(h.GenerateURL)).Name = "document:generate-signed-url"
}

// @Summary Generate Presigned URL
// @Description Generate Presigned URL
// @Tags documents
// @Accept json
// @Produce json
// @Param request body dto.PresignedURLRequestBody true "Presigned URL Request Body"
// @Success      201  {object}  wrapper.SuccessResponse{data=dto.PresignedURLResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents/generate-presigned-url [post]
func (h *DocumentHandler) GenerateURL(c echo.Context) wrapper.Response {
	req := &dto.PresignedURLRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	uid := c.Get("uid").(string)
	url, err := h.documentUC.GeneratePresignedURL(c.Request().Context(), req, uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Data: url, Status: http.StatusCreated}
}

// @Summary Save Document
// @Description Save Document
// @Tags documents
// @Accept json
// @Produce json
// @Param request body dto.UploadDocumentRequestBody true "Upload Document Request Body"
// @Success      201  {object}	wrapper.SuccessResponse{data=nil}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /documents [post]
func (h *DocumentHandler) SaveDocument(c echo.Context) wrapper.Response {
	req := &dto.UploadDocumentRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	uid := c.Get("uid").(string)
	err := h.documentUC.SaveDocument(c.Request().Context(), req, uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	return wrapper.Response{Status: http.StatusCreated}
}
