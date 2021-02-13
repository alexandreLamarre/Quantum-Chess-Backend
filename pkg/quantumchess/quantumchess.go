package quantumchess

import (
	"fmt"
	"math"
)

//DEBUG_QUANTUM_CHESS_STRUCTS toggles debug messages for the structs and their methods/helpers
var DEBUG_QUANTUM_CHESS_STRUCTS bool = false

//WHITE represents the color of white player as an int
var WHITE int = 0

//BLACK represents the color of black player as an int
var BLACK int = 1

//Board a struct representing the positions of quantum pieces on a board in 1d integer array.
// A value of 0 indicates an empty tile. A non-zero value stores the pieceID of the piece at that tile
type Board struct {
	Positions []int // 0 is the equivalent of null for javascript board
}

//Entanglements is a struct that maps piece ids to their Entanglement.
type Entanglements struct {
	List map[int]*Entanglement
}

//Entanglement stores the data needed to specify entanglements. A list of piece ID's concerned in the entanglement.
// The whole state of the entanglement.
type Entanglement struct {
	Elements []int
	State    [][2]float64
}

//Pieces is a struct that maps piece ids to their Piece datatype.
type Pieces struct {
	List map[int]*Piece
}

//Piece stores the relevant information of a quantum piece.
type Piece struct {
	Action       string
	Color        int
	InitialState map[string][2]float64
	StateSpace   []string
	State        map[string][2]float64
	Moved        bool
}

//SetupInitialQuantumChess sets up the initial quantum chess board.
func SetupInitialQuantumChess(board *Board, entanglements *Entanglements, pieces *Pieces) {
	positions := [64]int{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32,
	}
	board.Positions = positions[:]

	//Rooks
	pieces.List = make(map[int]*Piece)
	qRookWhite := __createMixedPiece("Rook", "Pawn",
		true, 0, "Measurement")
	pieces.List[25] = qRookWhite
	qRookWhite2 := __createMixedPiece("Rook", "Pawn",
		true, 0, "Measurement")
	pieces.List[32] = qRookWhite2

	qRookBlack := __createMixedPiece("Rook", "Pawn",
		true, 1, "Measurement")
	pieces.List[1] = qRookBlack
	qRookBlack2 := __createMixedPiece("Rook", "Pawn",
		true, 1, "Measurement")
	pieces.List[8] = qRookBlack2

	qKnightWhite := __createMixedPiece("Knight", "Pawn",
		true, 0, "PauliX")

	qKnightWhite2 := __createMixedPiece("Knight", "Pawn",
		true, 0, "PauliX")
	pieces.List[26] = qKnightWhite
	pieces.List[31] = qKnightWhite2

	qKnightBlack := __createMixedPiece("Knight", "Pawn",
		true, 1, "PauliX")

	qKnightBlack2 := __createMixedPiece("Knight", "Pawn",
		true, 1, "PauliX")
	pieces.List[2] = qKnightBlack
	pieces.List[7] = qKnightBlack2

	qBishopWhite := __createMixedPiece("Bishop", "Pawn",
		true, 0, "SqrtNOT")
	qBishopWhite2 := __createMixedPiece("Bishop", "Pawn",
		true, 0, "SqrtNOT")
	pieces.List[27] = qBishopWhite
	pieces.List[30] = qBishopWhite2

	qBishopBlack := __createMixedPiece("Bishop", "Pawn",
		true, 1, "SqrtNOT")
	qBishopBlack2 := __createMixedPiece("Bishop", "Pawn",
		true, 1, "SqrtNOT")
	pieces.List[3] = qBishopBlack
	pieces.List[6] = qBishopBlack2

	qQueenWhite := __createMixedPiece("Queen", "Pawn",
		true, 0, "Hadamard")
	pieces.List[28] = qQueenWhite
	qQueenBlack := __createMixedPiece("Queen", "Pawn",
		true, 1, "Hadamard")
	pieces.List[5] = qQueenBlack

	qKingWhite := __createKing(0)
	pieces.List[29] = qKingWhite
	qKingBlack := __createKing(1)
	pieces.List[4] = qKingBlack

	qPawnRookWhite := __createMixedPiece("Pawn", "Rook", false, 0,
		"PauliZ")
	qPawnRookWhite2 := __createMixedPiece("Pawn", "Rook", false, 0,
		"PauliZ")

	pieces.List[17] = qPawnRookWhite
	pieces.List[24] = qPawnRookWhite2

	qPawnRookBlack := __createMixedPiece("Pawn", "Rook", false, 1,
		"PauliZ")
	qPawnRookBlack2 := __createMixedPiece("Pawn", "Rook", false, 1,
		"PauliZ")

	pieces.List[9] = qPawnRookBlack
	pieces.List[16] = qPawnRookBlack2

	qPawnKnightWhite := __createMixedPiece("Pawn", "Knight", false, 0,
		"PauliZ")
	qPawnKnightWhite2 := __createMixedPiece("Pawn", "Knight", false, 0,
		"PauliZ")

	pieces.List[18] = qPawnKnightWhite
	pieces.List[23] = qPawnKnightWhite2

	qPawnKnightBlack := __createMixedPiece("Pawn", "Knight", false, 1,
		"PauliZ")
	qPawnKnightBlack2 := __createMixedPiece("Pawn", "Knight", false, 1,
		"PauliZ")

	pieces.List[10] = qPawnKnightBlack
	pieces.List[15] = qPawnKnightBlack2

	qPawnBishopWhite := __createMixedPiece("Pawn", "Bishop", false, 0,
		"PauliZ")
	qPawnBishopWhite2 := __createMixedPiece("Pawn", "Bishop", false, 0,
		"PauliZ")

	pieces.List[19] = qPawnBishopWhite
	pieces.List[22] = qPawnBishopWhite2

	qPawnBishopBlack := __createMixedPiece("Pawn", "Bishop", false, 1,
		"PauliZ")
	qPawnBishopBlack2 := __createMixedPiece("Pawn", "Bishop", false, 1,
		"PauliZ")

	pieces.List[11] = qPawnBishopBlack
	pieces.List[14] = qPawnBishopBlack2

	qPawnKingWhite := __createMixedPiece("Pawn", "King", false, 0,
		"PauliZ")
	qPawnQueenWhite := __createMixedPiece("Pawn", "Queen", false, 0,
		"PauliZ")

	pieces.List[21] = qPawnKingWhite
	pieces.List[20] = qPawnQueenWhite

	qPawnKingBlack := __createMixedPiece("Pawn", "King", false, 1,
		"PauliZ")
	qPawnQueenBlack := __createMixedPiece("Pawn", "Queen", false, 1,
		"PauliZ")

	pieces.List[12] = qPawnKingBlack
	pieces.List[13] = qPawnQueenBlack

	for i := 1; i < 33; i++ {
		entanglements.List[i] = nil
	}
}

