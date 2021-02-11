package quantumchess

import "fmt"

//InvalidMove is an error returned when we try to perform an invalid move on the board.
// Returns the square we want to move to.
type InvalidMove int

//InvalidPiece is an error returned when we expect a piece to be on the board, but it isn't
// Returns the position where we expected a piece.
type InvalidPiece int

//InvalidPieceAccess is an error returned when we get a piece ID from the board but it does not exist in the struct Pieces.List.
// Returns the piece ID of the piece that does not exist.
type InvalidPieceAccess int

//InvalidAction is an error returned when we parse a string and it is not one of the defined quantum actions.
// Returns the unexpected string.
type InvalidAction string

//InvalidEntanglementDelete is an error returned when we try to delete a piece from entanglements that is still entangled.
// Returns the piece id of such a piece.
type InvalidEntanglementDelete int

//InvalidMissingState is an error returned when a quantum piece has all zero states.
// Returns the piece's state space.
type InvalidMissingState []string

func (e InvalidMove) Error() string {
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

func (e InvalidEntanglementDelete) Error() string {
	return fmt.Sprintf("Cannot delete id %d from entanglements: still entangled", e)
}

func (e InvalidMissingState) Error() string {
	return fmt.Sprintf("States `%s` are both 0 ", InvalidMissingState{})
}
