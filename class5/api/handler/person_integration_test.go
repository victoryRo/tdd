package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/tdd/class5/api/model"
	"github.com/victoryRo/tdd/class5/api/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDB(t *testing.T) {
	correctJson := []byte(`{"name": "Victor", "age": 35}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(correctJson))
	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	// Act
	store := storage.NewPsql()
	t.Cleanup(func() { // execute at the end
		cleanDatabase(t, store)
	})

	p := newPerson(&store)
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

type responsePerson struct {
	MessageType string         `json:"message_type"`
	Message     string         `json:"message"`
	Data        []model.Person `json:"data"`
}

func TestPersonGetAll(t *testing.T) {
	insertData(t) // insert some persons

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	// Act
	store := storage.NewPsql()
	t.Cleanup(func() { // execute at the end
		cleanDatabase(t, store)
	})

	p := newPerson(&store)
	err := p.getAll(ctx)
	if err != nil {
		t.Errorf("I didn't expect error, we got: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	res := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		t.Errorf("error cannot decode response of w.Body: %v", err)
	}
	// Assert
	assert.Equal(t, 2, len(res.Data))
	assert.Equal(t, "Victor", res.Data[0].Name)
}

func insertData(t *testing.T) {
	t.Helper()

	store := storage.NewPsql()
	defer store.CloseDB()

	err := store.Create(&model.Person{
		Name: "Victor",
		Age:  35,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	err = store.Create(&model.Person{
		Name: "Acertivo",
		Age:  34,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
}

// cleanDatabase deletes all data from the database.
func cleanDatabase(t *testing.T, db storage.Psql) {
	t.Helper()

	defer db.CloseDB()
	err := db.DeleteAll()
	if err != nil {
		fmt.Errorf("error deleting database: %v", err).Error()
	}
}
