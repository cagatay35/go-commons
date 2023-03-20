package util

import "strings"

func ContainInt(a []int, v int) bool {
	for _, value := range a {
		if value == v {
			return true
		}
	}
	return false
}

func ContainString(a []string, v string) bool {
	for _, value := range a {
		if value == v {
			return true
		}
	}
	return false
}

func ContainsStringEqualIgnoreCase(list []string, x string) bool {
	for _, n := range list {
		if strings.EqualFold(n, x) {
			return true
		}
	}
	return false
}

func AppendUniqueString(a []string, s ...string) []string {
	m := map[string]struct{}{}
	a = append(a, s...)

	var result []string

	for _, val := range a {
		m[val] = struct{}{}
	}

	// preserve the order when appending into result.
	for _, val := range a {
		if _, ok := m[val]; ok {
			result = append(result, val)
			delete(m, val)
		}
	}

	return result
}

func AppendUniqueInt(a []int, s ...int) []int {
	m := map[int]struct{}{}
	a = append(a, s...)

	var result []int

	for _, val := range a {
		m[val] = struct{}{}
	}

	// preserve the order when appending into result.
	for _, val := range a {
		if _, ok := m[val]; ok {
			result = append(result, val)
			delete(m, val)
		}
	}

	return result
}

func AppendUniqueLong(a []int64, s ...int64) []int64 {
	m := map[int64]struct{}{}
	a = append(a, s...)

	var result []int64

	for _, val := range a {
		m[val] = struct{}{}
	}

	// preserve the order when appending into result.
	for _, val := range a {
		if _, ok := m[val]; ok {
			result = append(result, val)
			delete(m, val)
		}
	}

	return result
}

func Diff(slice1 map[string]string, slice2 map[string]string) []string {
	var diff []string
	for i := 0; i < 2; i++ {
		for s1 := range slice1 {
			found := false
			for s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}

			if !found {
				diff = append(diff, s1)
			}
		}

		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
