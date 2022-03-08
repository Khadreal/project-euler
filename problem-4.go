package main

import (
	"fmt"
	"strconv"
)


func checkIfNumberIsPalindrome(number int64) bool {
	if number < 0 {
		number = -number
	}

	var s = strconv.FormatInt(number, 10 )

	bound := (len(s) / 2) + 1
    for i := 0; i < bound; i++ {
    	if s[i] != s[ len(s) - 1 - i ] {
        	return false
        }
    }
    
    return true
}


func main() {
	var max, maxOne, maxTwo = 0, 0, 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j

	        if max > product {
	        	break
	        }

	        if checkIfNumberIsPalindrome( int64(product) ) {
            	max = product
            	maxOne = i
            	maxTwo = j

            	continue
          	}
		}
	}

	retval := fmt.Sprintf("A product of two 3-digit is %d x %d = %d \n", maxOne, maxTwo, max)

	fmt.Print(retval)
}