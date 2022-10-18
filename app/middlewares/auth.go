package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT      string
	ExpireDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
	}
}

// Generating Token
func (cj *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	listedToken, _ := token.SignedString([]byte(cj.SecretJWT))

	whitelist = append(whitelist, listedToken)
	return listedToken

}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)

	if isListed := CheckToken(user.Raw); !isListed {
		return nil
	}

	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func CheckToken(token string) bool {
	for _, listedToken := range whitelist {
		if listedToken == token {
			return true
		}
	}
	return false
}

func Logout(token string) bool {
	for i, listedToken := range whitelist {
		if listedToken == token {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
		}
	}
	return true
}
