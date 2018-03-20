package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one [][]int
}

func Test46(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{1, 2, 3}},
			ans{[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, permute(p.one), "输入:%v", p)
	}
}
