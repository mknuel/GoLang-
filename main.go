package main

import "fmt"

func main() {
	// for as while loop
	/* 	x := 0

	   	for x < 5 {
	   		x++
	   		fmt.Printf("The value of x is %d\n", x)

	   	} */

	// regular for loop
	/* 	for i:=0; i<5;i++{
		fmt.Println("The value of i is",i)
	} */

	names := []string{"Mk", "Emmanuel", "Luigi", "Hannah", "Debra", "Kentuck"}

	/* for i := 0; i < len(names); i++ {

		fmt.Printf("%s is placed at index %d\n", names[i], i)

		if i==len(names)-1 {

			fmt.Printf("That list is %d names long", len(names))

		}
	} */

	/* 
	for y := 1; y <= 12; y++ {
		
		x := 0
		fmt.Println("\nThe multiplication table for", y )
		for x<=11{

			
			x++
				
			
			fmt.Printf("%d * %d = %d\n", y,x, x*y )
		}
	} */

/* 	for index, value := range names{
		fmt.Printf("The Value at index %d is %s \n", index, value)
	}
 */
	for _, value := range names{
		fmt.Printf("The Value is %s \n",  value)
	}
}
