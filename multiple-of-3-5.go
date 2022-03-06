package main

import ("fmt")

func main(){
	sum := 0;
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
			fmt.Println(i, " is multiple of 3")
			continue 
		}
		if i%5 == 0 {
			sum += i
			fmt.Println(i, " is multiple of 5")
			continue
		}
	}

	fmt.Println("The sum of multiple of 3 or 5 below 1000 is ", sum)
}