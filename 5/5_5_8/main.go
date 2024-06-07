package main

import "fmt"

func main() {
	var n, group, taxi int
	groups := make([]int, 4)

	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&group)
		groups[group-1]++
	}

	if groups[3] > 0 {
		taxi += groups[3]
	}

	if groups[2] > 0 {
		taxi += groups[2]
		if groups[0] > 0 {
			for groups[0] > 0 && groups[2] > 0 {
				groups[0]--
				groups[2]--
			}
		}
	}

	doubleTwo := groups[1] / 2
	groups[1] = groups[1] % 2
	taxi += doubleTwo
	taxi += groups[1]
	if groups[1] > 0 {
		for i := 0; i < 2 && groups[0] > 0; i++ {
			groups[0]--
		}
	}

	for groups[0] > 0 {
		taxi++
		groups[0] -= 4
	}

	fmt.Println(taxi)
}
