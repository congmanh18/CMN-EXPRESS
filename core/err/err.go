package err

import (
	"express_be/core/log"
	"fmt"

	"go.uber.org/zap"
)

var errorPrefixMap = map[int]string{
	100: "WP",
	200: "AU",
	300: "VE",
	400: "DB",
	500: "SE",
}

type Err struct {
	Message     string
	Code        int
	Debug       error
	Description string
}

func newError(code int, message string, debug error, description string) *Err {
	prefix := getPrefix(code)
	return &Err{
		Message:     fmt.Sprintf("%s%d", prefix, code),
		Code:        code,
		Debug:       debug,
		Description: fmt.Sprintf("%s: %s", message, description),
	}
}

func getPrefix(code int) string {
	base := (code / 100) * 100
	if prefix, ok := errorPrefixMap[base]; ok {
		return prefix
	}
	return "UN" // Default prefix for unknown errors
}

func (e *Err) Error() string {
	if e.Debug != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.Debug.Error())
	}
	return e.Message
}

func NewError(code int, message string, err error) *Err {
	e := &Err{
		Code:    code,
		Message: message,
		Debug:   err,
	}
	log.Error("Usecase Error: ",
		zap.Int("code", e.Code),
		zap.String("message", e.Message),
		zap.Error(e.Debug),
	)
	return e
}
