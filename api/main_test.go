package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	log "github.com/sirupsen/logrus"
)

func TestDefaultRoute(t *testing.T) {
	router := setupRouter()

	expectedResponse := gin.H{"data": "hello"}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		log.Fatal("Error while unmarshaling body response, ", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponse["data"], response["data"])
}
