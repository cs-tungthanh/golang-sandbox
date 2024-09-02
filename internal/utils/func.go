package utils

type mapFunc[E any] func(E) E
type keepFunc[E any] func(E) bool
type reduceFunc[E any] func(cur, next E) E

// Map applies a function to each element of a slice and returns a new slice with the results.
//
// s := []string{"a", "b", "c"}
// fmt.Println(Map(s, strings.ToUpper))
// Output: [A B C]
func Map[S ~[]E, E any](s S, f mapFunc[E]) S {
	result := make(S, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

// Filter returns a new slice containing only the elements of the slice that satisfy the predicate f.
//
// s := []int{1, 2, 3, 4}
//
//	fmt.Println(Filter(s, func(v int) bool {
//	    return v%2 == 0
//	}))
func Filter[S ~[]E, E any](s S, f keepFunc[E]) S {
	result := S{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies a function to each element of a slice, combining the results.

// s := []int{1, 2, 3, 4}
//
//	sum := Reduce(s, 0, func(cur, next int) int {
//	    return cur + next
//	})
func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}
