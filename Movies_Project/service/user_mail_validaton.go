package service

import (
	"Movies_Project/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
)

func isValidEmail(email string) bool {
	// Define a regular expression pattern for email validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailPattern)

	// Match the email against the regular expression
	return re.MatchString(email)
}

func ValidateEmail(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if !isValidEmail(user.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email address")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Email is valid.",
	})
}
