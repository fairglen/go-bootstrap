package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://dev.to/quii/learn-go-by-writing-tests-arrays-and-slices-ahm

func TestSum(t *testing.T){
	scenarios:= []struct{
		numbers []int
		expSum int
	}{
		{
			numbers: []int{1,2,3,4,5},
			expSum: 15,
		},
		{
			numbers: []int{1,2,3,4},
			expSum: 10,
		},
		{
			numbers: []int{1,1},
			expSum: 2,
		},
	}
	for _, scenario := range scenarios {
		t.Run("returns sum of array numbers", func(t *testing.T){
			res := Sum(scenario.numbers)
			require.Equal(t, scenario.expSum, res)
		})
	}

}
