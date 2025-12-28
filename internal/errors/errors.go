package errors

import (
	"fmt"
	"net/http"
)

// AppError 应用错误类型
type AppError struct {
	Code       string // 错误代码
	Message    string // 错误消息
	HTTPStatus int    // HTTP状态码
	Err        error  // 原始错误（可选）
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 实现 errors.Unwrap 接口
func (e *AppError) Unwrap() error {
	return e.Err
}

// 预定义的通用错误类型（只保留最基础的）
var (
	ErrNotFound = &AppError{
		Code:       "NOT_FOUND",
		Message:    "resource not found",
		HTTPStatus: http.StatusNotFound,
	}

	ErrUnauthorized = &AppError{
		Code:       "UNAUTHORIZED",
		Message:    "unauthorized",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrForbidden = &AppError{
		Code:       "FORBIDDEN",
		Message:    "access denied",
		HTTPStatus: http.StatusForbidden,
	}

	ErrBadRequest = &AppError{
		Code:       "BAD_REQUEST",
		Message:    "bad request",
		HTTPStatus: http.StatusBadRequest,
	}
)

// NewAppError 创建新的应用错误
func NewAppError(code, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// WrapAppError 包装错误为应用错误
func WrapAppError(err error, appErr *AppError) *AppError {
	return &AppError{
		Code:       appErr.Code,
		Message:    appErr.Message,
		HTTPStatus: appErr.HTTPStatus,
		Err:        err,
	}
}

// WithMessage 为错误添加自定义消息
func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    message,
		HTTPStatus: e.HTTPStatus,
		Err:        e.Err,
	}
}

// IsAppError 检查错误是否为应用错误
func IsAppError(err error) (*AppError, bool) {
	if appErr, ok := err.(*AppError); ok {
		return appErr, true
	}
	return nil, false
}

// 辅助函数：快速创建常见错误
// NotFound 创建404错误
func NotFound(message string) *AppError {
	if message == "" {
		return ErrNotFound
	}
	return ErrNotFound.WithMessage(message)
}

// Forbidden 创建403错误
func Forbidden(message string) *AppError {
	if message == "" {
		return ErrForbidden
	}
	return ErrForbidden.WithMessage(message)
}

// BadRequest 创建400错误
func BadRequest(message string) *AppError {
	if message == "" {
		return ErrBadRequest
	}
	return ErrBadRequest.WithMessage(message)
}
