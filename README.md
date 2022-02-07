# auth-middleware
Standard API response of Microservices
## errorResponse inside the middleware
1. The error code and error map reuired in the middleware
```go
const (
	ErrAuthorizationTokenEmpty   int = 20003
	ErrAuthorizationTokenInvalid int = 20004
)

var ErrorsMap = map[int]response.ErrorStruct{
	ErrAuthorizationTokenEmpty:   {ErrorMsg: "Authorization token not provided", RespCode: http.StatusUnauthorized},
	ErrAuthorizationTokenInvalid: {ErrorMsg: "Authorization token invalid", RespCode: http.StatusUnauthorized},
}
```




## Functions inside the middleware
1. VerifyJwtToken() this function verifies token and also returns claims .
```go
VerifyJwtToken(c *gin.Context, jwtSecret string) (jwt.MapClaims, bool, int, error)
```

2. SetCorsForDev will set the necessary Header for Development 
```go
func SetCorsForDev(c *gin.Context) 
```

3. SetCorsForProd will set the necessary Header for Production environment 
```go
func SetCorsForProd(c *gin.Context) 
```

