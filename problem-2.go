package main

import ("fmt")

func main() {
	var a, b, i, sum uint64 = 1, 2, 3, 2
	for i < 4000000 {
		next := i 

		fmt.Println(next)

		if next%2 == 0 {
			sum += next 
		}
		a = b
		b = next
		i = a+b
	}

	fmt.Println(sum, "is the sum of the even-valued terms below 4,000,000")
}