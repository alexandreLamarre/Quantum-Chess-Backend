package quantum

type QuantumState struct{
  Real float64
  Imaginary float64
}

func (*QuantumState) ApplyGate(gate Gate){
  return
}
