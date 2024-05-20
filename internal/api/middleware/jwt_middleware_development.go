package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DevelopmentEmployeeTokenMiddleware(c *fiber.Ctx) error {
    testTokenString := "YOUR_DEVELOPMENT_TOKEN" 
    
    tokenString := c.Get("Authorization")
    if tokenString != testTokenString {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": "Invalid test token"})
    }
    
    return c.Next()
}

func DevelopmentStudentTokenMiddleware(c *fiber.Ctx) error {
    testTokenString := "YOUR_DEVELOPMENT_TOKEN" 
    
    tokenString := c.Get("Authorization")
    if tokenString != testTokenString {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": "Invalid test token"})
    }
    
    return c.Next()
}
