package funcs_test

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/pauloo27/toolkit/funcs"
)

func TestFunction(t *testing.T) {
	var _ funcs.Function[int, string] = strconv.Itoa
}

func TestConsumer(t *testing.T) {
	var _ funcs.Consumer[func()] = t.Cleanup
}

func TestUnaryOperator(t *testing.T) {
	var _ funcs.UnaryOperator[float64] = math.Abs
}

func TestBinaryOperator(t *testing.T) {
	var _ funcs.BinaryOperator[float64] = math.Pow
}

func TestPredicate(t *testing.T) {
	var _ funcs.Predicate[float64] = math.IsNaN
}

func TestSuplier(t *testing.T) {
	var _ funcs.Suplier[int] = rand.Int
}
