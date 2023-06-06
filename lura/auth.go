package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func init() {
	publicKey = []byte("your-secret-key")
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("authMiddleware")
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not found Authorization"})

			ctx.Abort() // stop next handler
			return
		}
		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, getKeyFunc)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})

			ctx.Abort() // stop next handler
			return
		}
		ctx.Next()
	}
}

func getKeyFunc(token *jwt.Token) (interface{}, error) {
	return publicKey, nil
}

var (
	publicKey interface{}
)
