package middleware

import (
	"net/http"

	response "github.com/Zbyteio/api-response-lib"
)

const (
	ErrAuthorizationTokenEmpty   int = 20003
	ErrAuthorizationTokenInvalid int = 20004
)

var ErrorsMap = map[int]response.ErrorStruct{
	ErrAuthorizationTokenEmpty:   {ErrorMsg: "Authorization token not provided", RespCode: http.StatusUnauthorized},
	ErrAuthorizationTokenInvalid: {ErrorMsg: "Authorization token invalid", RespCode: http.StatusUnauthorized},
}
