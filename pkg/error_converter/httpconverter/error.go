package httpconverter

import (
	"errors"
	"fmt"
	"gameapp/pkg/richerror"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func mapRichErrorCodeToHttpErrorCode(kind richerror.KindError) int {
	switch kind {
	case richerror.NotFound:
		return http.StatusNotFound
	case richerror.Forbidden:
		return http.StatusForbidden
	case richerror.Invalid:
		return http.StatusBadRequest
	case richerror.Unexpected:
		return http.StatusInternalServerError
	default:
		log.Println("kind this error not implementation:", kind)
		return http.StatusServiceUnavailable
	}
}

func RaiseError(err error) *echo.HTTPError {
	var rErr richerror.RichError
	ok := errors.As(err, &rErr)
	if ok {
		sCode := mapRichErrorCodeToHttpErrorCode(rErr.Kind)
		fmt.Println("\n--------------------log---------------------------")
		log.Printf("error in %s : %s", rErr.Operation, rErr.Meta)
		fmt.Println("--------------------log---------------------------\n")

		if rErr.ValidationErrors != nil {
			return echo.NewHTTPError(sCode, rErr.ValidationErrors)
		}
		return echo.NewHTTPError(sCode, echo.Map{
			"error": rErr.Message,
		})
	}
	return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
		"error": err.Error(),
	})

}
