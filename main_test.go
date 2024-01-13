package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetOrders(t *testing.T) {
	// Set up a test router
	router := gin.Default()
	router.GET("/api/orders", getOrders)

	// Create a mock request to /api/orders
	req, err := http.NewRequest("GET", "/api/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, recorder.Code)
}
