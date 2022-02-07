package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

func VerifyJwtToken(c *gin.Context, jwtSecret string) (jwt.MapClaims, bool, int, error) {
	auth_token := c.Request.Header["Authorization"]

	// no auth token error
	if len(auth_token) == 0 {
		return nil, false, ErrAuthorizationTokenEmpty, ErrorsMap.ErrAuthorizationTokenEmpty
	}

	jwttoken := strings.Split(auth_token[0], " ")[1]

	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(jwttoken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, false, ErrAuthorizationTokenInvalid, ErrorsMap.ErrAuthorizationTokenInvalid
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, false, ErrAuthorizationTokenInvalid, ErrorsMap.ErrAuthorizationTokenInvalid
	}
	if !tkn.Valid {
		return nil, false, ErrAuthorizationTokenInvalid, ErrorsMap.ErrAuthorizationTokenInvalid
	}
	return claims, true, 0, nil
}

func SetCorsForDev(c *gin.Context) {
	// set cors access
	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.WriteHeader(200)
		return
	}
	log.Printf(c.Request.Method, c.Request.URL, "Allowed")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	return
}

func SetCorsForProd(c *gin.Context) {
	// set cors access
	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "app.zbyte.io")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "app.zbyte.io")
		c.Writer.WriteHeader(200)
		return
	}
	log.Printf(c.Request.Method, c.Request.URL, "Allowed")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "app.zbyte.io")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "app.zbyte.io")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	return
}
