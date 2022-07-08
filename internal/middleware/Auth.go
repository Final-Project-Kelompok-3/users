package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Final-Project-Kelompok-3/users/internal/model"
	res "github.com/Final-Project-Kelompok-3/users/pkg/util/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	jwtKey = os.Getenv("JWT_KEY")
)

func AuthJWT(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		return next(c)
	}
}

func CreateToken(u model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = u.Model.ID
	claims["roleId"] = u.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}

func AuthApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if (len(c.Request().Header["Api-Key"]) > 0) {
			if (c.Request().Header["Api-Key"][0] == os.Getenv("API_KEY")) {
				return next(c)
			}
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}