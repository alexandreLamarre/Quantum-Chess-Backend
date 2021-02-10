package quantumchess

type Board struct{
  Positions []int // 0 is the equivalent of null for javascript board
}

type Entanglements struct{
  List map[int]*Entanglement
}

type Entanglement struct{
  Elements []int
  State []float64
}

type Pieces struct{
  List map[int]*Piece
}

type Piece struct{
  Action string
  Color int
  InitialState map[string]float64
  StateSpace []string
  State map[string]float64
  Moved bool
}

func (board *Board) getID(id int) int{
  return board.Positions[id]
}

func (piece *Piece) getAction() string{
  return piece.Action
}



func (piece *Piece) getAreaOfInfluence(board *Board, newPos int) (map[int]bool, error){
  aof := make(map[int] bool)
  states, err := piece._getActivatedStates()
  for _, state := range states{
    legalTiles := _getPiecesInAoF(state, newPos, piece.Color)
    for _, v := range legalTiles{
      if !aof[v] {
        aof[v] = true
      }
    }
  }
  if err != nil{
    return aof, err
  }
  return aof, nil
}



func (piece *Piece) _getActivatedStates() ([]string, error) {
  var activatedStates []string
  for state, v := range piece.State{
    if v != 0{
      activatedStates = append(activatedStates, state)
    }
  }
  if len(activatedStates) == 0 {
    return activatedStates, InvalidMissingState(piece.StateSpace)
  }
  return activatedStates, nil
}

func _getPiecesInAoF(state string, pos int, color int) []int{
  var validSquares []int

  if state == "Pawn"{ //when checking this piece is already moved so we dont need to check two squares forward
    validSquares = checkForward(pos, color)
  } else if state == "Knight" {
    checkKnight(pos)
  } else if state == "Bishop"{
    checkBishop(pos)
  } else if state == "Rook"{
    checkRook(pos)
  } else if state == "Queen"{
    checkQueen(pos)
  } else if state == "King"{
    checkKing(pos)
  }

  return validSquares
}


func checkForward(pos int, color int) []int{
  var valid []int

  return valid
}
func checkBishop(pos int) []int{
  var valid []int

  return valid
}

func checkKnight(pos int) []int{
  var valid []int

  return valid
}

func checkRook(pos int) []int{
  var valid []int

  return valid
}

func checkQueen(pos int) []int{
  var valid []int

  return valid
}

func checkKing(pos int) []int{
  var valid []int

  return valid
}

