package main

import "fmt"

func main() {

	// arrays
	// var ages [3]int = [3]int{1, 2, 3}

	ages := [3]int{23, 30, 9}

	var names = [3]string{"Miguel", "jamal", "Newby"}
	fmt.Println(ages, len(ages), names)

	// slices
	// can be manipulated
	var scores = []int{100, 50, 60}

	scores[0] = 58
	scores = append(scores, 87)

	scoreRange := scores[0:3]
	rangeOne := scores[:1]
	rangeTwo := scores[1:]

	fmt.Println(scores, len(scores), scoreRange)

	scoreRange = append(scoreRange, 48)

	fmt.Println(rangeOne, rangeTwo, scoreRange)
	// slice ranges

}
