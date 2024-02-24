package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/tdd/class5/api/model"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("empty person", func(t *testing.T) {
		// Arrange
		m := NewMemory()
		// Act
		err := m.Create(nil)
		// Assert
		assert.EqualError(t, err, model.ErrPersonCanNotBeNil.Error())
	})

	t.Run("validate currentID is increasing", func(t *testing.T) {
		// Arrange
		m := NewMemory()
		p := model.Person{Name: "Victor"}

		// Act
		err := m.Create(&p)
		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 1, m.currentID)
	})
}
