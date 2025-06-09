package exception

import "fmt"

// CustomError định nghĩa lỗi có mã và message
type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Hàm khởi tạo lỗi mới
func NewCustomError(code int, message string) error {
    return &CustomError{
        Code:    code,
        Message: message,
    }
}
