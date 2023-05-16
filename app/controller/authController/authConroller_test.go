package authController

import (
	"blog_api/app/auth/role"
	"blog_api/app/controller/dto"
	"blog_api/app/db"
	"blog_api/app/model"
	"blog_api/app/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	e        *echo.Echo
	authCont AuthController
	j        int
)

func init() {
	j = 1
	e = util.NewServer()
	database = db.GetDB("sqlite")
	database.Migrator().DropTable(&model.User{})
	database.AutoMigrate(&model.User{})
	authCont = InitAuthController(database)
}
func BenchmarkSignUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		item := dto.UserRequest{
			Name:     "test",
			Password: "password",
			Email:    "test" + strconv.Itoa(j) + "@example.com",
		}
		j += 1
		itemJson, _ := json.Marshal(item)

		req := httptest.NewRequest(http.MethodPost, "/signUp", bytes.NewReader(itemJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/auth")

		err := authCont.SignUp(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func TestSignUp(t *testing.T) {
	item := dto.UserRequest{
		Name:     "test",
		Password: "password",
		Email:    "test@example.com",
	}
	itemJson, _ := json.Marshal(item)

	req := httptest.NewRequest(http.MethodPost, "/signUp", bytes.NewReader(itemJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/auth")

	err := authCont.SignUp(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("response body: %v", rec.Body.String())
	assert.Equal(t, http.StatusCreated, rec.Code)
	user := new(dto.UserResponse)
	json.Unmarshal(rec.Body.Bytes(), user)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, role.Member.String(), user.Role)
}

func TestSignIn(t *testing.T) {
	item := dto.SignIn{
		Email:    "test@example.com",
		Password: "password",
	}
	itemJson, _ := json.Marshal(item)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(itemJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/auth")

	err := authCont.SignIn(c)
	if err != nil {
		assert.Nil(t, err, err.Error())
	}

	t.Logf("response body: %v", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)

}
