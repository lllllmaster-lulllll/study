package test

import "fmt"

func FallThrough() {
	var i int = 1
	for i <= 32 {
		switch i {
		case 1:
			fmt.Printf("i = %d\n", i)
		case 2:
			fmt.Printf("i = %d\n", i)
		case 4:
			fmt.Printf("i = %d\n", i)
		case 8:
			fmt.Printf("i = %d\n", i)
		case 16:
			fmt.Printf("i = %d\n", i)
		case 32:
			fmt.Printf("i = %d\n", i)
		default:

		}
		i = i * 2
	}
}

func sliceCut() {
	s1 := []int{
		1, 2, 3, 4, 5,
	}
	for index, s := range s1 {
		fmt.Printf("s1_%d = %d\n", index, s)
	}

	s2 := s1[:len(s1)-1]
	for index, s := range s2 {
		fmt.Printf("s2_%d = %d\n", index, s)
	}
}
