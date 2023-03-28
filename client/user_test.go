package client

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData struct {
	client    *Client
	container testContainer
}

func TestMain(m *testing.M) {

	testData.client, testData.container, _ = runHeadscale()

	code := m.Run()

	testData.container.Terminate()

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {

	userName := "foo"
	userStatus, _, err := testData.client.CreateUser(testData.container.Context, userName)
	assert.NoError(t, err)
	expectedUserStatus := []UserStatus{
		UserCreated,
	}
	assert.Contains(t, expectedUserStatus, userStatus)

	userStatus, _, err = testData.client.CreateUser(testData.container.Context, userName)
	assert.NoError(t, err)
	expectedUserStatus = []UserStatus{
		UserExists,
	}
	assert.Contains(t, expectedUserStatus, userStatus)
}

func TestDeleteUser(t *testing.T) {

	userName := "bar"
	userStatus, _, _ := testData.client.CreateUser(testData.container.Context, userName)
	expectedUserStatus := []UserStatus{
		UserCreated,
	}
	assert.Contains(t, expectedUserStatus, userStatus)
	userStatus, err := testData.client.DeleteUser(testData.container.Context, userName)
	assert.Nil(t, err)
	expectedUserStatus = []UserStatus{
		UserDeleted,
	}
	assert.Contains(t, expectedUserStatus, userStatus)

	userStatus, err = testData.client.DeleteUser(testData.container.Context, userName)
	assert.ErrorIs(t, err, ErrUserNotFound)
	expectedUserStatus = []UserStatus{
		UserUnknown,
	}
	assert.Contains(t, expectedUserStatus, userStatus)
}

func TestUnauthorized(t *testing.T) {

	testData.client.APIKey = "wrongkey"
	userName := "bar"
	_, _, err := testData.client.CreateUser(context.Background(), userName)
	assert.ErrorIs(t, err, ErrUnauthorized)
}
