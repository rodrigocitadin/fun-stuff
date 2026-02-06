package countminsketch

import (
	"fmt"
	"math/rand/v2"
)

const p = uint(1<<31 - 1)

type cms struct {
	width int
	depth int
	a     []uint
	b     []uint
	table [][]uint
}

type CMS interface {
	Update(i, c uint)
	Estimate(i uint) uint
}

func (cms *cms) Estimate(i uint) uint {
	var min uint
	for row := 0; row < cms.depth; row++ {
		col := cms.index(row, i)
		v := cms.table[row][col]
		if row == 0 || v < min {
			min = v
		}
	}
	return min
}

func (cms *cms) Update(i, c uint) {
	for row := 0; row < cms.depth; row++ {
		col := cms.index(row, i)
		cms.table[row][col] += c
	}
}

func NewCMS(w, d int) CMS {
	// guarantees power of 2
	if w <= 0 || (w&(w-1)) != 0 {
		panic("w must be a power of two")
	}

	// guarantees valid integer
	if d <= 0 {
		panic("d must be > 0")
	}

	table := make([][]uint, d)
	for i := range d {
		table[i] = make([]uint, w)
	}

	a := make([]uint, d)
	b := make([]uint, d)
	for i := range d {
		a[i] = uint(rand.IntN(int(p-1)) + 1)
		b[i] = uint(rand.IntN(int(p-1)) + 1)
	}

	return &cms{
		width: w,
		depth: d,
		a:     a,
		b:     b,
		table: table,
	}
}

func modP(x uint) uint {
	x = (x & p) + (x >> 31)
	if x >= p {
		x -= p
	}
	return x
}

func (cms *cms) index(row int, x uint) int {
	h := modP(cms.a[row]*x + cms.b[row])
	return int(h & uint(cms.width-1)) // power-of-two mod
}

func main() {
	cms := NewCMS(65_536, 10)
	model := make(map[uint]uint)

	for n := range 1_000_000 {
		i := uint(rand.IntN(1_000))
		c := uint(rand.IntN(10))
		model[i] += c
		cms.Update(i, c)
		if n%10_000 == 0 {
			fmt.Printf("model[i] %+v estimate %+v\n", model[i], cms.Estimate(i))
		}
	}
}
