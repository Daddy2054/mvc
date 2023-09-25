package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user id 0")
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusNotFound,err.StatusCode)
	
	if user != nil {
		t.Error("we were not expecting a user id 0")
	}

	if err == nil {
		t.Error("we were expecting a user id 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when a user not found")
	}

}
