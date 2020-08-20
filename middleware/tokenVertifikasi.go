package middleware

import (
	"ecommerce/model"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TokenVerifikasiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// autDetail, err := ExtractTokenAuth(c)
		// if err != nil {
		// 	utils.HandleError(c, http.StatusUnauthorized, err.Error())
		// 	c.Abort()
		// 	return
		// }
		// fmt.Printf("User acces => user_id: %v, role: %v, pass: %v \n", autDetail.ID, autDetail.Role, autDetail.Password)

		err := TokenValid(c)
		if err != nil {
			utils.HandleError(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenValid(c *gin.Context) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("security.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(c *gin.Context) string {
	keys := c.Request.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	c.Writer.Header()
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1]
	} else {
		return ""
	}
}

func ExtractTokenAuth(c *gin.Context) (*model.AuthDetail, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		pass, ok := claims["pass"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		authDetail := model.AuthDetail{
			ID:   int(userId),
			Password: pass,
		}
		return &authDetail, nil
	}
	return nil, err
}
