package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// AuthorizeJWT validates the token user given, return 401(Unauthorized) if not valid
func AuthorizeJWT(jwtService service.JWTService, userService service.UserService, redisService service.RedisService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		//Melakukan pengecekan pada header request yang dikirimkan oleh user
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		//Jika ada, cek header bearer token ke dalam service ValidateToken
		token, err := jwtService.ValidateToken(authHeader)
		//jika valid, melakukan mapping claim
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["user_id"]
		if token.Valid {
			//jika tokennya valid
			checkRedis := redisService.CheckValueKey(ID.(string), authHeader)
			//jika key dan value terdaftar
			if checkRedis {
				us, _ := userService.GetProfile(ID.(string))
				ctx.Set("UserLog", us)
			} else {
				response := helper.BuildErrorResponse("Token Invalid", "Key Or Value Not Found", nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		} else {
			//jika token sudah tidak valid
			redisService.ClearToken(ID.(string))
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}

func CheckRole(jwtService service.JWTService, userServ service.UserService, redisServ service.RedisService, roles string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		//Melakukan pengecekan pada header request yang dikirimkan oleh user
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		//Jika ada, cek header bearer token ke dalam service ValidateToken
		token, errTok := jwtService.ValidateToken(authHeader)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["user_id"]
		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			//ID := claims["ID"]
			checkRedis := redisServ.CheckValueKey(ID.(string), authHeader)
			if checkRedis {
				role := claims["role"]
				if role == roles {
					us, _ := userServ.GetProfile(ID.(string))
					context.Set("UserLog", us)
				} else {
					resp := helper.BuildErrorResponse("Role Invalid", "Role no have permission", nil)
					context.AbortWithStatusJSON(http.StatusForbidden, resp)
					return
				}
			} else {
				//redisServ.ClearToken(ID.(string))
				resp := helper.BuildErrorResponse("Token Invalid", "Key Or Value Not Found", nil)
				context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
				return
			}
		} else {
			redisServ.ClearToken(ID.(string))
			resp := helper.BuildErrorResponse("Token Invalid", errTok.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
	}
}
