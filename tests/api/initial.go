package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	database "github.com/nugrohosam/gosampleapi/services/databases"
	httpServe "github.com/nugrohosam/gosampleapi/services/http"
	viper "github.com/spf13/viper"
)

// Routes ...
var Routes *gin.Engine

// InitialTest ...
func InitialTest(t *testing.T) {
	// initial call to envinronment variable
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../../.env.test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}

	if err := database.ConnOrm(); err != nil {
		t.Error(err.Error())
		panic(err)
	}

	httpServe.Prepare()

	Routes = httpServe.Routes
}

// PerformRequest ...
func PerformRequest(r http.Handler, method, endpoint, contentType string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, endpoint, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
