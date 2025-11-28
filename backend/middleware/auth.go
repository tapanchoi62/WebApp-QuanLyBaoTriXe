package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra thuật toán
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.JwtSecret, nil
		})

		if err != nil {
			// Kiểm tra lỗi token hết hạn
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				c.Abort()
				return
			}
			// Lỗi token khác
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Kiểm tra claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check thủ công exp (nếu muốn chắc chắn)
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
					c.Abort()
					return
				}
			}

			// Lưu thông tin user vào context
			if username, ok := claims["username"].(string); ok {
				c.Set("username", username)
			}
			if role, ok := claims["role"].(string); ok {
				c.Set("role", role)
			}

			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
	}
}
