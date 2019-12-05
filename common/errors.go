package common

type (
    HttpAuthError struct {
        Code    int
        Message string
    }
)

func NewHttpAuthError() error {
    return HttpAuthError{Code: HTTP_AUTH_ERROR, Message: HTTP_AUTH_ERROR_MSG}
}

func (hae HttpAuthError) Error() string {
    return hae.Message
}
