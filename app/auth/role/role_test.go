package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedFail    = false
	expectedSuccess = true
)

func TestRoleValid(t *testing.T) {

	userAuthority := Member.GenerateAuthority()
	assert.Equal(t, expectedSuccess, Member.Valid(userAuthority))
	assert.Equal(t, expectedFail, Admin.Valid(userAuthority))
}

func TestGenerateAutority(t *testing.T) {
	adminAuthority := Admin.GenerateAuthority()

	assert.Equal(t, expectedSuccess, adminAuthority[Member])
	assert.Equal(t, expectedSuccess, adminAuthority[Admin])

	userAuthority := Member.GenerateAuthority()

	assert.Equal(t, expectedSuccess, userAuthority[Member])
	assert.Equal(t, expectedFail, userAuthority[Admin])
}
