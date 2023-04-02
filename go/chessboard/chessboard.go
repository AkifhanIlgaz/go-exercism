package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if givenFile, ok := cb[file]; ok {
		count := 0
		for _, isOccupied := range givenFile {
			if isOccupied {
				count++
			}
		}
		return count
	}

	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

	if rank <= 8 && rank >= 1 {
		for _, file := range cb {
			if file[rank-1] {
				count++
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for i := 1; i <= 8; i++ {
		count += CountInRank(cb, i)
	}

	return count
}
