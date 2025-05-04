package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	_ = godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET must be set")
	}
	jwtSecret = []byte(secret)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			return
		}

		tokenString := parts[1]

		// Parse e validação do token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Verifica o método de assinatura
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		// Tratamento de erros
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "malformed token"})
			case errors.Is(err, jwt.ErrTokenExpired):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			case errors.Is(err, jwt.ErrTokenNotValidYet):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not yet valid"})
			case errors.Is(err, jwt.ErrSignatureInvalid):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			}
			return
		}

		// Verificação das claims
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		// Extrai o user_id das claims
		rawUserID, exists := claims["user_id"]
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id claim missing"})
			return
		}

		// Converte o user_id para uint
		userID, ok := rawUserID.(float64) // JSON numbers são float64 por padrão
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id format"})
			return
		}

		// Armazena o user_id no contexto
		c.Set("user_id", uint(userID))
		c.Next()
	}
}
