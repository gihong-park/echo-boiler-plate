package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedFail    = false
	expectedSuccess = true
)

func TestRoleString(t *testing.T) {
	assert.Equal(t, "Member", Member.String())
	assert.Equal(t, "Admin", Admin.String())
}

func TestRoleValid(t *testing.T) {

	userAuthority := Member.GenerateAuthority()
	userAuthInterface := make(map[string]interface{})
	for k, v := range userAuthority {
		userAuthInterface[k] = v
	}
	assert.Equal(t, expectedSuccess, Member.Valid(userAuthInterface))
	assert.Equal(t, expectedFail, Admin.Valid(userAuthInterface))
}

func TestGenerateAutority(t *testing.T) {
	adminAuthority := Admin.GenerateAuthority()

	assert.Equal(t, expectedSuccess, adminAuthority[Member.String()])
	assert.Equal(t, expectedSuccess, adminAuthority[Admin.String()])

	userAuthority := Member.GenerateAuthority()

	assert.Equal(t, expectedSuccess, userAuthority[Member.String()])
	assert.Equal(t, expectedFail, userAuthority[Admin.String()])
}
