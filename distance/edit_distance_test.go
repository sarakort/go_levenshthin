package distance_test 

import (
	"dp/algorithm/distance"
	"testing"
	"gotest.tools/assert"
	"github.com/bradfitz/iter"
)


type ExpectValue struct {
	Str1 string 
	Str2 string
	Distance int
	Dimension [][]int
}


func TestMake2Dim(t *testing.T){
	const rows = 5
	const cols = 4
	result := distance.Make2Dim(rows,cols)

	expect := [][]int{
		{0,1,2,3},
		{1,0,0,0},
		{2,0,0,0},
		{3,0,0,0},
		{4,0,0,0},
	}

	for i := range iter.N(rows) {
		for j := range iter.N(cols) {
			t.Logf("[%d,%d] %d\n" ,i, j, result[i][j])
			assert.Equal(t, result[i][j], expect[i][j] )
		}
	}
}

func TestLevenshteinDistance(t *testing.T){

	valueTests := []ExpectValue{
		makeExpectValue("LAWN", "FLAW", 2 , [][]int {
			{0,1,2,3,4},
			{1,1,1,2,3},
			{2,2,2,1,2},
			{3,3,3,2,1},
			{4,4,4,3,2},
		} ),

		makeExpectValue("HONDA", "HYUDAI" , 3 , nil ),
		
		makeExpectValue("a cat", "an act", 3 , [][]int {
			{0,1,2,3,4,5,6},
			{1,0,1,2,3,4,5},
			{2,1,1,1,2,3,4},
			{3,2,2,2,2,2,3},
			{4,3,3,3,2,3,3},
			{5,4,4,4,3,3,3},
		} ),
		
		makeExpectValue("money", "monkey", 1 , [][]int {
			{0,1,2,3,4,5,6},
			{1,0,1,2,3,4,5},
			{2,1,0,1,2,3,4},
			{3,2,1,0,1,2,3},
			{4,3,2,1,1,1,2},
			{5,4,3,2,2,2,1},
		} ),

		makeExpectValue("sitting", "kitten", 3 , [][]int {
			{0,1,2,3,4,5,6},
			{1,1,2,3,4,5,6},
			{2,2,1,2,3,4,5},
			{3,3,2,1,2,3,4},
			{4,4,3,2,1,2,3},
			{5,5,4,3,2,2,3},
			{6,6,5,4,3,3,2},
			{7,7,6,5,4,4,3},
		} ),

		makeExpectValue("Sunday" ,"saturday", 3, [][]int{
		{0,1,2,3,4,5,6,7,8},
		{1,0,1,2,3,4,5,6,7},
		{2,1,1,2,2,3,4,5,6},
		{3,2,2,2,3,3,4,5,6},
		{4,3,3,3,3,4,3,4,5},
		{5,4,3,4,4,4,4,3,4},
		{6,5,4,4,5,5,5,4,3},
		}),

		makeExpectValue("RELEVANT", "ELEPHANT", 3 , [][]int {
			{0,1,2,3,4,5,6,7,8},
			{1,1,2,3,4,5,6,7,8},
			{2,1,2,2,3,4,5,6,7},
			{3,2,1,2,3,4,5,6,7},
			{4,3,2,1,2,3,4,5,6},
			{5,4,3,2,2,3,4,5,6},
			{6,5,4,3,3,3,3,4,5},
			{7,6,5,4,4,4,4,3,4},
			{8,7,6,5,5,5,5,4,3},
		} ),
		makeExpectValue("look at" ,"google", 5, [][]int{
		// {0,1,2,3,4,5,6},
		// {1,0,1,2,3,4,5},
		// {2,1,2,2,2,3,4},
		// {3,2,2,2,2,3,4},
		// {4,3,2,2,3,3,4},
		// {5,4,3,3,3,4,4},
		// {6,5,4,4,4,4,5},
		// {7,6,5,5,5,5,5},
		}),
	}

	for _ , v := range valueTests {
		levenshteinDistance(t, v)
	}
}

func makeExpectValue(str1, str2 string, distance int, dim [][]int) ExpectValue{
	return ExpectValue{
		Str1 : str1,
		Str2 : str2,
		Distance : distance,
		Dimension : dim,
	}
}

func levenshteinDistance(t *testing.T , e ExpectValue){
	v := distance.NewEditDistance(e.Str1, e.Str2)
	assert.Equal(t, v.Distance(), e.Distance)
	v.Print()
	dim := v.Dimension()
	if e.Dimension != nil && len(e.Dimension) > 0 {
		for i := range iter.N(len(e.Str1) +1) {
			for j := range iter.N(len(e.Str2)+1) {
				t.Logf("[%d,%d] %d\n" ,i, j, dim[i][j])
				assert.Equal(t, dim[i][j], e.Dimension[i][j] )
			}
		}
	}
}