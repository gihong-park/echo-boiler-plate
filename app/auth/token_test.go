package auth

import (
	"blog_api/app/auth/role"
	"blog_api/app/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	tokenPrefix = "Bearer "
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(1, role.Member)

	t.Logf("token: %v\n", token)
	if err != nil {
		t.Logf("[ERROR] token signing error occured: %v", err)
		t.Error(err)
	}

	assert.NotEmpty(t, token)
}

func TestGenerateMapClaim(t *testing.T) {
	var user_id uint32 = 1
	claims := *(GenerateMapClaim(user_id, role.Member))

	assert.Equal(t, user_id, claims["user_id"])
	assert.Equal(t, role.Member.GenerateAuthority(), claims["authority"])
	assert.Equal(t, true, claims["authorized"])
	assert.IsType(t, int64(1), claims["exp"])
}

func TestExtractToken(t *testing.T) {
	expectedToken := "abced"
	req := httptest.NewRequest(http.MethodPost, "/", nil)

	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+expectedToken)

	assert.Equal(t, "Bearer "+expectedToken, req.Header.Get("Authorization"))

	token := ExtractToken(req)
	assert.Equal(t, expectedToken, token)
}

func TestTokenValidByRoleError(t *testing.T) {
	token := "abcde"

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, tokenPrefix+token)

	err := TokenValidByRole(role.Member)(req)
	t.Log(err)
	assert.NotNil(t, err)
}

func TestTokenValidByRole(t *testing.T) {
	token, _ := CreateToken(1, role.Member)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, tokenPrefix+token)
	err := TokenValidByRole(role.Member)(req)

	t.Log("[INFO] token valid error:", err)

	assert.Nil(t, err)
	err = TokenValidByRole(role.Admin)(req)
	assert.NotNil(t, err)
}

func TestExtractClaim(t *testing.T) {
	e := util.NewServer()
	token, _ := CreateToken(1, role.Member)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, tokenPrefix+token)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	claims, err := ExtractClaims(c)
	if err != nil {
		t.Log(err)
	}

	var user_id float64 = 1
	authority := make(map[string]interface{})
	authority["0"] = true

	assert.Equal(t, user_id, claims["user_id"])
	assert.Equal(t, authority, claims["authority"])
	assert.Equal(t, true, claims["authorized"])
	assert.IsType(t, float64(1), claims["exp"])
}

func TestExtractID(t *testing.T) {
	req := tokenRequest()

	id, err := ExtractID(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	assert.Equal(t, uint32(1), id)
}

func tokenRequest() *http.Request {
	token, _ := CreateToken(1, role.Member)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, tokenPrefix+token)

	return req
}

// func TestTokenValidByRole(t *testing.T) {
// 	assert.Equal(t, true, TokenValidByRole(role.Member)(tokenRequest()))
// 	assert.Equal(t, false, TokenValidByRole(role.Member)(tokenRequest()))
// }
