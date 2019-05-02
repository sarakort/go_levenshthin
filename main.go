package main

import (
	"github.com/fatih/color"
	"fmt"
	"dp/algorithm/distance"
)



func main(){
	var word1 = "sunday"  
	var word2 = "Saturday"
	ed := distance.NewEditDistance(word1,word2)
	distance := ed.Distance()

	green := color.New(color.FgGreen, color.Bold)
	green.Println("The Levenshtein distance")
	fmt.Println("string metric for measuring the difference between two sequences.")
	fmt.Println("https://en.wikipedia.org/wiki/Levenshtein_distance")
	fmt.Println("")
	fmt.Printf("Compare 2 sequences %s and %s .\n", word1, word2)
	ed.Print()
	fmt.Println("")
	green.Printf("Distance = %d \n", distance)
}