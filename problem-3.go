package main

import ("fmt")


func getLargestPrimeFactor(number int64 ) (pfs []int64) {
	for number%2 == 0 {
		pfs = append(pfs, 2)
		number = number / 2
	}

	var i int64
	for i = 3; i*i <= number; i = i + 2 {
		// while i divides number, append i and divide number
		for number%i == 0 {
		      pfs = append(pfs, i)
		      number = number / i
		}
    }

    if number > 2 {
    	pfs = append(pfs, number)
    }

    return
}

func main() {
	//var primeNumber = 600851475143;

	fmt.Println(getLargestPrimeFactor(600851475143))
}