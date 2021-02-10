package quantumchess

import (
	"fmt"
)

var DEBUGAPPLYMOVE bool = true
var ACTIONS [5]string = [5]string{"None", "Hadamard", "PauliX", "PauliZ", "Measurement"}

type InvalidMove int
type InvalidPiece int
type InvalidPieceAccess int
type InvalidAction string

func(e InvalidMove) Error() string {
	return fmt.Sprintf("Illegal move to position %d", e)
}

func (e InvalidPiece) Error() string {
	return fmt.Sprintf(" Illegal piece on Board at position %d", e)
}

func (e InvalidPieceAccess) Error() string {
	return fmt.Sprintf("Illegal piece accessed in pieces struct for id %d", e)
}

func (e InvalidAction) Error() string {
	return fmt.Sprintf("Unrecognized action %s", e)
}

func ApplyMove(board *Board, entanglements * Entanglements, pieces * Pieces,
	startSquare int, endSquare int)(*Board, *Entanglements, *Pieces, error){
	if DEBUGAPPLYMOVE {fmt.Println("Applying move from ", startSquare, " to ", endSquare)}

	// CHECK CAPTURE
	movedPiece := board.Positions[startSquare]
	potentialPiece := board.Positions[endSquare]
	capture, err := checkCapture(pieces, movedPiece, potentialPiece)
	if err != nil {
		return board, entanglements, pieces, err
	}

	if capture {
		piece1, piece2 := board.getID(startSquare), board.getID(endSquare)
		if piece1 == 0{
			return board, entanglements, pieces, InvalidPiece(startSquare)
		}
		if piece2 == 0{
			return board, entanglements, pieces, InvalidPiece(endSquare)
		}

		measure2(pieces, entanglements, piece1, piece2)
		processCapture(board, startSquare, endSquare)

	} else{ // CHECK OTHER BOARD STATES IF CAPTURE DOES NOT OCCUR
		piece1 := board.getID(startSquare)
		if piece1 == 0{
			return board, entanglements, pieces, InvalidPiece(startSquare)
		}
		piece := pieces.List[piece1]
		if piece == nil{
			return board, entanglements, pieces, InvalidPieceAccess(piece1)
		}
		action := piece.getAction()
		if !validAction(action){
			return board, entanglements, pieces, InvalidAction(action)
		}
		if action == "None"{
			move(board, startSquare, endSquare) //move a piece as normal
		} else if action == "Measurement"{
			AoF := piece.getAreaOfInfluence(board, endSquare)
			measureOnAoF(board, entanglements, pieces, AoF)
			move(board, startSquare, endSquare)
		} else{
			AoF := piece.getAreaOfInfluence(board, endSquare)
			updateEntanglements(board, entanglements, pieces, piece1, action, AoF)
			move(board, startSquare, endSquare)
		}
	}

	return board, entanglements, pieces, nil
}

//checkCapture checks whether or not moving a piece from start square to end square will
//perform a capture on the board
func checkCapture(pieces *Pieces, startSquare int, endSquare int) (capture bool, err error) {
	if pieces.List[startSquare] == nil{
		fmt.Println("FATAL ERROR: PIECE TO MOVE DOESN'T EXIST")
		return false, InvalidMove(endSquare)
	}
	if pieces.List[endSquare] == nil{
		return false, nil
	} else if pieces.List[startSquare].color == pieces.List[endSquare].color {
		fmt.Println("FATAL ERROR: PIECE MOVES TO AN INVALID SQUARE: OCCUPIED BY SAME COLOR")
		return false, InvalidMove(endSquare)
	}

	return true, nil
}

func measureOnAoF(board *Board, entanglements *Entanglements, pieces *Pieces, aof []int){
	for _, square := range aof{
		id := board.getID(square)
		if board.getID(square) != 0 {
			measure(pieces, entanglements, id)
		}
	}
}


//measure2 measures the states of all pieces entangled to piece1 and piece2.
// piece1 and piece2 are ids of the pieces being checked.
func measure2(pieces *Pieces, entanglements *Entanglements, piece1 int, piece2 int ){
	// if the pieces share entanglements perform measure on the only entangled systems
	if entanglements.List[piece2] != nil && find(entanglements.List[piece2].Elements, piece1 ){
		measure(pieces, entanglements, piece2)
	} else{ // perform measure on both separate entangled systems
		measure(pieces, entanglements, piece1)
		measure(pieces, entanglements, piece2)
	}

}

//measure1 measures the states of all pieces entangled to piece
func measure(pieces *Pieces, entanglements *Entanglements, piece int){

}

//processCapture removes the piece being captured from board then moves the piece who captured
// it to its location. Checks that piece captured is no longer entangled or else raises an error.
// Deletes captured piece from entanglements and pieces struct.
func processCapture(board *Board, startSquare int, endSquare int){

}


func updateEntanglements(board *Board, entanglements *Entanglements, pieces *Pieces,
	pieceId int, action string, aof []int) error{


	return nil
}


// find function return true if el is in arr.
func find(arr []int, el int) bool{
	for _, item := range arr{
		if item == el{
			return true
		}
	}
	return false
}

func validAction(a string) bool {
	for _, val := range ACTIONS{
		if val == a { return true}
	}
	return false
}


// move moves a piece from startSquare to endSquare, after the rest of processing
//for the turn is done
func move(board *Board, startSquare int, endSquare int){
	tempPieceId := startSquare
	board.Positions[startSquare] = 0
	board.Positions[endSquare] = tempPieceId
}



