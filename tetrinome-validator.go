package tetris

var validTetrinomes = []rune{46, 46, 46, 35, 13, 10, 46, 46, 46, 35, 13, 10, 46, 46, 46, 35, 13, 10, 46, 46, 46, 35}

func TetrinomeValidator(r []rune) bool {
	for _, v := range validTetrinomes {
		if r[0] == v {
			return true
		}
	}
	return false
}
