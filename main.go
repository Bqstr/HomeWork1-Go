package awesomeProject3

import "fmt"

func sortSlice(s []int) {
	for a := 0; a < len(s)-1; a++ {
		for b := a + 1; b < len(s); b++ {
			if s[a] > s[b] {
				var ss = s[a]
				s[a] = s[b]
				s[b] = ss

			}
		}

	}

}

func IncrementOdd(s []int) {
	for a := 0; a < len(s); a++ {
		if a%2 != 0 {
			s[a]++
		}
	}
}

func PrintSlice(s []int) {
	fmt.Println(s)
}

func ReverseSlice(s []int) {
	var b = len(s) - 1
	for a := 0; a < len(s)/2; a++ {
		ss := s[a]
		s[a] = s[b]
		s[b] = ss
		b--
	}

}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(s []int) {
		dst(s)
		for _, a := range src {
			a(s)
		}
	}
}
