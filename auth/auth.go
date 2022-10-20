package auth

import (
	"System/System"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	// todo get the Token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	count := 0
	for _, i := range tokenString {
		if i == '.' {
			count++
		}
	}
	if count != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//todo Decode & validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("1o2m3a4r5z6i7z8o9"), nil
	})
	//todo Check the Exp
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//todo Continue
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)

}

//func Auth(c *gin.Context) {
//	Token := c.GetHeader("Authorization")
//	check, row := getUserByToken(Token)
//	if !check || row.Role != "Student" {
//		c.AbortWithStatus(http.StatusUnauthorized)
//	}
//	c.Next()
//}

func AuthAdmin(c *gin.Context) {
	// todo get the role header
	role := c.GetHeader("role")
	if role != "Admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
		System.BadRequest(c, "you can't access this")
	}
	c.Next()
}
