package int

import (
	"fmt"
	"testing"
)

func TestPerm2(t *testing.T) {
	counter := 0
	n := 60
	unique := make(map[string][2]int)
	result := [2]int{-1, -1}
	v := Perm2(n, func(v [2]int) bool {
		t.Logf("%v", v)
		counter++
		unique[fmt.Sprintf("%v", v)] = v
		return false
	})
	per := n * (n - 1) / 2
	if counter != per {
		t.Errorf("Counter is wrong %v is not %v", counter, per)
	}
	if v != result {
		t.Errorf("Result is wrong expect -1,-1 got %v", v)
	}
	if len(unique) != per {
		t.Errorf("Unique is wrong expect %v got %v", per, unique)
	}

}
func TestPerm4(t *testing.T) {
	counter := 0
	n := 60
	unique := make(map[string][4]int)
	result := [4]int{-1, -1, -1, -1}
	v := Perm4(n, func(v [4]int) bool {
		t.Logf("%v", v)
		counter++
		unique[fmt.Sprintf("%v", v)] = v
		return false
	})
	per := n * (n - 1) * (n - 2) * (n - 3) / (4 * 3 * 2)
	if counter != per {
		t.Errorf("Counter is wrong expected %v got %v", per, counter)
	}
	if v != result {
		t.Errorf("Result is wrong expect -1,-1,-1,-1 got %v", v)
	}
	if len(unique) != per {
		t.Errorf("Unique is wrong expect %v got %v", per, len(unique))
	}

}
func TestPerm3(t *testing.T) {
	counter := 0
	n := 60
	unique := make(map[string][3]int)
	result := [3]int{-1, -1, -1}
	v := Perm3(n, func(v [3]int) bool {
		t.Logf("%v", v)
		counter++
		unique[fmt.Sprintf("%v", v)] = v
		return false
	})
	per := n * (n - 1) * (n - 2) / (3 * 2)
	if counter != per {
		t.Errorf("Something is wrong expected %v got %v", per, counter)
	}
	if v != result {
		t.Errorf("Something is wrong expect -1,-1,-1 got %v", v)
	}
	if len(unique) != per {
		t.Errorf("Something is wrong expect %v got %v", per, len(unique))
	}

}
func TestComb(t *testing.T) {
	v := Comb(uint64(6), uint64(4))
	if v != 15 {
		t.Errorf("Something is wrong expect %v got %v", 15, v)
	}
}
func TestPerm(t *testing.T) {
	set := make(map[string]bool)
	n := 8
	d := 4
	Perm(n, d, func(v []int) bool {
		set[fmt.Sprint(v)] = true
		return false
	})

	per := len(set)
	comb := Comb(uint64(n), uint64(d))
	if per != int(comb) {
		t.Errorf("Something is wrong permutaions %v combinations %v", per, comb)
	}
}

func TestFactorSum(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 11
	elements := 3
	all := true
	FactorSum(values, sum, elements, all)
	all = false
	FactorSum(values, sum, elements, all)
}
