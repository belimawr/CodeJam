package main

import "fmt"

type rope struct {
	a int
	b int
}

func main() {
	var cases int
	fmt.Scanf("%d", &cases)

	for i := 0; i < cases; i++ {
		var n int
		fmt.Scanf("%d", &n)
		ropes := []rope{}

		for j := 0; j < n; j++ {
			var a, b int
			fmt.Scanf("%d %d", &a, &b)
			ropes = append(ropes, rope{a, b})
		}

		count := 0
		for a := 0; a < len(ropes); a++ {
			for b := a; b < len(ropes); b++ {
				if (ropes[a].a-ropes[b].a)*(ropes[a].b-ropes[b].b) < 0 {
					count++
				}
			}
		}
		fmt.Printf("Case #%d: %d\n", i+1, count)
	}
}
