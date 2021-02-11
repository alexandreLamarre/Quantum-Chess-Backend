package quantumchess

import (
	"log"
	"math"
	"testing"
)

//debug toggles whether or not to log some debug messages for the tests
var debug bool = true

func TestAllInfluence(t *testing.T) {
	t.Log("Testing area of influence")
	debug = false
	DEBUG_QUANTUM_CHESS_STRUCTS = false
	testBishopAof(t)
	testKnightAoF(t)
	testRookAoF(t)
	testQueenAoF(t)
	testKingAoF(t)
}

func createBoard(t *testing.T, board *Board, positions []int) {
	if len(positions) != 64 {
		t.Errorf("Warning:")
	}
	board.Positions = positions

	for i, v := range board.Positions {
		if positions[i] != v {
			t.Errorf("Expected board positions %d to "+
				"be %d, got : %d", i, positions[i], v)
		}
	}
}

func testBoardGetID(t *testing.T, board *Board, id int, expected int) {
	if board.getID(id) != expected {
		t.Errorf("Expected board %d to be %d instead got %d", id, expected, board.getID(id))
	}
}

func createPiece(t *testing.T, piece *Piece, action string, color int, initstate map[string][2]float64,
	statespace []string, state map[string][2]float64, moved bool) {
	piece.Action = action
	piece.Color = color
	piece.InitialState = initstate
	if len(piece.InitialState) == 0 {
		t.Errorf("Qpiece needs to have initial states")
	}
	piece.StateSpace = statespace
	if len(piece.StateSpace) == 0 {
		t.Errorf("Qpiece needs to have state space")
	}
	piece.State = state
	if len(piece.State) == 0 {
		t.Errorf("Qpiece needs to have a state")
	}
	piece.Moved = moved
}

func testAoF(t *testing.T, piece *Piece, pos int, board *Board, expected map[int]bool) {
	AoF, err := piece.getAreaOfInfluence(board, pos)
	if err != nil {
		t.Logf("AoF error")
		log.Fatal(err)
	}
	if debug {
		t.Logf("length of AoF %d", len(AoF))
		for id := range AoF {
			t.Log(id)
		}
	}

	if len(expected) != len(AoF) {
		t.Errorf(
			"Incorrect amount of pieces in aof: %d, expected %d", len(AoF), len(expected))
	}

	for id := range expected {
		if AoF[id] != true {
			t.Errorf("Expected %d to be in Area of Influence but was not", id)
		}
	}
}

