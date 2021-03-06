package int

import (
	"fmt"
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
		panic(fmt.Sprintf("n: %v,d: %v", n, d))
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
func Comb2(n uint64, draw uint64) (res uint64) { //TODO make a safe version. 52 over 22 look like the limit, and we need 52 over 23
	if n == draw {
		res = 1
	} else {
		if n < draw {
			panic(fmt.Sprintf("n: %v must not be smaller than m: %v", n, draw))
		}
		if draw > n/2 {
			draw = n - draw
		}
		nv := make([]uint64, draw)
		dv := make([]uint64, draw)
		for i := uint64(0); i < draw; i++ {
			nv[i] = n - i
			dv[i] = draw - i
		}
		if draw > 10 {
			res = uint64(1)
			for _, ni := range nv {
				r := ni
				found := -1
				for i := 0; i < len(dv); i++ {
					di := dv[i]
					if di != uint64(1) {
						if ni%di == 0 {
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
		} else {
			res = uint64(1)
			for _, ni := range nv {
				res = res * ni
			}
			for _, di := range dv {
				res = res / di
			}
		}
	}
	return res
}
func Comb(n uint64, draw uint64) (res uint64) { //TODO make a safe version. 52 over 23 look like the limit
	if n == draw {
		res = 1
	} else {
		if n < draw {
			panic(fmt.Sprintf("n: %v must not be smaller than m: %v", n, draw))
		}
		if draw > n/2 {
			draw = n - draw
		}
		nv := make([]uint64, draw)
		dv := make([]uint64, draw)
		for i := uint64(0); i < draw; i++ {
			nv[i] = n - i
			dv[i] = draw - i
		}
		if draw > 10 {
			res = uint64(1)
			reduce(nv, dv)
			reduce(dv, nv)

			for _, ni := range nv {
				if ni != 1 {
					res = res * ni
				}
			}

			for _, di := range dv {
				if di != 1 {
					res = res / di
				}
			}
		} else {
			res = uint64(1)
			for _, ni := range nv {
				res = res * ni
			}
			for _, di := range dv {
				res = res / di
			}
		}
	}
	return res
}
func reduce(nv, dv []uint64) {
	for nix, ni := range nv {
		r := ni
		found := -1
		for i := 0; i < len(dv); i++ {
			di := dv[i]
			if di != uint64(1) && ni > di {
				if ni%di == 0 {
					found = i
					break
				}
			}
		}
		if found != -1 {
			r = ni / dv[found]
			dv[found] = 1
			nv[nix] = r
		}
	}
}

//FactorSum finds all the possible ways to factor a sum with the numbers of elements.
//Result indexes is nil if none was found.
func FactorSum(values []int, sum int, elements int, all bool) (indexes [][]int) {
	n := len(values)
	if n >= elements {
		Perm(n, elements, func(is []int) (stop bool) {
			s := 0
			for _, index := range is {
				s = s + values[index]
			}
			if sum == s {
				e := make([]int, elements)
				for i, index := range is {
					e[i] = index
				}
				indexes = append(indexes, e)
				if !all {
					stop = true
				}
			}
			return stop
		})
	}
	return indexes
}
