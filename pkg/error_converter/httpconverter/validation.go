package httpconverter

import (
	"errors"
	"fmt"
	"gameapp/pkg/richerror"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func RaiseValidationError(err error) *echo.HTTPError {
	var rErr richerror.RichError
	ok := errors.As(err, &rErr)
	if ok {
		statusCode := mapRichErrorCodeToHttpErrorCode(rErr.Kind)

		fmt.Println("\n--------------------log---------------------------")
		log.Printf("error in %s : %s", rErr.Operation, rErr.Meta)
		fmt.Println("--------------------log---------------------------\n")
		return echo.NewHTTPError(statusCode, rErr.ValidationErrors)
	}
	return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
		"error": err.Error(),
	})
}