func __copyMap(input map[string][2]float64) map[string][2]float64 {
	newMap := make(map[string][2]float64)

	for k, v := range input {
		newMap[k] = v
	}

	return newMap
}

func __createMixedPiece(state1 string, state2 string, normal bool, color int,
	action string) *Piece {
	var stateVec [][2]float64
	if normal {
		stateVec = [][2]float64{{1 / math.Sqrt(2), 0.0}, {1 / math.Sqrt(2), 0.0}}
	} else {
		stateVec = [][2]float64{{1.0, 0.0}, {0.0, 0.0}}
	}
	ss := []string{state1, state2}
	is := map[string][2]float64{state1: stateVec[0], state2: stateVec[1]}
	state := __copyMap(is)

	return &Piece{Action: action, Color: color, InitialState: is,
		StateSpace: ss, State: state, Moved: false}
}

func __createKing(color int) *Piece {
	return &Piece{Action: "None", Color: color, InitialState: map[string][2]float64{"King": {1.0, 0.0}},
		StateSpace: []string{"King"}, State: map[string][2]float64{"King": {1.0, 0.0}}, Moved: false}
}

func (board *Board) getID(id int) int {
	return board.Positions[id]
}

func (piece *Piece) getAction() string {
	return piece.Action
}

func (piece *Piece) getStateVector() [][2]float64 {
	stateVec := make([][2]float64, 0, 0)

	for _, cmplx := range piece.State {
		stateVec = append(stateVec, cmplx)
	}
	return stateVec
}

