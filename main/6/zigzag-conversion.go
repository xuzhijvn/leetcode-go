package main

import "fmt"

/**
时间复杂度：O(n)
*/
func convert(s string, numRows int) string {
	//step1: 参数检查
	if s == "" || numRows <= 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	//step2: 构造Z字形二维矩阵
	matrix := make([][]string, numRows)
	row := 0
	col := 1
	for i := 0; i < len(s); i++ {
		if row < numRows {
			mod := (col - 1) % (numRows - 1)
			if mod != 0 {
				row = numRows - mod - 1
				matrix[row] = append(matrix[row], string(s[i]))
				col++
				row = 0
			} else {
				matrix[row] = append(matrix[row], string(s[i]))
				row++
			}
		} else {
			col++
			row = 0
			i--
		}
	}
	//fmt.Println(matrix)

	//step3: 遍历矩阵
	res := ""
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res += matrix[i][j]
		}
	}

	return res
}

func main() {
	matrix := make([][]string, 4)
	matrix[0] = append(matrix[0], "e")
	matrix[0] = append(matrix[0], "t")
	fmt.Println(matrix)
	fmt.Println(convert("LEETCODEISHIRING", 3))

}
