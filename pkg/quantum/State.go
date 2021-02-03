package quantum

import(
  "fmt"
  "math"
)

var DEBUG_STATE = false

type QuantumState struct{
  amplitudes []complex128
}

func MakeState(qbit_size int) QuantumState{
  var res QuantumState
  if qbit_size <= 0 {
    return res
  }
  size := int(math.Pow(2,float64(qbit_size)))
  res.amplitudes = make([]complex128, size, size)
  return res
}

func (q *QuantumState) SetState(vals []complex128) {
  if(DEBUG_STATE){fmt.Println(len(q.amplitudes), len(vals))}
  if(len(q.amplitudes) != len(vals)){return }
  if(DEBUG_STATE) {fmt.Println("Setting state")}
  for i:= 0; i < len(vals); i++{
    q.amplitudes[i] = vals[i]
  }
}

func (q *QuantumState) ApplyGate(gate Gate){
  size := int(math.Sqrt(float64(len(gate.matrix))))

  if(DEBUG_STATE){fmt.Println(size, len(q.amplitudes))}
  if size != len(q.amplitudes){
    return
  }

  if(DEBUG_STATE){fmt.Println("Multiplying state and gate...")}

  temp_arr := make([]complex128, len(q.amplitudes), len(q.amplitudes))
  //left matrix multiplcation: gate x state
  for col := 0; col < size; col++{
    var temp complex128 = complex(0.0, 0.0)
    for row := 0; row < size; row++{
      temp += gate.matrix[col*size+ row] * q.amplitudes[row] * gate.constant
    }
    temp_arr[col] = temp
  }
  if(DEBUG_STATE){fmt.Println("Obtained", temp_arr)}

  for i := 0; i < size; i++ {
    q.amplitudes[i] = temp_arr[i]
  }

}
