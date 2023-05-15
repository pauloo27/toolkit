package slices_test

import (
	"testing"

	"github.com/pauloo27/toolkit/slices"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	assert.Equal(t, 4, *slices.Find([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 3
	}))

	assert.Equal(t, (*int)(nil), slices.Find([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 100
	}))

	type User struct {
		ID   int
		Name string
	}

	users := []User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jack"},
	}

	assert.Equal(t, &User{ID: 2, Name: "Jane"}, slices.Find(users, func(u User) bool {
		return u.ID == 2
	}))

}

func TestFindOr(t *testing.T) {
	assert.Equal(t, 4, slices.FindOr([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 3
	}, 0))

	assert.Equal(t, 0, slices.FindOr([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 100
	}, 0))

	assert.Equal(t, 1337, slices.FindOr([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 10000
	}, 1337))

	assert.Equal(t, -123, slices.FindOr([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return false
	}, -123))
}

func TestFindIndex(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	assert.Equal(t, 1, slices.FindIndex(slc, func(i int) bool {
		return i > 3
	}))

	assert.Equal(t, -1, slices.FindIndex(slc, func(i int) bool {
		return i > 100
	}))
}

func TestSome(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	assert.True(t, slices.Some(slc, func(i int) bool {
		return i > 3
	}))

	assert.False(t, slices.Some(slc, func(i int) bool {
		return i > 100
	}))

	assert.True(t, slices.Some(slc, func(i int) bool {
		return i%2 == 0
	}))

	assert.True(t, slices.Some(slc, func(i int) bool {
		return i%2 == 1
	}))
}

func TestEvery(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	assert.True(t, slices.Every(slc, func(i int) bool {
		return i > 0
	}))

	assert.False(t, slices.Every(slc, func(i int) bool {
		return i > 3
	}))
}

func TestMap(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	assert.Equal(t, slc, slices.Map(slc, func(i int) int {
		return i
	}))

	assert.Equal(t, []int{2, 10, 6, 8, 4}, slices.Map(slc, func(i int) int {
		return i * 2
	}))
}

func TestForEach(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	var sum int
	slices.ForEach(slc, func(i int) {
		sum += i
	})
	assert.Equal(t, 15, sum)

	slc = []int{1, 2, 4}
	sum = 0
	callback := func(i int) {
		sum += i
	}
	slices.ForEach(slc, callback)
	assert.Equal(t, 7, sum)

}

func TestFilter(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	assert.Equal(t, []int{1, 5, 3}, slices.Filter(slc, func(i int) bool {
		return i%2 == 1
	}))

	assert.Equal(t, ([]int)(nil), slices.Filter(slc, func(i int) bool {
		return i%2 == 22 // imposible?
	}))

	assert.Len(t, slices.Filter(slc, func(i int) bool {
		return i%2 == 32 // also imposible?
	}), 0)
}

func TestReduce(t *testing.T) {
	slc := []int{1, 5, 3, 4, 2}

	sum := func(n1, n2 int) int {
		return n1 + n2
	}

	biggest := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}

	assert.Equal(t, 15, slices.Reduce(slc, sum, 0))
	assert.Equal(t, 5, slices.Reduce(slc, biggest, 0))
}
