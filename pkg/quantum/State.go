package quantum

import(
  "math"
)

type QuantumState struct{
  amplitudes []complex128
}

func MakeState(qbit_size int) QuantumState{
  var res QuantumState
  if qbit_size <= 0 {
    return res
  }
  size := int(math.Pow(2,float64(qbit_size)))
  res.amplitudes = make([]complex128, 0, size)
  return res
}

func SetState(q *QuantumState, vals []complex128){
  if(len(q.amplitudes) != len(vals)){return}

  for i:= 0; i < len(vals); i++{
    q.amplitudes[i] = vals[i]
  }
}

func (*QuantumState) ApplyGate(gate Gate){

  return
}
