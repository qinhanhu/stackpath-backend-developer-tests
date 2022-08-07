package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/kinbiko/jsonassert"
	"github.com/stackpath/backend-developer-tests/rest-service/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	app := gin.Default()
	router := app.Group("")
	routers.AddRouters(router)
	return app
}

// APITestCase represents the data needed to describe an API test case.
type APITestCase struct {
	Name         string
	Method       string
	URL          string
	Body         string
	Header       http.Header
	WantStatus   int
	WantResponse string
}

// Endpoint tests an HTTP endpoint using the given APITestCase spec.
func Endpoint(t *testing.T, router *gin.Engine, tc APITestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		req := httptest.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
		if tc.Header != nil {
			req.Header = tc.Header
		}
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, tc.WantStatus, res.Code, "status mismatch")
		if tc.WantResponse != "" {
			ja := jsonassert.New(t)
			ja.Assertf(res.Body.String(), tc.WantResponse)
		}
	})
}
