package handler_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/tdd/class4/handler"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	r := httptest.NewRequest("http.MethodGet", "/", nil)

	// Act
	handler.Get(w, r)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("got %d, want %d", w.Code, http.StatusOK)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("we want %s, and we got %s", "application/json", w.Header().Get("Content-Type"))
	}

	// Arrange
	p := handler.Person{}
	err := json.NewDecoder(w.Body).Decode(&p)
	if err != nil {
		t.Error(err)
	}
	// Assert
	assert.Equal(t, "Victoria", p.Name)

	// Arrange
	per := handler.Person{
		Name: "Victoria",
		Age:  29,
	}
	// Assert
	if !reflect.DeepEqual(p, per) {
		t.Errorf("we want %v, got %v", per, p)
	}
}
