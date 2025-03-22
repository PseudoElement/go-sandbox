package funcs

type CoordsDelta struct {
	x int
	y int
}

func rotate(matrix [][]int) {
	n := len(matrix)

	// Step 1: Transpose the matrix (swap matrix[i][j] with matrix[j][i])
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		// j < n/2 cause for row 0,1,2,3,4 after 2 swaps row will look like 4,3,2,1,0, then it will swap in start direction
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

/*
   yx 4*4

   top to right
   00 -> 03
   01 -> 13
   02 -> 23
   03 -> 33

   10 -> 02
   11 -> 12
   12 -> 22
   13 -> 32

   20 -> 01
   21 -> 11
   22 -> 21
   23 -> 31

   30 -> 00
   31 -> 01
   32 -> 02
   33 -> 03


*/

/*
   yx 3*3

   top to right
   00 -> 02
   01 -> 12
   02 -> 22

   right to bottom
   02 -> 22
   12 -> 21
   22 -> 20

   bottom to left
   22 -> 20
   21 -> 10
   20 -> 00

   left to top
   20 -> 00
   10 -> 01
   00 -> 02
*/
