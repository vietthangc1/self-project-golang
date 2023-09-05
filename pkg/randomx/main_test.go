package randomx

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	lenString := rand.Intn(100)
	outString := RandString(lenString)

	require.Equal(t, lenString, len(outString))
}

func TestRandomInt(t *testing.T) {
	min := 10
	max := 100

	num_test := 10
	for i:=0; i<num_test; i++ {
		outInt := RandInt(10, 100)
		require.LessOrEqual(t, outInt, max)
		require.GreaterOrEqual(t, outInt, min)
	}
}

func TestErrorRandomInt(t *testing.T) {
	min := 10
	max := 9

	out := RandInt(min, max)
	require.Equal(t, 0, out)
}
