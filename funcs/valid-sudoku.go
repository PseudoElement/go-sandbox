package funcs

import (
	"log"
	"strconv"
	"strings"
)

func IsValidSudoku(board [][]byte) bool {
	var cells = map[string][]string{
		//
		// first row
		"00_01_02_10_11_12_20_21_22": []string{},
		"03_04_05_13_14_15_23_24_25": []string{},
		"06_07_08_16_17_18_26_27_28": []string{},
		// 2nd row
		"30_31_32_40_41_42_50_51_52": []string{},
		"33_34_35_43_44_45_53_54_55": []string{},
		"36_37_38_46_47_48_56_57_58": []string{},
		// 3th
		"60_61_62_70_71_72_80_81_82": []string{},
		"63_64_65_73_74_75_83_84_85": []string{},
		"66_67_68_76_77_78_86_87_88": []string{},
	}

	rowNums := ""
	colNums := ""

	for colIdx, _ := range board {
		for rowIdx, _ := range board {
			numInRow := board[colIdx][rowIdx]
			numInCol := board[rowIdx][colIdx]
			rowNums += string(numInRow)
			colNums += string(numInCol)

			squareKey := GetSquareKey(rowIdx, colIdx, cells)
			cells[squareKey] = append(cells[squareKey], string(numInRow))
		}
		if HasSameChar(rowNums) || HasSameChar(colNums) {
			return false
		}
		rowNums = ""
		colNums = ""
	}

	for _, value := range cells {
		if HasDuplicatesSlice(value) {
			return false
		}
	}

	return true
}

func GetSquareKey(row int, col int, m map[string][]string) string {
	colStr := strconv.Itoa(col)
	rowStr := strconv.Itoa(row)
	colRowStr := colStr + rowStr
	for key, _ := range m {
		if IsColRowKey(key, colRowStr) {
			return key
		}
	}

	return ""
}

func IsColRowKey(str string, code string) bool {
	splitted := strings.Split(str, "_")
	for _, c := range splitted {
		log.Println("c ==> ", c)
		if c == code {
			return true
		}
	}
	return false
}

func HasSameChar(str string) bool {
	m := map[string]int{}
	for _, runa := range str {
		if string(runa) != "." {
			prev, _ := m[string(runa)]
			m[string(runa)] = prev + 1
		}
	}

	for _, value := range m {
		if value > 1 {
			return true
		}
	}

	return false
}

func HasDuplicatesSlice(nums []string) bool {
	seen := make(map[string]bool)

	for _, char := range nums {
		if seen[string(char)] && string(char) != "." {
			return true
		}
		seen[string(char)] = true
	}

	return false
}
