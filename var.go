package main

import "fmt"


func main(){

	//string

	var text string = "my name is mk";
	var textTwo = "janie"
	var textThree string;


	fmt.Println(text, textTwo, textThree)
	
	textFour := "john"
	
	fmt.Println(textFour)

	//int
	var intOne = 23
	var intTwo int = 24
	intThree := 21

	fmt.Println(intOne, intTwo, intThree)

	//bits and memory
	var intFour int8 = 10
	var intFive int8 = -123
	var intSix uint16= 256

	// floats
	var scoreOne float32= 45.92938
	var scoreTwo float64= 445463728292025.92938
	scoreThree := 292025.92938
	
	fmt.Println(intFour, intFive, intSix)
	fmt.Println(scoreOne, scoreTwo, scoreThree)



}

