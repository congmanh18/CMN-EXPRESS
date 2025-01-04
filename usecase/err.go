package usecase

import (
	"express_be/core/log"

	"go.uber.org/zap"
)

type Error struct {
	Code    int    // Mã lỗi
	Message string // Thông báo lỗi
	Err     error  // Lỗi gốc (nếu có)
}

// Hàm tạo lỗi với logging
func NewError(code int, message string, err error) *Error {
	e := &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
	log.Error("Usecase Error: ",
		zap.Int("code", e.Code),
		zap.String("message", e.Message),
		zap.Error(e.Err),
	)
	return e
}

// Triển khai interface error
func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}
