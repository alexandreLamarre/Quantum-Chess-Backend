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

//InvalidDeterminedState is an error returned when a quantum piece is passed in as a mixed state
// but is actually in a determined state
type InvalidDeterminedState []string

//InvalidSetState is an error returned when setting a state has too many values
// returns the state space of the piece whose state we tried to set
type InvalidSetState []string

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
	s := e[:]
	return fmt.Sprintf("Unrecognized action: %v", s)
}

func (e InvalidEntanglementDelete) Error() string {
	return fmt.Sprintf("Cannot delete id %d from entanglements: still entangled", e)
}

func (e InvalidMissingState) Error() string {
	s := e[:]
	return fmt.Sprintf("States `%v` are both 0 ", s)
}

func (e InvalidDeterminedState) Error() string {
	s := e[:]
	return fmt.Sprintf("%v state was passed in as a mixed state", s)
}

func (e InvalidSetState) Error() string {
	s := e[:]
	return fmt.Sprintf("Tried to set state of %v but failed", s)
}
