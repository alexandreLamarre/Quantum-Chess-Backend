package quantumchess

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

func (board *Board) getID(id int) int {
	return board.Positions[id]
}

func (piece *Piece) getAction() string {
	return piece.Action
}

func (piece *Piece) getAreaOfInfluence(board *Board, newPos int) (map[int]bool, error) {
	aof := make(map[int]bool)
	states, err := piece._getActivatedStates()
	for _, state := range states {
		legalTiles := _getPiecesInAoF(state, newPos, piece.Color, board)
		for _, v := range legalTiles {
			if !aof[v] {
				aof[v] = true
			}
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

func _getPiecesInAoF(state string, pos int, color int, board *Board) []int {
	var validSquares []int

	if state == "Pawn" { //when checking this piece is already moved so we dont need to check two squares forward
		validSquares = checkForward(pos, color, board)
	} else if state == "Knight" {
		checkKnight(pos, board)
	} else if state == "Bishop" {
		checkBishop(pos, board)
	} else if state == "Rook" {
		checkRook(pos, board)
	} else if state == "Queen" {
		checkQueen(pos, board)
	} else if state == "King" {
		checkKing(pos, board)
	}

	return validSquares
}

func checkForward(pos int, color int, board *Board) []int {
	var valid []int

	return valid
}
func checkBishop(pos int, board *Board) []int {
	var valid []int

	return valid
}

func checkKnight(pos int, board *Board) []int {
	var valid []int

	return valid
}

func checkRook(pos int, board *Board) []int {
	var valid []int

	return valid
}

func checkQueen(pos int, board *Board) []int {
	var valid []int

	return valid
}

func checkKing(pos int, board *Board) []int {
	var valid []int

	return valid
}
