package int

import (
	"math"
)

func Perm2(n int, eval func([2]int) bool) [2]int {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := [2]int{i, j}
			if eval(v) {
				return v
			}
		}
	}
	return [2]int{-1, -1}
}

func Perm3(n int, eval func([3]int) bool) [3]int {
	for i0 := 0; i0 < n-2; i0++ {
		for i1 := i0 + 1; i1 < n-1; i1++ {
			for i2 := i1 + 1; i2 < n; i2++ {
				v := [3]int{i0, i1, i2}
				if eval(v) {
					return v
				}
			}
		}
	}
	return [3]int{-1, -1, -1}
}
func Perm4(n int, eval func([4]int) bool) [4]int {
	for i0 := 0; i0 < n-3; i0++ {
		for i1 := i0 + 1; i1 < n-2; i1++ {
			for i2 := i1 + 1; i2 < n-1; i2++ {
				for i3 := i2 + 1; i3 < n; i3++ {
					v := [4]int{i0, i1, i2, i3}
					if eval(v) {
						return v
					}
				}
			}
		}
	}
	return [4]int{-1, -1, -1, -1}
}
func Perm(n int, d int, eval func([]int) bool) (per []int) {
	if n < d || d < 1 || n < 1 {
		panic("Illegal values")
	}
	per = make([]int, d)
	for i := 0; i < d; i++ {
		per[i] = i
	}
	if !eval(per) && n != d {
	Loop:
		for {
			//Search loop
			foundix := -1
			for i := 1; i <= d; i++ {
				if per[d-i] != n-i {
					foundix = d - i
					break
				}
			}
			if foundix != -1 {
				per[foundix] = per[foundix] + 1
				for j := foundix + 1; j < d; j++ {
					per[j] = per[j-1] + 1
				}
				if eval(per) {
					break Loop
				}
			} else {
				break Loop
			}
		} //Loop
	}
	return per
}
func Comb(n uint64, draw uint64) (res uint64) {
	nv := make([]uint64, draw)
	dv := make([]uint64, draw)
	for i := uint64(0); i < draw; i++ {
		nv[i] = n - i
		dv[i] = draw - i
	}
	res = uint64(1)
	for _, ni := range nv {
		r := ni
		found := -1
		for i, di := range dv {
			if di != uint64(1) {
				if math.Remainder(float64(ni), float64(di)) == 0 {
					found = i
					break
				}
			}
		}
		if found != -1 {
			r = ni / dv[found]
			dv[found] = 1
		}
		res = res * r

	}
	for _, di := range dv {
		if di != 1 {
			res = res / di
		}
	}

	return res
}
