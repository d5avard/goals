package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestAllRoute(t *testing.T) {
	var r *gin.Engine = gin.Default()
	AddGoalsRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/goals/", nil)
	r.ServeHTTP(w, req)

	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Get all goals", dat["message"])
}

func TestGetRoute(t *testing.T) {
	var r *gin.Engine = gin.Default()
	AddGoalsRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/goals/1", nil)
	r.ServeHTTP(w, req)

	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Get goal 1", dat["message"])
}

func TestCreateRoute(t *testing.T) {
	var r *gin.Engine = gin.Default()
	AddGoalsRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/goals/", nil)
	r.ServeHTTP(w, req)

	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Create a new goal", dat["message"])
}

func TestUpdateRoute(t *testing.T) {
	var r *gin.Engine = gin.Default()
	AddGoalsRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/goals/1", nil)
	r.ServeHTTP(w, req)

	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Update goal 1", dat["message"])
}

func TestDeleteRoute(t *testing.T) {
	var r *gin.Engine = gin.Default()
	AddGoalsRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/goals/1", nil)
	r.ServeHTTP(w, req)

	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Delete goal 1", dat["message"])
}
