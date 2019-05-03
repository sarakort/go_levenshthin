package distance 

import (
	"github.com/bradfitz/iter"
	"strings"
	"fmt"
	"github.com/fatih/color"
)

type EditDistance struct{
	StrFirst string
	StrSecond string
	dimension [][]int
}

type CalcDistance interface {
	Dimension()  [][]int
	Distance()  int
	MatchCost(r1 , r2 rune) int 
	MinCost(i ,j , cost int) int
	Print()
}

func (e *EditDistance) Dimension() [][]int {
	tmp := make([][]int, len(e.dimension) )
	copy ( tmp, e.dimension)
	return   tmp
}

func (e *EditDistance) MinCost(i,j, cost int) int {
	// fmt.Printf("i: %d, j: %d , cost: %d\n", i,j, cost)
	v := []int {
		e.dimension[i][j+1] +1 , // delete / insert from str1
		e.dimension[i+1][j] +1,  // delete / insert from str2
		e.dimension[i][j] + cost, // delete /insert from both
	}
	return min(v)
}

func (e EditDistance) MatchCost(r1, r2 rune) int{
	// fmt.Printf("%s(%d),%s(%d) \n" , string(r1), unicode.ToLower(r1), string(r2), unicode.ToLower(r2))
	if r1 == r2 {
		return 0
	}
	return 1
}

func (e *EditDistance) Distance() int {
	r1 := []rune(strings.ToLower(e.StrFirst))
	r2 := []rune(strings.ToLower(e.StrSecond))
	str1Len := len(r1)
	str2Len := len(r2)
	e.dimension = Make2Dim(str1Len +1 , str2Len +1)

	// fmt.Println(r1)
	// fmt.Println(r2)
	for i := 0 ; i < str1Len ; i++{
		for j:= 0 ; j < str2Len ; j++ {
			cost := e.MatchCost(r1[i], r2[j])
			e.dimension[i+1][j+1] = e.MinCost(i, j , cost)
		}
	}
	// fmt.Println(e.dimension)
	return e.dimension[str1Len][str2Len]
}

func stringToArray(str string) []string{
	return strings.Split(str, "")
}

func (e *EditDistance) Print() {
	
	var (
		strCols = stringToArray(" "+e.StrSecond)
		strRows = stringToArray(" "+e.StrFirst)
		blue = color.New(color.FgBlue , color.Bold)	
		green = color.New(color.FgGreen)
		red = color.New(color.FgRed, color.Bold).Add(color.BgHiWhite)
		write = color.New(color.FgWhite)
	)

	fmt.Print(" ")
	blue.Printf("%v\n",strCols)

	dim := e.dimension
	for i:= range iter.N(len(strRows)) {
		// first chr is blue color
		blue.Print(strRows[i])


		// first element in array is green color
		if( i == 0 ){
			green.Printf("%v\n", dim[i])
			continue
		}

		// set color when corner right bottom martix is red color orther write color
		var printcolor *color.Color = write
		if i == len(dim)-1{
			printcolor = red
		}

		arrcol := dim[i]
		for j := range arrcol{
			v := arrcol[j]
			switch j {
			case 0 : 
				green.Printf("[%d ", v)
			case len(arrcol) -1 :
				printcolor.Printf("%d", v)
				green.Println("]") 
			default : 
				write.Printf("%d ",v)
			}
		}
		
	}
}

func Trim(str string)  string{
	return strings.TrimSpace(str)
}

func NewEditDistance(str1, str2 string) EditDistance{
	str1 = Trim(str1)
	str2 = Trim(str2)
	v := EditDistance{
		StrFirst : str1,
		StrSecond : str2,
	}
	return v
}

