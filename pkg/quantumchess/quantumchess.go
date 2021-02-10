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
  color int
  initialState map[string]float64
  stateSpace []string
  state map[string]float64
}

func (board *Board) getID(id int) int{
  return board.Positions[id]
}

func (piece *Piece) getAction() string{
  return piece.Action
}

func (piece *Piece) getAreaOfInfluence(board *Board, newPos int) []int{
  var aof []int
  return aof
}