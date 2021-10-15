package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	assert.True(Contains([]string{"foo", "bar"}, "bar"))
	assert.False(Contains([]string{"foo", "bar"}, "baz"))

	assert.True(Contains([]int{114, 514}, 114))
	assert.False(Contains([]int{114, 514}, 233))

	assert.True(Contains([2]int{114, 514}, 114))
	assert.False(Contains([2]int{114, 514}, 233))
}

func TestJoinInt64Array(t *testing.T) {
	assert := assert.New(t)

	arr := []int64{
		123,
		456,
		789,
	}

	result := JoinInt64(arr, "|")
	assert.Equal("123|456|789", result)
}

func TestIntersectionInt64(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	a1 := []int64{1, 1}
	a2 := []int64{2, 2}
	intersection := IntersectionInt64(a1, a2)
	require.Empty(intersection)

	a1 = []int64{1, 2}
	a2 = []int64{2, 3}
	intersection = IntersectionInt64(a1, a2)
	require.NotEmpty(intersection)
	assert.Equal([]int64{2}, IntersectionInt64(a1, a2))

	a1 = []int64{1, 2}
	a2 = []int64{3, 4}
	intersection = IntersectionInt64(a1, a2)
	require.Empty(intersection)
}

func TestGroupInt64(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	var s1 []int64
	length := 2
	result := GroupInt64(s1, length)
	require.Len(result, 1)
	assert.Empty(result[0])

	s1 = make([]int64, 0)
	result = GroupInt64(s1, length)
	require.Len(result, 1)
	assert.Empty(result[0])

	s1 = []int64{1, 2, 3}
	result = GroupInt64(s1, length)
	require.Len(result, 2)
	assert.Equal(result, [][]int64{{1, 2}, {3}})

	s1 = []int64{1, 2, 3, 4}
	result = GroupInt64(s1, length)
	require.Len(result, 2)
	assert.Equal(result, [][]int64{{1, 2}, {3, 4}})
}

func TestDiffInt64(t *testing.T) {
	assert := assert.New(t)
	a := []int64{0, 1, 2}
	var b []int64
	assert.ElementsMatch([]int64{0, 1, 2}, DiffInt64(a, b))
	assert.ElementsMatch([]int64{0, 1, 2}, DiffInt64(b, a))

	b = []int64{1, 2, 4, 5}
	assert.ElementsMatch([]int64{0, 4, 5}, DiffInt64(a, b))

	b = []int64{1, 2, 4, 5}
	assert.Empty(DiffInt64(b, b))
}

func TestFindInt64Duplicates(t *testing.T) {
	assert := assert.New(t)

	list := []int64{1, 2, 3, 4, 5}
	repeatList := FindDuplicateInt64(list)
	assert.Len(repeatList, 0)

	list = []int64{1, 1, 2}
	repeatList = FindDuplicateInt64(list)
	assert.Equal([]int64{1}, repeatList)
}

func TestSplitToInt64(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	r, err := SplitToInt64("1,2,3,4", ",")
	require.NoError(err)
	assert.Equal([]int64{1, 2, 3, 4}, r)

	r, err = SplitToInt64("", ",")
	require.NoError(err)
	assert.Empty(r)

	_, err = SplitToInt64("test", ",")
	assert.Error(err)
}
