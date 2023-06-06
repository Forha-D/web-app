package middleware

import (
	token "github.com/Forha-D/ecommerce/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"No authorization header provided"})
             c.Abort()
			return
		}
      claims , err := token.ValidateToken(ClientToken)
      if err != "" {
      	c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		  return
	  }
	  c.Set( "email",claims.Email)
      c.Set("uid" , claims.Uid)
      c.Next()

	}
}