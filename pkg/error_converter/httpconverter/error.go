package httpconverter

import (
	"errors"
	"fmt"
	"gameapp/pkg/richerror"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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
	var richEr richerror.RichError
	ok := errors.As(err, &richEr)
	if ok {
		statCode := mapRichErrorCodeToHttpErrorCode(richEr.Kind)
		fmt.Println("\n--------------------log---------------------------")
		log.Printf("error in %s : %s", richEr.Operation, richEr.Meta)
		fmt.Println("--------------------log---------------------------\n")
		return echo.NewHTTPError(statCode, echo.Map{
			"error": richEr.Message,
		})
	}
	return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
		"error": err.Error(),
	})

}
