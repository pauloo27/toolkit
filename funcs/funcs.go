package funcs

// Function takes T T and returns R
type Function[T, R any] func(t T) R

// Consumer takes T param and returns nothing
type Consumer[T any] func(t T)

// UnaryOperator takes T param and returns T
type UnaryOperator[T any] func(t T) T

// BinaryOperator takes two T params and returns T
type BinaryOperator[T any] func(t1, t2 T) T

// Predicate takes T param and returns bool
type Predicate[T any] func(t T) bool

// Suplier takes nothing and returns T
type Suplier[T any] func() T
