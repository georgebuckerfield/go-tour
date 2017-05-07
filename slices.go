package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8,dy)

	for i := range s {
		s[i] = make([]uint8, dx)

		for ii := range s[i] {
			s[i][ii] = uint8(i+ii*3)
		}
	}


	fmt.Println(len(s))
	return s
}

func main() {
	pic.Show(Pic)
}