func testBishopAof(t *testing.T) {
	positions := [64]int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 2, 0, 0,
		0, 0, 0, 0, 5, 0, 0, 0, // test bishop at position id = 5, i.e. pos = 20
		0, 0, 0, 3, 0, 4, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 13, 27, 29
	board := &Board{}
	createBoard(t, board, positions[:])

	qBishop := &Piece{}
	stateSpace := make([]string, 0, 0)
	stateSpace = append(stateSpace, "Bishop")
	stateSpace = append(stateSpace, "Pawn")
	bishopInitialState := make(map[string][2]float64)
	bishopInitialState["Bishop"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qBishop, "Hadamard", 1,
		bishopInitialState, stateSpace, bishopInitialState, false)
	res := map[int]bool{
		13: true,
		27: true,
		29: true,
	}
	testAoF(t, qBishop, 20, board, res)

	positions1 := [64]int{
		0, 0, 2, 0, 0, 0, 3, 0,
		0, 0, 0, 1, 4, 0, 0, 0,
		0, 0, 0, 0, 5, 0, 0, 0, // test bishop at position id = 5, i.e. pos = 20
		0, 0, 0, 0, 9, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 7, 0,
		0, 0, 0, 0, 0, 0, 0, 8,
		6, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 6, 10, 12, 38, 48
	board1 := &Board{}
	createBoard(t, board1, positions1[:])

	qBishop1 := &Piece{}
	stateSpace1 := make([]string, 0, 0)
	stateSpace1 = append(stateSpace, "Bishop")
	stateSpace1 = append(stateSpace, "Pawn")
	bishopInitialState1 := make(map[string][2]float64)
	bishopInitialState1["Bishop"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState1["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qBishop1, "Hadamard", 1,
		bishopInitialState1, stateSpace1, bishopInitialState1, false)
	res1 := map[int]bool{
		6:  true,
		11: true,
		28: true,
		38: true,
		48: true,
	}
	testAoF(t, qBishop1, 20, board1, res1)

	positions2 := [64]int{
		0, 0, 2, 0, 0, 0, 3, 0,
		0, 0, 0, 1, 4, 0, 0, 0,
		0, 0, 0, 0, 5, 0, 0, 0, // test bishop at position id = 5, i.e. pos = 20
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 7, 0,
		0, 0, 0, 0, 0, 0, 0, 8,
		6, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 6, 10, 12, 38, 48
	board2 := &Board{}
	createBoard(t, board2, positions2[:])

	qBishop2 := &Piece{}
	stateSpace2 := make([]string, 0, 0)
	stateSpace2 = append(stateSpace, "Bishop")
	stateSpace2 = append(stateSpace, "Pawn")
	bishopInitialState2 := make(map[string][2]float64)
	bishopInitialState2["Bishop"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState2["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qBishop2, "Hadamard", 0,
		bishopInitialState2, stateSpace2, bishopInitialState2, false)
	res2 := map[int]bool{
		6:  true,
		11: true,
		12: true,
		38: true,
		48: true,
	}
	testAoF(t, qBishop2, 20, board2, res2)

	positions3 := [64]int{
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 0, 2, 2, 2, 2,
		2, 2, 0, 2, 2, 2, 2, 2,
		0, 0, 2, 2, 2, 2, 2, 2,
		1, 2, 2, 2, 2, 2, 2, 2, // test bishop at position id = 1, i.e. pos = 32
		0, 0, 2, 2, 2, 2, 2, 2,
		2, 2, 0, 2, 2, 2, 2, 2,
		2, 2, 2, 0, 2, 2, 2, 2,
	} //squares to be expected NONE
	board3 := &Board{}
	createBoard(t, board3, positions3[:])

	qBishop3 := &Piece{}
	stateSpace3 := make([]string, 0, 0)
	stateSpace3 = append(stateSpace, "Bishop")
	stateSpace3 = append(stateSpace, "Pawn")
	bishopInitialState3 := make(map[string][2]float64)
	bishopInitialState3["Bishop"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState3["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qBishop3, "Hadamard", 0,
		bishopInitialState3, stateSpace3, bishopInitialState3, false)
	res3 := map[int]bool{}
	testAoF(t, qBishop3, 32, board3, res3)
}

func testKnightAoF(t *testing.T) {
	positions := [64]int{
		0, 0, 0, 1, 0, 2, 0, 0,
		0, 0, 3, 0, 0, 0, 6, 0,
		0, 0, 0, 0, 5, 0, 0, 0, // test knight at position id = 5, i.e. pos = 20
		0, 0, 4, 0, 2, 0, 7, 0,
		0, 0, 0, 8, 0, 9, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 3,5,10,14,26,28,30,35,37
	board := &Board{}
	createBoard(t, board, positions[:])

	qKnight := &Piece{}
	stateSpace := make([]string, 0, 0)
	stateSpace = append(stateSpace, "Knight")
	stateSpace = append(stateSpace, "Pawn")
	bishopInitialState := make(map[string][2]float64)
	bishopInitialState["Knight"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKnight, "Hadamard", 1,
		bishopInitialState, stateSpace, bishopInitialState, false)
	res := map[int]bool{
		3:  true,
		5:  true,
		10: true,
		14: true,
		26: true,
		28: true,
		30: true,
		35: true,
		37: true,
	}
	testAoF(t, qKnight, 20, board, res)

	positions1 := [64]int{
		0, 0, 2, 0, 0, 2, 0, 0,
		5, 0, 0, 0, 0, 0, 6, 0,
		0, 0, 3, 0, 5, 0, 0, 0, // test knight at position id = 5, i.e. pos = 20
		0, 4, 4, 0, 2, 0, 7, 0,
		0, 0, 0, 8, 0, 9, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 2, 18, 25
	board1 := &Board{}
	createBoard(t, board1, positions1[:])

	qKnight1 := &Piece{}
	stateSpace1 := make([]string, 0, 0)
	stateSpace1 = append(stateSpace, "Knight")
	stateSpace1 = append(stateSpace, "Pawn")
	bishopInitialState1 := make(map[string][2]float64)
	bishopInitialState1["Knight"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState1["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKnight1, "Hadamard", 1,
		bishopInitialState1, stateSpace1, bishopInitialState1, false)
	res1 := map[int]bool{
		2:  true,
		18: true,
		25: true,
	}
	testAoF(t, qKnight1, 8, board1, res1)

	positions2 := [64]int{
		2, 2, 2, 0, 2, 0, 2, 2,
		2, 2, 0, 2, 0, 2, 0, 2,
		2, 2, 2, 2, 5, 2, 2, 2, // test knight at position id = 5, i.e. pos = 20
		2, 2, 0, 2, 0, 2, 0, 2,
		2, 2, 2, 0, 2, 0, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
	} //squares to be expected NONE
	board2 := &Board{}
	createBoard(t, board2, positions2[:])

	qKnight2 := &Piece{}
	stateSpace2 := make([]string, 0, 0)
	stateSpace2 = append(stateSpace, "Knight")
	stateSpace2 = append(stateSpace, "Pawn")
	bishopInitialState2 := make(map[string][2]float64)
	bishopInitialState2["Knight"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState2["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKnight2, "Hadamard", 1,
		bishopInitialState2, stateSpace2, bishopInitialState2, false)
	res2 := map[int]bool{}
	testAoF(t, qKnight, 20, board2, res2)
}

func testRookAoF(t *testing.T) {
	positions := [64]int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 2, 0, 6, 0,
		0, 0, 0, 2, 5, 0, 0, 2, // test knight at position id = 5, i.e. pos = 20
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 2, 0, 0, 0,
		0, 0, 0, 0, 2, 0, 0, 0,
	} //squares to be expected 12,19,23,52
	board := &Board{}
	createBoard(t, board, positions[:])

	qRook := &Piece{}
	stateSpace := make([]string, 0, 0)
	stateSpace = append(stateSpace, "Rook")
	stateSpace = append(stateSpace, "Pawn")
	bishopInitialState := make(map[string][2]float64)
	bishopInitialState["Rook"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qRook, "Measurement", 1,
		bishopInitialState, stateSpace, bishopInitialState, false)
	res := map[int]bool{
		12: true,
		19: true,
		23: true,
		52: true,
	}
	testAoF(t, qRook, 20, board, res)

	positions1 := [64]int{
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		0, 0, 0, 0, 5, 0, 0, 0, // test knight at position id = 5, i.e. pos = 20
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
	} //squares to be expected NONE
	board1 := &Board{}
	createBoard(t, board1, positions1[:])

	qRook1 := &Piece{}
	stateSpace1 := make([]string, 0, 0)
	stateSpace1 = append(stateSpace, "Rook")
	stateSpace1 = append(stateSpace, "Pawn")
	bishopInitialState1 := make(map[string][2]float64)
	bishopInitialState1["Rook"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState1["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qRook1, "Measurement", 1,
		bishopInitialState1, stateSpace1, bishopInitialState1, false)
	res1 := map[int]bool{}
	testAoF(t, qRook1, 20, board1, res1)

	positions2 := [64]int{
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		5, 0, 0, 0, 0, 0, 0, 0, // test knight at position id = 5, i.e. pos = 20
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
	} //squares to be expected NONE
	board2 := &Board{}
	createBoard(t, board2, positions2[:])

	qRook2 := &Piece{}
	stateSpace2 := make([]string, 0, 0)
	stateSpace2 = append(stateSpace, "Rook")
	stateSpace2 = append(stateSpace, "Pawn")
	bishopInitialState2 := make(map[string][2]float64)
	bishopInitialState2["Rook"] = [2]float64{math.Sqrt(2), 0}
	bishopInitialState2["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qRook2, "Measurement", 1,
		bishopInitialState2, stateSpace2, bishopInitialState2, false)
	res2 := map[int]bool{
		8:  true,
		24: true,
	}
	testAoF(t, qRook2, 16, board2, res2)
}

func testQueenAoF(t *testing.T) {
	positions := [64]int{
		0, 0, 2, 0, 2, 0, 0, 0,
		0, 0, 0, 0, 0, 2, 0, 0,
		2, 0, 0, 0, 5, 0, 0, 2, // test queen at position id = 5, i.e. pos = 20
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 2, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		2, 0, 0, 0, 2, 0, 0, 0,
		0, 0, 0, 0, 2, 0, 0, 0,
	} //squares to be expected 2, 4, 13, 16, 23, 38, 48, 52
	board := &Board{}
	createBoard(t, board, positions[:])

	qQueen := &Piece{}
	stateSpace := make([]string, 0, 0)
	stateSpace = append(stateSpace, "Queen")
	stateSpace = append(stateSpace, "Pawn")
	queenInitialState := make(map[string][2]float64)
	queenInitialState["Queen"] = [2]float64{math.Sqrt(2), 0}
	queenInitialState["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qQueen, "Measurement", 1,
		queenInitialState, stateSpace, queenInitialState, false)
	res := map[int]bool{
		2:  true,
		4:  true,
		13: true,
		16: true,
		23: true,
		38: true,
		48: true,
		52: true,
	}
	testAoF(t, qQueen, 20, board, res)

	positions1 := [64]int{
		2, 2, 0, 2, 0, 2, 0, 2,
		2, 2, 2, 0, 0, 0, 2, 2,
		0, 0, 0, 0, 5, 0, 0, 0, // test queen at position id = 5, i.e. pos = 20
		2, 2, 2, 0, 0, 0, 2, 2,
		2, 2, 0, 2, 0, 2, 0, 2,
		2, 0, 2, 2, 0, 2, 2, 0,
		0, 2, 2, 2, 0, 2, 2, 2,
		2, 2, 2, 2, 0, 2, 2, 2,
	} //squares to be expected NONE
	board1 := &Board{}
	createBoard(t, board1, positions1[:])

	qQueen1 := &Piece{}
	stateSpace1 := make([]string, 0, 0)
	stateSpace1 = append(stateSpace, "Queen")
	stateSpace1 = append(stateSpace, "Pawn")
	queenInitialState1 := make(map[string][2]float64)
	queenInitialState1["Queen"] = [2]float64{math.Sqrt(2), 0}
	queenInitialState1["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qQueen1, "Measurement", 1,
		queenInitialState1, stateSpace1, queenInitialState1, false)
	res1 := map[int]bool{}
	testAoF(t, qQueen, 20, board1, res1)
}

func testKingAoF(t *testing.T) {
	positions := [64]int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 2, 2, 2, 0, 0,
		0, 0, 0, 2, 5, 2, 0, 0, // test king at position id = 5, i.e. pos = 20
		0, 0, 0, 2, 2, 2, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	} //squares to be expected 11,12,13,19,21,27,28,29
	board := &Board{}
	createBoard(t, board, positions[:])

	qKing := &Piece{}
	stateSpace := make([]string, 0, 0)
	stateSpace = append(stateSpace, "King")
	stateSpace = append(stateSpace, "Pawn")
	kingInitialState := make(map[string][2]float64)
	kingInitialState["King"] = [2]float64{math.Sqrt(2), 0}
	kingInitialState["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKing, "Measurement", 1,
		kingInitialState, stateSpace, kingInitialState, false)
	res := map[int]bool{
		11: true,
		12: true,
		13: true,
		19: true,
		21: true,
		27: true,
		28: true,
		29: true,
	}
	testAoF(t, qKing, 20, board, res)

	positions1 := [64]int{
		2, 2, 2, 2, 2, 0, 0, 0,
		2, 2, 2, 0, 0, 0, 2, 2,
		2, 2, 2, 0, 5, 0, 2, 2, // test king at position id = 5, i.e. pos = 20
		2, 2, 2, 0, 0, 0, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
	} //squares to be expected NONE
	board1 := &Board{}
	createBoard(t, board1, positions1[:])

	qKing1 := &Piece{}
	stateSpace1 := make([]string, 0, 0)
	stateSpace1 = append(stateSpace, "King")
	stateSpace1 = append(stateSpace, "Pawn")
	kingInitialState1 := make(map[string][2]float64)
	kingInitialState1["King"] = [2]float64{math.Sqrt(2), 0}
	kingInitialState1["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKing1, "Measurement", 1,
		kingInitialState1, stateSpace1, kingInitialState1, false)
	res1 := map[int]bool{}
	testAoF(t, qKing, 20, board1, res1)

	positions2 := [64]int{
		5, 2, 2, 2, 2, 0, 0, 0,
		2, 2, 2, 0, 0, 0, 2, 2,
		2, 2, 2, 0, 0, 0, 2, 2, // test king at position id = 5, i.e. pos = 0
		2, 2, 2, 0, 0, 0, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
	} //squares to be expected NONE
	board2 := &Board{}
	createBoard(t, board2, positions2[:])

	qKing2 := &Piece{}
	stateSpace2 := make([]string, 0, 0)
	stateSpace2 = append(stateSpace, "King")
	stateSpace2 = append(stateSpace, "Pawn")
	kingInitialState2 := make(map[string][2]float64)
	kingInitialState2["King"] = [2]float64{math.Sqrt(2), 0}
	kingInitialState2["Pawn"] = [2]float64{math.Sqrt(2), 0}

	createPiece(t, qKing2, "Measurement", 1,
		kingInitialState2, stateSpace2, kingInitialState2, false)
	res2 := map[int]bool{
		1: true,
		8: true,
		9: true,
	}
	testAoF(t, qKing, 0, board2, res2)
}
