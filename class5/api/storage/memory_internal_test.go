package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/tdd/class5/api/model"
)

func TestMemoryCreate(t *testing.T) {
	t.Run("chack currentID", func(t *testing.T) {

		// Arrange
		m := NewMemory()
		p := model.Person{Name: "Acertivo"}

		// Act
		err := m.Create(&p)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 1, m.currentID)
	})
}
