package problem0827

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{{1, 0}, {0, 1}},
		3,
	},

	{
		[][]int{{1, 1}, {1, 0}},
		4,
	},

	{
		[][]int{{1, 1}, {1, 1}},
		4,
	},

	// 可以有多个 testcase
}

func Test_largestIsland(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largestIsland(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_largestIsland(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestIsland(tc.grid)
		}
	}
}