func (piece *Piece) setState(input [][2]float64) error {
	if len(piece.StateSpace) != len(input) {
		return InvalidSetState(piece.StateSpace)
	}
	i := 0

	for _, state := range piece.StateSpace {
		piece.State[state] = input[i]
		i++
	}

	return nil
}

func (piece *Piece) inMixedState() bool {
	if len(piece.StateSpace) == 1 {
		return false
	}
	for _, v := range piece.State {
		if !nonZero(v) {
			return false
		}
	}
	return true
}

func (piece *Piece) getAreaOfInfluence(board *Board, newPos int, pieces *Pieces) (map[int]bool, error) {
	aof := make(map[int]bool)
	if !piece.inMixedState() {
		return aof, InvalidDeterminedState(piece.StateSpace)
	}
	states, err := piece._getActivatedStates()
	for _, state := range states {
		legalTiles := _getPiecesInAoF(state, newPos, piece.Color, board, pieces)
		for _, v := range legalTiles {
			aof[v] = true
		}
	}
	if err != nil {
		return aof, err
	}
	return aof, nil
}

//nonZero checks if a complex number in the form [2]float64
// is zero or not
func nonZero(cmplx [2]float64) bool {
	return cmplx[0] != 0 || cmplx[1] != 1
}

func (piece *Piece) _getActivatedStates() ([]string, error) {
	var activatedStates []string
	for state, v := range piece.State {
		if nonZero(v) {
			activatedStates = append(activatedStates, state)
		}
	}
	if len(activatedStates) == 0 {
		return activatedStates, InvalidMissingState(piece.StateSpace)
	}
	return activatedStates, nil
}

func _getPiecesInAoF(state string, pos int, color int, board *Board, pieces *Pieces) []int {
	var validSquares []int

	if state == "Pawn" { //when checking this piece is already moved so we dont need to check two squares forward
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking Pawn moves")
		}
		validSquares = checkForward(pos, color, board, validSquares)
	} else if state == "Knight" {
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking Knight moves")
		}
		validSquares = checkKnight(pos, board, validSquares)
	} else if state == "Bishop" {
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking Bishop moves")
		}
		validSquares = checkBishop(pos, board, validSquares)
	} else if state == "Rook" {
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking Rook moves")
		}
		validSquares = checkRook(pos, board, validSquares)
	} else if state == "Queen" {
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking Queen moves")
		}
		validSquares = checkQueen(pos, board, validSquares, color, pieces)
	} else if state == "King" {
		if DEBUG_QUANTUM_CHESS_STRUCTS {
			fmt.Println("Checking King moves")
		}
		validSquares = checkKing(pos, board, validSquares)
	}

	return validSquares
}

func checkForward(pos int, color int, board *Board, valid []int) []int {

	var dx int
	if color == WHITE {
		dx = -1
	} else {
		dx = 1
	}
	nextPos := pos + dx*8
	if inBoard(nextPos) && board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
	}
	return valid

}

func checkBishop(pos int, board *Board, valid []int) []int {
	valid = checkDiagonal(pos, board, 1, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, -1, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, 1, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, -1, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	return valid
}

// dx represents left -1, right + 1, dy represents up -1, down +1
func checkDiagonal(pos int, board *Board, dx int, dy int, valid []int) []int {
	nextPos := pos + dy*8
	if !inBoard(nextPos) {
		return valid
	}
	nextRow := getRow(nextPos)
	nextPos = nextPos + dx
	if nextRow != getRow(nextPos) {
		return valid
	}
	if board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
		return valid
	}

	return checkDiagonal(nextPos, board, dx, dy, valid)

}

