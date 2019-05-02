package main

import (
	"strings"
	"bufio"
	"os"
	"os/signal"
	"syscall"
	"github.com/fatih/color"
	"fmt"
	"dp/algorithm/distance"
)

var (
	green = color.New(color.FgGreen, color.Bold)
	blue = color.New(color.FgBlue, color.BgHiWhite)
)


func display(word1, word2 string){
	ed := distance.NewEditDistance(word1,word2)
	fmt.Println("")
	distance := ed.Distance()
	ed.Print()
	fmt.Println("")
	green.Printf("Distance = %d \n", distance)
}


func main(){
	green.Println("The Levenshtein distance")
	fmt.Println("string metric for measuring the difference between two sequences.")
	fmt.Println("https://en.wikipedia.org/wiki/Levenshtein_distance")

	word1 := "สบายดี"
	word2 := "สวัสดี"
	green.Printf(" Compare 2 sequences %s and %s .\n", word1, word2)
	display( word1, word2)

	c := make(chan os.Signal , 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<- c
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		blue.Print("->")
		text1  := readString(reader)
		blue.Print("->")
		text2  := readString(reader)
		fmt.Println("-----------------")
		display(text1, text2)
	}
}

func readString(reader *bufio.Reader)  string {
		text , _ := reader.ReadString('\n')
		// convert CRLF to LF
		return strings.Replace(text, "\n", "",-1)
}