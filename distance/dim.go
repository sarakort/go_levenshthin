package distance

func Make2Dim(rows , columns int ) [][]int{
	// use number of row to determine the size
	// fmt.Printf("\nrow: %d, column: %d\n", rows, columns)
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
	return dim
}

func initDistancePostion(nums int) []int{
	n := make([]int, nums)
	for i := 0 ; i < nums ; i++{
		n[i] = i
	}
	return n
}

func min(n []int) int {
	m := 0
	for i, v := range n {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}