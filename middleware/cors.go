package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, X-Requested-With, Origin, Content-Type, Content-Length, Accept-Encoding, X-XSRF-Token, X-CSRF-Token, Authorization")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
		//if c.Request.Method != "OPTIONS" {
		//	c.Next()
		//} else {
		//	c.Header("Access-Control-Allow-Origin", "*")
		//	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		//	//c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		//	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-CSRF-Token")
		//	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		//	c.Header("Content-Type", "application/json")
		//	c.AbortWithStatus(http.StatusOK)
		//}
	}
}
