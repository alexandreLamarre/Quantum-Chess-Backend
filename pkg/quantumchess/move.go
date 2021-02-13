package quantumchess

import (
	"fmt"
	"math"
	"math/rand"
)

// DEBUGAPPLYMOVE toggles debug messages for the apply move function.
var DEBUGAPPLYMOVE bool = true

//ACTIONS is a string array representing the valid string quantum gates accepted by apply move
var ACTIONS [7]string = [7]string{"None", "Hadamard", "PauliX", "PauliZ", "Measurement", "PauliY", "SqrtNOT"}

//ApplyMove applies a move to a board state : (board, entanglements, pieces)
// and updates its components in place.
// Returns nil if successful and an appropriate error if the assumptions are not met.
func ApplyMove(board *Board, entanglements *Entanglements, pieces *Pieces,
	startSquare int, endSquare int) (err error) {
	if DEBUGAPPLYMOVE {
		fmt.Println("Applying move from ", startSquare, " to ", endSquare)
	}

	// CHECK CAPTURE
	movedPiece := board.Positions[startSquare]
	potentialPiece := board.Positions[endSquare]
	if DEBUGAPPLYMOVE {fmt.Println("Checking capture")}
	capture, err := checkCapture(pieces, movedPiece, potentialPiece)
	if err != nil {
		return err
	}

	if capture {
		if DEBUGAPPLYMOVE{ fmt.Println("A capture has occurred")}
		piece1, piece2 := board.getID(startSquare), board.getID(endSquare)
		if piece1 == 0 {
			return InvalidPiece(startSquare)
		}
		if piece2 == 0 {
			return InvalidPiece(endSquare)
		}

		if DEBUGAPPLYMOVE{fmt.Println("Measuring pieces involved in capture")}
		measure2(pieces, entanglements, piece1, piece2)
		if DEBUGAPPLYMOVE {fmt.Println("Processing captures...")}
		err := processCapture(board, entanglements, pieces, endSquare)
		if err != nil {
			return err
		}
		if DEBUGAPPLYMOVE {fmt.Println("Moving")}
		move(board, pieces, startSquare, endSquare)

	} else { // CHECK OTHER BOARD STATES IF CAPTURE DOES NOT OCCUR
		if DEBUGAPPLYMOVE{ fmt.Println("A capture has not occured")}
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
		if DEBUGAPPLYMOVE{fmt.Println("parsing actions....", action)}
		if action == "None" {
			if DEBUGAPPLYMOVE{fmt.Println("Moving")}
			move(board, pieces, startSquare, endSquare) //move a piece as normal
		} else if action == "Measurement" && piece.inMixedState() {
			if DEBUGAPPLYMOVE {fmt.Println("Measuring AoF")}
			AoF, err := piece.getAreaOfInfluence(board, endSquare, pieces)
			if err != nil {
				return err
			}
			measureOnAoF(board, entanglements, pieces, AoF)
			move(board, pieces, startSquare, endSquare)
		} else if piece.inMixedState() {
			if DEBUGAPPLYMOVE{fmt.Println("Update entanglements based on circuits")}
			AoF, err := piece.getAreaOfInfluence(board, endSquare, pieces)
			if err != nil {
				return err
			}
			eErr:= updateEntanglements(board, entanglements, pieces, piece1, action, AoF)
			if eErr != nil{
				return err
			}
			if DEBUGAPPLYMOVE{fmt.Println("Moving")}
			move(board, pieces, startSquare, endSquare)
		} else { //piece is in a determined state and cant exert its quantum action
			if DEBUGAPPLYMOVE {fmt.Println("Determined state piece moving...")}
			move(board, pieces, startSquare, endSquare)
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
		measure(pieces, entanglements, id)

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

	for _, v := range entanglements.List[piece].Elements {
		if len(pieces.List[piece].StateSpace) == 1 {
			continue
		}
		randInteger := rand.Float64()
		cur := 0.0
		maxProb := ""
		selected := ""
		for state, c := range pieces.List[v].State {
			pr := modulus(c)
			if pr > cur {
				maxProb = state
			}
			cur += pr
			if randInteger < cur {
				selected = state
				break
			}
		}
		//if a rounding error occurs, it selects the max probability result
		if selected == "" {
			selected = maxProb
		}

		for state := range pieces.List[v].State {
			if state == selected {
				pieces.List[v].State[state] = [2]float64{1.0, 0.0}
			} else {
				pieces.List[v].State[state] = [2]float64{0.0, 0.0}
			}
		}
		entanglements.List[piece] = nil //reset their entanglements
	}
}

func modulus(cmplx [2]float64) float64 {
	return math.Sqrt(math.Pow(cmplx[0], 2) + math.Pow(cmplx[1], 2))
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

	kroneckerProductStack := make([][][2]float64, 0, 0)
	idStack := make([]int, 0, 0)
	entanglements.List[pieceId] = &Entanglement{}

	// append to Entangled elements recursively while checking not to add duplicates
	if DEBUGAPPLYMOVE{fmt.Println("Checking AoF on...")}
	for id := range aof {
		pid := board.getID(id)
		if DEBUGAPPLYMOVE{fmt.Println(pid)}
		if !pieces.List[pid].inMixedState() {
			if DEBUGAPPLYMOVE{fmt.Println("Piece found to be in determined state")}
			//DO NOT ADD PIECE TO ENTANGLEMENTS, Apply the gate as normal
			//TODO: Apply gate to size 1 state vector of piece with id pid
			if len(pieces.List[pid].StateSpace) != 1 {
				if DEBUGAPPLYMOVE{fmt.Println("Apply action ", action)}
				newState := ApplyCircuit(action, 1, pieces.List[pid].getStateVector())
				i := 0
				if DEBUGAPPLYMOVE{fmt.Println("Copying result of quantum circuit to state")}
				for _, state := range pieces.List[pid].StateSpace {
					pieces.List[pid].State[state] = newState[i]
					i++
				}
				if DEBUGAPPLYMOVE{fmt.Println("State resulting from quantum action", pieces.List[pid].State)}
			}

		} else {
			if DEBUGAPPLYMOVE {fmt.Println("Piece found to be in mixed state")}

			if !checkEntangledWith(entanglements, pieceId, pid) {

				if DEBUGAPPLYMOVE {fmt.Println("Piece is not already entangled")}

				if entanglements.List[pid] != nil {
					if DEBUGAPPLYMOVE {fmt.Println("Piece is entangled to a different system")}
					for _, el := range entanglements.List[pid].Elements {
						idStack = append(idStack, el)
					}
					kroneckerProductStack = append(kroneckerProductStack,
						entanglements.List[pieceId].State)
				} else {
					if DEBUGAPPLYMOVE{fmt.Println("Piece is not entangled to a different system")}
					kroneckerProductStack = append(kroneckerProductStack, pieces.List[pid].getStateVector())
					idStack = append(idStack, pid)
				}
			}

		}
	}
	if DEBUGAPPLYMOVE{
		fmt.Println("=================")
		fmt.Println("id stack", idStack)
		fmt.Println("kronecker product", kroneckerProductStack)
	}

	for _, id:= range idStack{
		entanglements.List[pieceId].Elements = append(entanglements.List[pieceId].Elements, id)
	}

	//No entanglements were added
	if len(entanglements.List[pieceId].Elements) == 0 || len(entanglements.List[pieceId].Elements) == 1{
		entanglements.List[pieceId] = nil
		fmt.Println("No updated entanglements")
		return nil
	}

	//Too many entanglements were added
	if len(entanglements.List[pieceId].Elements) >= 8 { //unstable quantum system collapses on itself (returns early)
		for _, id := range entanglements.List[pieceId].Elements {
			measure(pieces, entanglements, id)
		}
		return nil
	}

	var finalState [][2]float64
	for _, vector := range kroneckerProductStack {
		entanglements.List[pieceId].State = kroneckerVectorProduct(entanglements.List[pieceId].State, vector)
	}

	//TODO: apply quantum circuit to the kronecker Product
	qbitSize := int(math.Log2(float64(len(entanglements.List[pieceId].Elements))))
	// finally apply quantum action to entangled
	finalState = ApplyCircuit(action, qbitSize, entanglements.List[pieceId].State)
	if DEBUGAPPLYMOVE {
		fmt.Println("final entangled state", finalState)
	}
	// update all entanglements [ids added] to the same value as entanglements[pieceID]

	states := unpackStatesFromEntangledState(finalState)
	if DEBUGAPPLYMOVE {
		fmt.Println(states)
	}

	//TODO: update each entanglement with the new state vector
	//TODO: set each pieces state based on its position in elements/Idstack to corresponding sum of entanglement state
	for i, state := range states {
		id := entanglements.List[pieceId].Elements[i]
		err := pieces.List[id].setState(state)
		if err != nil {
			return err
		}
		entanglements.List[id] = entanglements.List[pieceId]
	}

	return nil
}

func unpackStatesFromEntangledState(allState [][2]float64) [][][2]float64 {
	size := len(allState)
	res := make([][][2]float64, 0, 0)
	for size/2 >= 1 {
		value1 := [2]float64{0.0, 0.0}
		value2 := [2]float64{0.0, 0.0}
		i := 0
		flipped := false
		for _, v := range allState {
			if !flipped {
				value1[0] += v[0]
				value1[1] += v[1]
			} else {
				value2[0] += v[0]
				value2[1] += v[1]
			}
			i++
			if i == size/2 {
				flipped = !flipped
				i = 0
			}
		}
		res = append(res, [][2]float64{value1, value2})
		size = size / 2
	}
	return res
}

func kroneckerVectorProduct(v1 [][2]float64, v2 [][2]float64) [][2]float64 {
	res := make([][2]float64, 0, len(v1)*len(v2))
	fmt.Println()
	// TODO: do kronecker product of two vectors
	for _, a1 := range v1 {
		for _, a2 := range v2 {
			if DEBUGAPPLYMOVE {
				//fmt.Println("complex coefficients",a1, a2)
				fmt.Println("result:", cmplxMult(a1, a2))
			}
			res = append(res, cmplxMult(a1, a2))

		}
	}

	return res
}

func cmplxMult(v1 [2]float64, v2 [2]float64) [2]float64 {
	a := v1[0]
	b := v1[1]
	c := v2[0]
	d := v2[1]
	return [2]float64{a*c - b*d, b*c + a*d}
}

func checkEntangledWith(entanglements *Entanglements, pieceId int, id int) bool {
	if entanglements.List[pieceId] == nil{return false}
	for _, sid := range entanglements.List[pieceId].Elements {
		if id == sid {
			return true
		}
	}
	return false
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
func move(board *Board, pieces *Pieces, startSquare int, endSquare int) {
	tempPieceId := board.getID(startSquare)
	pieces.List[tempPieceId].Moved = true
	board.Positions[startSquare] = 0
	board.Positions[endSquare] = tempPieceId
}
