package design


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	service := UserService{}
	user, err := service.Login(10506000010, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &UserInfo{Name: "login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(10506000010, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &UserInfo{Name: "login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := UserService{}
	user, err := service.Register(10506000010, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &UserInfo{Name: "register"}, user)
}