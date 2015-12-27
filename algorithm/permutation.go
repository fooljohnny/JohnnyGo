package algorithm

import "fmt"

func permute() {
	var m, n, i, j, t, s int
	var a [30]int

	n = 5
	m = 1
	i = 1
	a[i] = 1
	for {
		t = 1
		for j = 1; j < i; j++ {
			if a[j] == a[i] {
				t = 0
				break
			}
		}
		if t && i == m {
			// here find one sequence
			s++
			for j = 1; j < m; j++ {
				fmt.Print("%d", a[j])
			}
			fmt.Printf(" ")
			if s%10 == 0 {
				fmt.Println("")
			}
		}
		if t && i < m {
			// finding next number
			i++
			a[i] = 1
			continue
		}
		for a[i] == n {
			i--
		}
		if i > 0 {
			a[i]++
		} else {
			break
		}
	}
	fmt.Println("s = %d", s)
}
