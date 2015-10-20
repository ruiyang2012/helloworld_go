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
	
}