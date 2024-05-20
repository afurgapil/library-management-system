package middleware

import (
	"net/http"

	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func EmployeeJWTMiddleware(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": "Missing or malformed JWT"})
    }

    claims, err := utils.ValidateEmployeeJWT(tokenString)
    if err != nil {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": err.Error()})
    }

    c.Locals("employeeID", claims.EmployeeID)
    return c.Next()
}

func StudentWTMiddleware(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": "Missing or malformed JWT"})
    }

    claims, err := utils.ValidateStudentJWT(tokenString)
    if err != nil {
        c.Status(http.StatusUnauthorized)
        return c.JSON(fiber.Map{"error": err.Error()})
    }

    c.Locals("studentID", claims.StudentID)
    return c.Next()
}