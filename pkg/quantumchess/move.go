package quantumchess

import (
	"fmt"
)

// DEBUGAPPLYMOVE toggles debug messages for the apply move function.
var DEBUGAPPLYMOVE bool = true

//ACTIONS is a string array representing the valid string quantum gates accepted by apply move
var ACTIONS [5]string = [5]string{"None", "Hadamard", "PauliX", "PauliZ", "Measurement"}

//ApplyMove applies a move to a board state : (board, entanglements, pieces)
// and updates its components in place.
// Returns nil if successful and an appropriate error if the assumptions are not met.
func ApplyMove(board *Board, entanglements *Entanglements, pieces *Pieces,
	startSquare int, endSquare int) error {
	if DEBUGAPPLYMOVE {
		fmt.Println("Applying move from ", startSquare, " to ", endSquare)
	}

	// CHECK CAPTURE
	movedPiece := board.Positions[startSquare]
	potentialPiece := board.Positions[endSquare]
	capture, err := checkCapture(pieces, movedPiece, potentialPiece)
	if err != nil {
		return err
	}

	if capture {
		piece1, piece2 := board.getID(startSquare), board.getID(endSquare)
		if piece1 == 0 {
			return InvalidPiece(startSquare)
		}
		if piece2 == 0 {
			return InvalidPiece(endSquare)
		}

		measure2(pieces, entanglements, piece1, piece2)
		err := processCapture(board, entanglements, pieces, endSquare)
		if err != nil {
			return err
		}
		move(board, startSquare, endSquare)

	} else { // CHECK OTHER BOARD STATES IF CAPTURE DOES NOT OCCUR
		piece1 := board.getID(startSquare)
		if piece1 == 0 {
			return InvalidPiece(startSquare)
		}
		piece := pieces.List[piece1]
		if piece == nil {
			return InvalidPieceAccess(piece1)
		}
		action := piece.getAction()
		if !validAction(action) {
			return InvalidAction(action)
		}
		if action == "None" {
			move(board, startSquare, endSquare) //move a piece as normal
		} else if action == "Measurement" && piece.inMixedState() {
			AoF, err := piece.getAreaOfInfluence(board, endSquare)
			if err != nil {
				return err
			}
			measureOnAoF(board, entanglements, pieces, AoF)
			move(board, startSquare, endSquare)
		} else if piece.inMixedState(){
			AoF, err := piece.getAreaOfInfluence(board, endSquare)
			if err != nil {
				return err
			}
			updateEntanglements(board, entanglements, pieces, piece1, action, AoF)
			move(board, startSquare, endSquare)
		} else{ //piece is in a determined state and cant exert its quantum action
			move(board, startSquare, endSquare)
		}
	}

	return nil
}

//checkCapture checks whether or not moving a piece from start square to end square will
//perform a capture on the board
func checkCapture(pieces *Pieces, startSquare int, endSquare int) (capture bool, err error) {
	if pieces.List[startSquare] == nil {
		fmt.Println("FATAL ERROR: PIECE TO MOVE DOESN'T EXIST")
		return false, InvalidMove(endSquare)
	}
	if pieces.List[endSquare] == nil {
		return false, nil
	} else if pieces.List[startSquare].Color == pieces.List[endSquare].Color {
		fmt.Println("FATAL ERROR: PIECE MOVES TO AN INVALID SQUARE: OCCUPIED BY SAME COLOR")
		return false, InvalidMove(endSquare)
	}

	return true, nil
}

func measureOnAoF(board *Board, entanglements *Entanglements, pieces *Pieces, aof map[int]bool) {
	for square := range aof {
		id := board.getID(square)
		if board.getID(square) != 0 {
			measure(pieces, entanglements, id)
		}
	}
}

//measure2 measures the states of all pieces entangled to piece1 and piece2.
// piece1 and piece2 are ids of the pieces being checked.
func measure2(pieces *Pieces, entanglements *Entanglements, piece1 int, piece2 int) {
	// if the pieces share entanglements perform measure on the only entangled systems
	if entanglements.List[piece2] != nil && find(entanglements.List[piece2].Elements, piece1) {
		measure(pieces, entanglements, piece2)
	} else { // perform measure on both separate entangled systems
		measure(pieces, entanglements, piece1)
		measure(pieces, entanglements, piece2)
	}

}

//measure1 measures the states of all pieces entangled to piece
func measure(pieces *Pieces, entanglements *Entanglements, piece int) {

}

//processCapture removes the piece being captured from board then moves the piece who captured
// it to its location. Checks that piece captured is no longer entangled or else raises an error.
// Deletes captured piece from entanglements and pieces struct.
func processCapture(board *Board, entanglements *Entanglements,
	pieces *Pieces, endSquare int) error {
	pieceToDeleteID := board.getID(endSquare)
	validDelete := checkNotEntangled(entanglements, pieceToDeleteID)
	if !validDelete {
		return InvalidEntanglementDelete(pieceToDeleteID)
	}
	delete(entanglements.List, pieceToDeleteID)
	delete(pieces.List, pieceToDeleteID)
	board.Positions[endSquare] = 0
	return nil
}

// checkNotEntangled checks all currently entangled elements,
// if the piece we are checking is still in entangled elements return true,
// else false
func checkNotEntangled(entanglements *Entanglements, id int) bool {
	for _, els := range entanglements.List {
		for _, pid := range els.Elements {
			if pid == id {
				return false
			}
		}
	}

	return true
}

func updateEntanglements(board *Board, entanglements *Entanglements, pieces *Pieces,
	pieceId int, action string, aof map[int]bool) error {

	kroneckerProductStack := make([][]float64, 0, 0)
	idStack := make([]int, 0, 0)
	fmt.Println(kroneckerProductStack, idStack)
	// append to Entangled elements recursively while checking not to add duplicates
	for id := range aof {
		els := entanglements.List[pieceId].Elements
		if !checkEntangledWith(entanglements, pieceId, id) {
			//if the piece is entangled but not with our piece add all its 'dependencies'
			if entanglements.List[id] != nil {
				err := addEntanglements(entanglements, pieceId, id, kroneckerProductStack, idStack)
				if err != nil {
					return err
				}
				// otherwise it is a standalone non-entangled state
			} else {
				entanglements.List[pieceId].Elements = append(els, id)
			}
		}
	}

	var finalState [][2]float64
	if len(entanglements.List[pieceId].Elements) >= 8 { //unstable quantum system collapses on itself (returns early)
		//measure all
		// and move.
		return nil
	}
	// pop state vector from kronecker stack and kronecker product it with
	// entanglements.List[pieceId].States

	// finally apply quantum action to entangled
	finalState = ApplyCircuit(action, len(entanglements.List[pieceId].Elements), entanglements.List[pieceId].State)
	fmt.Println(finalState)
	// update all entanglements [ids added] to the same value as entanglements[pieceID]

	return nil
}

func checkEntangledWith(entanglements *Entanglements, pieceId int, id int) bool {
	return true
}

func addEntanglements(entanglements *Entanglements, pieceId int, id int,
	kroneckerProductStack [][]float64, idStack []int) error {
	return nil
}

// find function return true if el is in arr.
func find(arr []int, el int) bool {
	for _, item := range arr {
		if item == el {
			return true
		}
	}
	return false
}

func validAction(a string) bool {
	for _, val := range ACTIONS {
		if val == a {
			return true
		}
	}
	return false
}

// move moves a piece from startSquare to endSquare, after the rest of processing
//for the turn is done
func move(board *Board, startSquare int, endSquare int) {
	tempPieceId := startSquare
	board.Positions[startSquare] = 0
	board.Positions[endSquare] = tempPieceId
}
