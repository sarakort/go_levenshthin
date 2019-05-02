package distance 

import (
	"strings"
	"unicode"
	"fmt"
)

type EditDistance struct{
	StrFirst string
	StrSecond string
	dimension [][]int
	init bool
}

type CalcDistance interface {
	Dimension()  [][]int
	Distance()  int
	MatchCost(r1 , r2 rune) int 
	MinCost(i ,j , cost int) int
	Print()
}

func (e EditDistance) Make2Dim(rows , columns int ) {
	dim := make([][]int, rows) 
	firstRow := 0
	firstColumn := 0

	// folloywed by columns
	for i := range dim {
		if i == firstRow {
			dim[i] = initDistancePostion(columns)
		}else {
			temp := make([]int, columns)
			temp[firstColumn] = i
			dim[i] = temp
		}
	}
	e.init = true
	e.dimension = dim 
}

func (e EditDistance) Dimension() [][]int {
	tmp := make([][]int, len(e.dimension) )
	copy ( tmp, e.dimension)
	return   tmp
}

func (e EditDistance) MinCost(i,j, cost int) int {
	// fmt.Printf("i: %d, j: %d , edit: %d", i,j, edit)
	v := []int {
		e.dimension[i][j+1] +1 , // delete / insert from str1
		e.dimension[i+1][j] +1,  // delete / insert from str2
		e.dimension[i][j] + cost, // delete /insert from both
	}
	return min(v)
}

func (e EditDistance) MatchCost(r1, r2 rune) int{
	// fmt.Printf("%s(%d),%s(%d) \n" , string(r1), unicode.ToLower(r1), string(r2), unicode.ToLower(r2))
	if unicode.ToLower(r1) == unicode.ToLower(r2) {
		return 0
	}
	return 1
}

func (e EditDistance) Distance() int {
	str1Len := len(e.StrFirst)
	str2Len := len(e.StrSecond)
	if !e.init {
		e.dimension = Make2Dim(str1Len +1 , str2Len +2)
	}
	r1 := []rune(strings.ToLower(e.StrFirst))
	r2 := []rune(strings.ToLower(e.StrSecond))
	fmt.Println(r1)
	fmt.Println(r2)
	for i := 0 ; i < str1Len ; i++{
		for j:= 0 ; j < str2Len ; j++ {
			cost := e.MatchCost(r1[i], r2[j])
			e.dimension[i+1][j+1] = e.MinCost(i, j , cost)
		}
	}
	// fmt.Println(e.dimension)

	return e.dimension[str1Len][str2Len]
}

func (e EditDistance) Print() {
	for i:= range e.dimension {
		fmt.Println(e.dimension[i])
	}
}

func NewEditDistance(str1, str2 string) EditDistance{
	v := EditDistance{
		StrFirst : strings.TrimSpace(str1),
		StrSecond : strings.TrimSpace(str2),
		dimension : Make2Dim(len(str1) +1, len(str2)+ 1),
		init : true,
	}
	return v
}

