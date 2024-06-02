package main

import (
	"fmt"
	"sort"
)

func main() {
	/* greetings:="Hello! there, How are you today"
	   greetings=strings.ReplaceAll(greetings, "Hello", "Hey")
	   fmt.Println(strings.Contains(greetings, "Hello"))
	   fmt.Println(strings.ToTitle(greetings))
	   fmt.Println(strings.Index(greetings, "y"))
	   fmt.Println("Trim space",strings.Split(strings.TrimSpace(greetings), " "))
	   fmt.Println("Trim",strings.ReplaceAll(greetings, " ", "")) */

	// slices
	ages := []int{24, 34, 3, 45, 7, 12, 90, 87, 43, 25}
names:=[]string{"Mk", "Emmanuel", "Luigi", "Hannah", "Debra", "Kentuck"}
	sort.Ints(ages)

	sort.Strings(names)

	var index = sort.SearchInts(ages, 1000)
	fmt.Println("Ages", ages)
	fmt.Println("Index", index)
	fmt.Println("Names", names)
	fmt.Println("Sorted", sort.SearchStrings(names,"Luigi"))

} 
