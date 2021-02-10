package quantumchess

import "fmt"

type InvalidMove int
type InvalidPiece int
type InvalidPieceAccess int
type InvalidAction string
type InvalidEntanglementDelete int
type InvalidMissingState []string

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

func (e InvalidEntanglementDelete) Error() string {
	return fmt.Sprintf("Cannot delete id %d from entanglements: still entangled", e)
}

func (e InvalidMissingState) Error() string {
	return fmt.Sprintf("States `%s` are both 0 ", InvalidMissingState{})
}