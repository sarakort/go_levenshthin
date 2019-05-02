package main

import (
	"fmt"
	"dp/algorithm/distance"
)



func main(){
	var word1 = "abc"  
	var word2 = "add"
	ed := distance.NewEditDistance(word1,word2)

	fmt.Printf("compare 2 word %s and %s distance: %d \n", word1, word2 , ed.Distance())
	// fmt.Printf("%v\n", ed.Dimension())
	ed.Print()
}