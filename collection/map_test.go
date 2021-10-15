package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMapWithSize(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValue("input must be a slice", func() { ToMapWithSize(1, "a", 1) })
	type foo struct {
		ID   int
		Name string
	}

	inputs := make([]*foo, 2)
	for i := 0; i < 2; i++ {
		inputs[i] = &foo{
			ID:   i,
			Name: "A",
		}
	}

	assert.Nil(ToMapWithSize(inputs, "empty", 1))
	result, ok := ToMapWithSize(inputs, "ID", 1).(map[int]*foo)
	assert.True(ok)

	for _, elem := range inputs {
		v, ok := result[elem.ID]
		assert.True(ok)
		assert.Equal(elem, v)
	}

}

func TestToMap(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValue("input must be a slice", func() { ToMap(1, "a") })
	type foo struct {
		ID   int
		Name string
	}

	inputs := make([]*foo, 2)
	for i := 0; i < 2; i++ {
		inputs[i] = &foo{
			ID:   i,
			Name: "A",
		}
	}

	assert.Nil(ToMap(inputs, "empty"))
	result, ok := ToMap(inputs, "ID").(map[int]*foo)
	assert.True(ok)

	for _, elem := range inputs {
		v, ok := result[elem.ID]
		assert.True(ok)
		assert.Equal(elem, v)
	}
}
