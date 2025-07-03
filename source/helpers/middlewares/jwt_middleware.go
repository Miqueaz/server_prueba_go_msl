package middleware

import (
	token "main/security/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token desde el encabezado Authorization
		t := c.GetHeader("Authorization")

		// Validar el token usando la función de tu paquete token
		_, err := token.ValidToken(t)
		if err != nil {
			// Si el token no es válido, respondemos con un error Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort() // Interrumpir el procesamiento de la solicitud
			return
		}

		// Si el token es válido, continuamos con la siguiente función del middleware
		c.Next()
	}
}
