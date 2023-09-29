package services

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetUserNotFoundInDatabase(t *testing.T){
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

}