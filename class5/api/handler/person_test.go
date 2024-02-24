package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPerson_Create_wrong_struct(t *testing.T) {
	// Arrange
	badJsonType := []byte(`{"name": 123, "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(badJsonType))
	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	// Arrange
	p := newPerson(&StorageMockOK{})

	// Act
	err := p.create(ctx)
	if err != nil {
		t.Errorf("we get an error %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := response{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		t.Errorf("Error cannot decode response of w.Body: %v", err)
	}
	// Assert
	wantMsg := "La persona no tiene una estructura correcta"
	assert.Equal(t, wantMsg, res.Message)
}

func TestPerson_Create_wrong_storage(t *testing.T) {
	correctJson := []byte(`{"name": "Victor", "age": 35}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(correctJson))
	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	// Arrange
	p := newPerson(&StorageMockError{})
	// Act
	err := p.create(ctx)
	if err != nil {
		t.Errorf("we get error: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	res := response{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		t.Errorf("Error cannot decode response of w.Body: %v", err)
	}

	// Assert
	wantMsg := "Hubo un problema al crear la persona"
	assert.Equal(t, wantMsg, res.Message)
}

func TestPerson_Create(t *testing.T) {
	correctJson := []byte(`{"name": "Victor", "age": 35}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(correctJson))
	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	// Arrange
	p := newPerson(&StorageMockOK{})
	// Act
	err := p.create(ctx)
	if err != nil {
		t.Errorf("we get error: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	res := response{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		t.Errorf("Error cannot decode response of w.Body: %v", err)
	}

	// Assert
	wantMsg := "Persona creada correctamente"
	assert.Equal(t, wantMsg, res.Message)
}