func checkKnight(pos int, board *Board, valid []int) []int {
	valid = checkKnightMove(pos, board, 1, 0, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkKnightMove(pos, board, -1, 0, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkKnightMove(pos, board, 0, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkKnightMove(pos, board, 0, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	return valid
}

func checkKnightMove(pos int, board *Board, dx int, dy int, valid []int) []int {
	if dx != 0 {
		curRow := getRow(pos)
		directionPos := pos + dx*2
		if getRow(directionPos) != curRow {
			return valid
		}
		directionPosUp := directionPos - 8
		directionPosDown := directionPos + 8
		if inBoard(directionPosUp) && board.getID(directionPosUp) != 0 {
			valid = append(valid, directionPosUp)
		}
		if inBoard(directionPosDown) && board.getID(directionPosDown) != 0 {
			valid = append(valid, directionPosDown)
		}
	} else if dy != 0 {
		directionPos := pos + 16*dy
		if !inBoard(directionPos) {
			return valid
		}
		curRow := getRow(directionPos)
		directionPosLeft := directionPos - 1
		directionPosRight := directionPos + 1
		if inBoard(directionPosLeft) && curRow == getRow(directionPosLeft) &&
			board.getID(directionPosLeft) != 0 {
			valid = append(valid, directionPosLeft)
		}
		if inBoard(directionPosRight) && curRow == getRow(directionPosRight) &&
			board.getID(directionPosRight) != 0 {
			valid = append(valid, directionPosRight)
		}
	}
	return valid
}

func checkRook(pos int, board *Board, valid []int) []int {
	curRow := getRow(pos)
	valid = checkHorizontal(pos, curRow, board, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkHorizontal(pos, curRow, board, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkVertical(pos, board, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkVertical(pos, board, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	return valid
}

//dx -1 : left, dx: +1 right
func checkHorizontal(pos int, curRow int, board *Board, dx int, valid []int) []int {

	nextPos := pos + dx
	if !inBoard(nextPos) || getRow(nextPos) != curRow {
		return valid
	} else if board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
		return valid
	}
	return checkHorizontal(nextPos, curRow, board, dx, valid)
}

//dy -1: up, dy: +1 down
func checkVertical(pos int, board *Board, dy int, valid []int) []int {
	nextPos := pos + dy*8
	if !inBoard(nextPos) {
		return valid
	}
	if board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
		return valid
	}
	return checkVertical(nextPos, board, dy, valid)
}

func checkQueen(pos int, board *Board, valid []int, color int, pieces *Pieces) []int {
	curRow := getRow(pos)
	valid = checkDiagonal(pos, board, 1, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, -1, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, 1, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkDiagonal(pos, board, -1, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}

	valid = checkHorizontal(pos, curRow, board, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkHorizontal(pos, curRow, board, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkVertical(pos, board, -1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	valid = checkVertical(pos, board, 1, valid)
	if DEBUG_QUANTUM_CHESS_STRUCTS {
		fmt.Println(valid)
	}
	//check to make sure its only friendly pieces
	actualValid := make([]int, 0, 0)
	for _, id := range valid {
		if pieces.List[id].Color == color {
			actualValid = append(actualValid, id)
		}
	}
	return valid

}

func checkKing(pos int, board *Board, valid []int) []int {
	valid = checkKingVertical(pos, board, 1, valid)
	valid = checkKingVertical(pos, board, -1, valid)
	valid = checkKingHorizontal(pos, board, 1, valid)
	valid = checkKingHorizontal(pos, board, -1, valid)
	return valid
}

func checkKingVertical(pos int, board *Board, dy int, valid []int) []int {
	nextPos := pos + 8*dy
	if !inBoard(nextPos) {
		return valid
	}
	nextRow := getRow(nextPos)
	nextPosL := nextPos - 1
	nextPosR := nextPos + 1
	if board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
	}
	if inBoard(nextPosL) && getRow(nextPosL) == nextRow && board.getID(nextPosL) != 0 {
		valid = append(valid, nextPosL)
	}
	if inBoard(nextPosR) && getRow(nextPosR) == nextRow && board.getID(nextPosR) != 0 {
		valid = append(valid, nextPosR)
	}
	return valid
}

func checkKingHorizontal(pos int, board *Board, dx int, valid []int) []int {
	nextPos := pos + dx
	if !inBoard(nextPos) {
		return valid
	}
	curRow := getRow(pos)
	nextRow := getRow(nextPos)
	if curRow == nextRow && board.getID(nextPos) != 0 {
		valid = append(valid, nextPos)
	}
	return valid
}

// Helpers

func inBoard(pos int) bool {
	return 0 <= pos && pos < 64
}

func inBoardRow(row int) bool {
	return 0 <= row && row < 8
}

func getRow(pos int) int {
	return int(math.Floor(float64(pos / 8)))
}
