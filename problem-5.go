package main

import ("fmt")

func main() {
	var retval = 1;

	for i := 2; i <= 20; i++ {
		retval = LCM(retval, i)
	}

	fmt.Println(retval)
}

func LCM(num, counter int ) int {
	return num * counter / GCD(num, counter)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
  }

  return a
}