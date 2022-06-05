package slices

import "github.com/Pauloo27/toolkit/funcs"

// Find returns a pointer to the first element in the slice where the predicate
// is true. If no element is found, it returns nil.
func Find[T any](slice []T, predicate funcs.Predicate[T]) *T {
	for _, v := range slice {
		if predicate(v) {
			return &v
		}
	}
	return nil
}

// FindOr returns the value of the first element in the slice where the predicate
// is true. If no element is found, it returns the default value.
func FindOr[T any](slice []T, predicate funcs.Predicate[T], defaultValue T) T {
	v := Find(slice, predicate)
	if v == nil {
		return defaultValue
	}
	return *v
}

// Map returns a new slice with the results of applying the mapper function
// to each element in the original slice.
func Map[T any, I any](slice []T, mapper funcs.Function[T, I]) []I {
	res := make([]I, len(slice))
	for i, v := range slice {
		res[i] = mapper(v)
	}
	return res
}

// FindIndex returns the index of the first element in the slice where
// the predicate is true. If no element is found, it returns -1.
func FindIndex[T any](slice []T, predicate funcs.Predicate[T]) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Some returns true if any element in the slice satisfies the predicate and
// false otherwise.
func Some[T any](slice []T, predicate funcs.Predicate[T]) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Every returns true if every element in the slice satisfies the predicate
// and false otherwise.
func Every[T any](slice []T, predicate funcs.Predicate[T]) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// ForEach iterates over the slice and calls the consumer functions on each
// element.
func ForEach[T any](slice []T, consumer funcs.Consumer[T]) {
	for _, v := range slice {
		consumer(v)
	}
}

// Filter returns a new slice with all elements that satisfy the predicate.
// If no elements satisfy the predicate, an empty slice is returned.
// The original slice is not modified.
func Filter[T any](slice []T, predicate funcs.Predicate[T]) []T {
	var res []T
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// Reduce is kinda hard to explain. Google that up.
func Reduce[T any](slice []T, reducer func(prevValue, curValue T) T, initialValue T) T {
	prev := initialValue
	for _, v := range slice {
		prev = reducer(prev, v)
	}
	return prev
}
