package middleware

import (
	"fmt"
	"go-jwt-auth/models"
	"go/token"
	"net/http"

	" github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context){
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorzation")

	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

//decode/validate it 

token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {


if claims, ok := token.Claims.(jwt.MapClaims); ok {
	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)

	}


//find the user with token sub
var user  models.User
initializers.DB.First(&user, claims["sub"])


if user.ID == 0{
	c.AbortWithStatus(http.StatusUnauthorized)
	
}

//attach to req
c.Set("user", user)

// continue 
c.Next()


} else {
	c.AbortWithStatus(http.StatusUnauthorized)
}