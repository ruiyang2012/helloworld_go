package hello_test

// to run test: goapp test  github.com/ruiyang2012/helloworld_go/


import (
	"github.com/ruiyang2012/helloworld_go"

	"io/ioutil"

	"net/http"
	"net/http/httptest"

	"testing"


	"github.com/stretchr/testify/assert"
)

func testRequest(t *testing.T, url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	assert.NoError(t, err)

	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, "it worked", string(body), "resp body should match")
	assert.Equal(t, "200 OK", resp.Status, "should get a 200")
}


func TestRunEmptyWithEnv(t *testing.T) {
	router := hello.SetupRoutes()
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "pong")	
	
}