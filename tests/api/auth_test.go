package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
	factories "github.com/nugrohosam/gosampleapi/tests/factories"
	utilities "github.com/nugrohosam/gosampleapi/tests/utilities"
	viper "github.com/spf13/viper"
	assert "github.com/stretchr/testify/assert"
)

var user userRepo.User

// TestRun ...
func TestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)

	user = factories.CreateUser()

	t.Log("Test Positive")
	t.Log("=======>>>> <<<<======")
	testAuthRegister(t)
	testAuthLogin(t)

	t.Log("Test Negative")
	t.Log("=======>>>> <<<<======")
	t.Log("Process...")
}

func testAuthRegister(t *testing.T) {
	url := viper.GetString("app.url")
	port := viper.GetString("app.port")

	data, err := json.Marshal(map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	},
	)

	if err != nil {
		t.Error(err.Error())
	}

	reader := bytes.NewBuffer(data)
	endpoint := "http://" + url + ":" + port + "/v1/auth/register"

	resp, err := http.Post(endpoint, "application/json", reader)
	if err != nil {
		t.Error(err.Error())
	} else {
		assert.Equal(t, 200, resp.StatusCode)
	}
}

func testAuthLogin(t *testing.T) {
	url := viper.GetString("app.url")
	port := viper.GetString("app.port")

	endpoint := "http://" + url + ":" + port + "/v1/auth/login"

	t.Log("Test Positive Login via email")
	loginWithEmail(t, endpoint)
	t.Log("Test Positive Login via username")
	loginWithUsername(t, endpoint)
}

func loginWithEmail(t *testing.T, endpoint string) {
	data, err := json.Marshal(map[string]interface{}{
		"emailOrUsername": user.Email,
		"password":        user.Password,
	},
	)

	if err != nil {
		t.Error(err.Error())
	}

	reader := bytes.NewBuffer(data)
	resp, _ := http.Post(endpoint, "application/json", reader)
	assert.Equal(t, 200, resp.StatusCode)
}

func loginWithUsername(t *testing.T, endpoint string) {
	data, err := json.Marshal(map[string]interface{}{
		"emailOrUsername": user.Username,
		"password":        user.Password,
	},
	)

	if err != nil {
		t.Error(err.Error())
	}

	reader := bytes.NewBuffer(data)
	resp, _ := http.Post(endpoint, "application/json", reader)
	assert.Equal(t, 200, resp.StatusCode)
}
