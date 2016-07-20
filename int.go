package int

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
func Comb(n int, draw int) (res int) {
	res = n
	for i := 1; i < draw; i++ {
		res = res * (n - i)
	}

	return res
}
