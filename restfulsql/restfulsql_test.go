package restfulsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	query := `["AND", ["a", ["AND", ["b"], [2]]], [10, ""]]`
	parser := NewRestfulSQLParser(query)
	rsql, err := parser.Parse()
	require.NoError(err)
	require.NotNil(rsql)

	assert.Equal("AND", rsql.Mode)
	require.NotEmpty(rsql.Fields)
	require.Len(rsql.Fields, 2)
	assert.Equal([]interface{}{
		"AND",
		[]interface{}{"b"},
		[]interface{}{2.0},
	}, rsql.Fields[1].([]interface{}))
	assert.Equal([]interface{}{10.0, ""}, rsql.Values)
}
