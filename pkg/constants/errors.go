package constants

import (
	"net/http"

	"github.com/www-printf/wepress-core/pkg/errors"
)

var (
	HTTPNotFound     = &errors.HTTPError{Status: http.StatusNotFound, Message: "Resource Not Found"}
	HTTPInternal     = &errors.HTTPError{Status: http.StatusInternalServerError, Message: "Internal Server Error"}
	HTTPUnauthorized = &errors.HTTPError{Status: http.StatusUnauthorized, Message: "Unauthorized"}
	HTTPForbidden    = &errors.HTTPError{Status: http.StatusForbidden, Message: "Forbidden"}
	HTTPBadRequest   = &errors.HTTPError{Status: http.StatusBadRequest, Message: "Bad Request"}
)
