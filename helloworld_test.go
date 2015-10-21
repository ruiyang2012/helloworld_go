package hello_test

// to run test: goapp test  github.com/ruiyang2012/helloworld_go/

import (
	"github.com/ruiyang2012/helloworld_go"

	"appengine/aetest"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = hello.SetupRoutes()

func TestRunEmptyWithEnv(t *testing.T) {
	inst, e := aetest.NewInstance(nil)
	assert.NoError(t, e)
	defer inst.Close()
	req, e := inst.NewRequest("GET", "/ping", nil)
	assert.NoError(t, e)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "pong")

}

func TestRunV1Count(t *testing.T) {
	inst, e := aetest.NewInstance(nil)
	assert.NoError(t, e)
	defer inst.Close()
	_, e0 := inst.NewRequest("DELETE", "/v1/count", nil)
	assert.NoError(t, e0)
	req, e := inst.NewRequest("GET", "/v1/count", nil)
	assert.NoError(t, e)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "0")

	req1, e1 := inst.NewRequest("POST", "/v1/count/add/15", nil)
	assert.NoError(t, e1)
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)

	assert.Equal(t, resp1.Body.String(), "15")

	req2, e2 := inst.NewRequest("POST", "/v1/count/subtract/13", nil)
	assert.NoError(t, e2)
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)

	assert.Equal(t, resp2.Body.String(), "2")

	req3, e3 := inst.NewRequest("POST", "/v1/count/multiple/3", nil)
	assert.NoError(t, e3)
	resp3 := httptest.NewRecorder()
	router.ServeHTTP(resp3, req3)

	assert.Equal(t, resp3.Body.String(), "6")
}
