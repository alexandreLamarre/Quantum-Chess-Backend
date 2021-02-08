package quantumchess

type Board struct{
  Positions []int
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
  color string
  initialState map[string]float64
  stateSpace []string
  state map[string]float64
}
