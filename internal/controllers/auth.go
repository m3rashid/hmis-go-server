package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/m3rashid-org/hmis-go-server/internal/models"
	"github.com/m3rashid-org/hmis-go-server/internal/params"
	"github.com/m3rashid-org/hmis-go-server/internal/services"
	"github.com/m3rashid-org/hmis-go-server/internal/utils"
)

func RefreshToken(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	claims := &utils.Claims{}
	token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET_KEY"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return c.SendStatus(200)
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	return c.SendStatus(200)
}

func CurrentUser(c *fiber.Ctx) error {

	return c.SendString("Hello, World!")
}

func AuthPing(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Login(c *fiber.Ctx) error {
	var Login params.LoginRequest
	c.BodyParser(Login)

	user, err := utils.PrismaClient.User.FindUnique(
		models.User.Email.Equals(Login.Email),
	).With(
		models.User.Profile.Fetch(),
	).With(
		models.User.Role.Fetch().With(
			models.Role.Permissions.Fetch(),
		),
	).Exec(context.Background())

	if err != nil {
		c.SendStatus(404)
	}

	fmt.Println(user)

	matched := services.CheckPasswordHash(user.Password, Login.Password)
	if !matched {
		c.SendStatus(401)
	}

	rawPermissions := user.Role().Permissions()
	permissions := rawPermissions
	fmt.Println(permissions)

	// create json web token
	// tokenString, expirationTime, err := utils.GenerateJwtToken(user.ID, user.Email, user.)

	// TODO: remove these error escapes
	tokenString := ""
	expirationTime := time.Now()

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return c.Status(200).JSON(fiber.Map{
		"token":       tokenString,
		"user":        user,
		"permissions": rawPermissions,
	})
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
