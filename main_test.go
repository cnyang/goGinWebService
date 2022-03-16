package main

import (
	//"iBP/models"
	"iBP/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

// performRequest 在test裡面用這個function 來call api
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Test_version testing of web route /api/BasicService/version should return version
func Test_verions(t *testing.T) {
	router := routes.SetupRouter()

	w := performRequest(router, "GET", "/api/BasicService/version")
	// 假設正確會回傳 200
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "version")

	// Convert the JSON response to a map
	var response map[string]string
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["version"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, viper.GetString("env.version"), value)
}

// Test_healthcheck should return ok
func Test_healthcheck(t *testing.T) {
	router := routes.SetupRouter()

	w := performRequest(router, "GET", "/api/BasicService/healthcheck")
	// 假設正確會回傳 200
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthcheck")

	// Convert the JSON response to a map
	var response map[string]string
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["healthcheck"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "ok", value)
}
