package main



func Define_Matrix(rows int, cols int) [][]int {
	var matrix [][]int

	matrix = make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	return matrix

}

func Get_Row(matrix [][]int, row_index int) []int {
	return matrix[row_index]
}

func Get_Column(matrix [][]int, col_index int) []int {

	col := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		col[i] = matrix[i][col_index]
	}

	return col
}

func Vector_Sum(vector_1 []int, vector_2 []int) []int {
	sum := make([]int, len(vector_1))
	for i := 0; i < len(vector_1); i++ {
		sum[i] = vector_1[i] + vector_2[i]
	}

	return sum
}

func Matrix_Sum(m1 [][]int, m2 [][]int) [][]int {

	sum := Define_Matrix(len(m1), len(m1[0]))
	for i := 0; i < len(m1); i++ {
		sum[i] = Vector_Sum(m1[i], m2[i])
	}
	return sum

}

func Inner_Product(vec_1 []int, vec_2 []int) int {

	var result int = 0

	for i := 0; i < len(vec_1); i++ {
		result += vec_1[i] * vec_2[i]
	}

	return result

}

func Matrix_Multiplication(m1 [][]int, m2 [][]int) [][]int {

	mult := Define_Matrix(len(m1), len(m1[0]))

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			mult[i][j] = Inner_Product(Get_Row(m1, i), Get_Column(m2, j))
		}
	}

	return mult
}

func Ex_1() {

	

}
