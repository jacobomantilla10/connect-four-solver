package game

import "testing"

func TestMakeBoard(t *testing.T) {
	board := MakeBoard()
	if board.w != 6 || board.h != 7 {
		t.Errorf("Board didn't create correctly expected dimensions 6 x 7, go %d x %d", board.w, board.h)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if board.position != 0 {
				t.Errorf("Board didn't initialize as empty")
			}
		}
	}
	if board.numMoves != 0 {
		t.Errorf("Board was created with wrong number of starting moves")
	}
}

func TestCanPlay(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  bool
	}{
		{"8 is out of bounds", 8, false},
		{"0 is out of bounds", 0, false},
		{"1 is a valid insert", 1, true},
		{"7 is a valid insert", 7, true},
		{"9065 is out of bounds", 9065, false},
		{"-2 is out of bounds", -2, false},
		{"3 is full so we can't play", 3, false},
	}

	board := MakeBoard()
	for i := 0; i < 6; i++ {
		board.Play(2)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := board.CanPlay(tt.input - 1)
			if ans != tt.want {
				t.Errorf("got %t expected %t", ans, tt.want)
			}
		})
	}
}

func TestTopMask(t *testing.T) {
	var tests = []struct {
		name string
		col  int
		want uint64
	}{
		{"First column top mask", 1, 0},
		{"Second column top mask", 2, 0},
		{"Third column top mask", 3, 0},
		{"Fourth column top mask", 4, 0},
		{"Fifth column top mask", 5, 0},
		{"Sixth column top mask", 6, 0},
		{"Seventh column top mask", 7, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := top_mask(tt.col)
			if ans != tt.want {
				t.Errorf("got %b expected %b", ans, tt.want)
			}
		})
	}
}

func TestIsWinningMove(t *testing.T) {
	var tests = []struct {
		name        string
		col         int
		boardString string
		want        bool
	}{
		{"3 is a winning move", 3, "445566", true},
		{"6 is a winning move", 6, "445533", true},
		{"2 is not a winning move", 2, "4455", false},
		{"3 is also winning move", 3, "444555542216", true},
		{"3 is another winning move", 3, "444555542213", true},
		{"3 is not a winning move", 3, "44455554221", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := MakeBoardFromString(tt.boardString)
			ans := board.IsWinningMove(tt.col - 1)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestNumMoves(t *testing.T) {
	var tests = []struct {
		name        string
		boardString string
	}{
		{"This board has 39 moves", "3135151421347443544172316522225776773566"},
		{"This board has 35 moves", "11513154335437461536136677566452744"},
		{"This board has 15 moves", "641335277532624"},
		{"This board has 4 moves", "4737"},
		{"This board has 0 moves", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := MakeBoardFromString(tt.boardString)
			if board.NumMoves() != len(tt.boardString) {
				t.Errorf("got %d want %d", board.NumMoves(), len(tt.boardString))
			}
		})
	}
}

func TestBottomBoardMask(t *testing.T) {
	t.Run("Bottom Mask test", func(t *testing.T) {
		want := 0b0000001000000100000010000001000000100000010000001
		res := bottom_board_mask(7, 6)
		if res != uint64(want) {
			t.Errorf("got %b want %b", res, want)
		}
	})
}

// func TestMakeBoardFromString(t *testing.T) {
// 	//Set up
// 	randBoardMatrix := [6][7]rune{}
// 	randBoardMatrix[0] = [7]rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}
// 	randBoardMatrix[1] = [7]rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}
// 	randBoardMatrix[2] = [7]rune{' ', ' ', ' ', 'O', 'X', ' ', ' '}
// 	randBoardMatrix[3] = [7]rune{' ', ' ', ' ', 'X', 'O', ' ', ' '}
// 	randBoardMatrix[4] = [7]rune{' ', 'O', ' ', 'O', 'X', ' ', ' '}
// 	randBoardMatrix[5] = [7]rune{'X', 'X', ' ', 'X', 'O', ' ', ' '}

// 	fullBoardMatrix := [6][7]rune{}
// 	fullBoardMatrix[0] = [7]rune{'O', 'X', 'O', 'X', 'O', 'X', 'O'}
// 	fullBoardMatrix[1] = [7]rune{'X', 'O', 'X', 'O', 'X', 'O', 'X'}
// 	fullBoardMatrix[2] = [7]rune{'O', 'X', 'O', 'X', 'O', 'X', 'O'}
// 	fullBoardMatrix[3] = [7]rune{'X', 'O', 'X', 'O', 'X', 'O', 'X'}
// 	fullBoardMatrix[4] = [7]rune{'O', 'X', 'O', 'X', 'O', 'X', 'O'}
// 	fullBoardMatrix[5] = [7]rune{'X', 'O', 'X', 'O', 'X', 'O', 'X'}

// 	emptyBoard := MakeBoard()
// 	randBoard := MakeBoardWithMatrix(randBoardMatrix)
// 	fullBoard := MakeBoardWithMatrix(fullBoardMatrix)
// 	var tests = []struct {
// 		name  string
// 		input string
// 		want  Board
// 	}{
// 		{"Empty board should be equal", "", emptyBoard},
// 		{"Random board should be equal", "44455554221", randBoard},
// 		{"Full board should be equal", "123456712345671234567123456712345671234567", fullBoard},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ans, _ := MakeBoardFromString(tt.input)
// 			if ans.data != tt.want.data {
// 				t.Errorf("got %v wanted %v", ans.data, tt.want.data)
// 			}
// 		})
// 	}
// }
