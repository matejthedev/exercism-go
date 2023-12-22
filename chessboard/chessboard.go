package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, v := range cb[file] {
		if v {
			count++
		}
	}
	return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank > 8 || rank < 1 {
		return 0
	}
	count := 0
	for _, v := range cb {
		if v[rank-1] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		count += len(v)
	}
	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, v := range file {
			if v {
				count++
			}
		}
	}
	return count
}
