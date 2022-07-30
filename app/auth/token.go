package auth

import (
	"blog_api/app/auth/role"
	"blog_api/app/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	API_SECRET string
)

func init() {
	if !util.IsTest() {

		if err := godotenv.Load(util.GetRootPath() + "/.env"); err != nil {
			log.Println("[FATAL] fail to load .env file")
			log.Fatalln(err)
		}
		API_SECRET = os.Getenv("API_SECRET")
	} else {
		API_SECRET = "12345"
	}
}

func CreateToken(user_id uint32, authority role.Role) (string, error) {
	claims := GenerateMapClaim(user_id, authority)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func GenerateMapClaim(user_id uint32, authority role.Role) *jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authority"] = authority.GenerateAuthority()

	return &claims
}

func TokenValidByRole(r role.Role) func(r *http.Request) error {
	return func(r *http.Request) error {
		tokenString := ExtractToken(r)
		token, err := jwt.Parse(tokenString, keyFunc)
		if err != nil {
			return err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			pretty(claims)
			err = claims.Valid()
			if err != nil {
				log.Println("claims is not valid", err)
			}
		}
		return nil
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(API_SECRET), nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}

	bearerToken := r.Header.Get(echo.HeaderAuthorization)
	splitedBearerToken := strings.Split(bearerToken, " ")
	if len(splitedBearerToken) == 2 {
		return splitedBearerToken[1]
	}

	return ""
}

func ExtractID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	claims.Valid()
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

func ExtractClaims(c echo.Context) (jwt.MapClaims, error) {
	tokenString := ExtractToken(c.Request())
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		log.Println("[ERROR] token parse error:", err)
		return nil, err
	}

	err = token.Claims.Valid()
	if err != nil {
		log.Println("[ERROR] token validation error:", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims is not exist")
	}

	return claims, nil
}

func pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
