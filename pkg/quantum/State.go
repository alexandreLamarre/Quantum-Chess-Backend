package quantum

import (
	"fmt"
	"math"
)

//DEBUG_STATE toggles debug messages on QuantumState and its methods
var DEBUG_STATE = false

//QuantumState represents the complex amplitudes of a quantum state.
type QuantumState struct {
	Amplitudes []complex128
}

//MakeState makes a QuantumState of size 2^n where n is qbitSize
func MakeState(qbitSize int) QuantumState {
	var res QuantumState
	if qbitSize <= 0 {
		return res
	}
	size := int(math.Pow(2, float64(qbitSize)))
	res.Amplitudes = make([]complex128, size, size)
	return res
}

//SetState sets the quantum state to the values specified
func (q *QuantumState) SetState(vals []complex128) {
	if DEBUG_STATE {
		fmt.Println(len(q.Amplitudes), len(vals))
	}
	if len(q.Amplitudes) != len(vals) {
		return
	}
	if DEBUG_STATE {
		fmt.Println("Setting state")
	}
	for i := 0; i < len(vals); i++ {
		q.Amplitudes[i] = vals[i]
	}
}

//ApplyGate applies a quantum gate to the quantum state.
func (q *QuantumState) ApplyGate(gate Gate) {
	size := int(math.Sqrt(float64(len(gate.matrix))))

	if DEBUG_STATE {
		fmt.Println(size, len(q.Amplitudes))
	}
	if size != len(q.Amplitudes) {
		return
	}

	if DEBUG_STATE {
		fmt.Println("Multiplying state and gate...")
	}

	tempArr := make([]complex128, len(q.Amplitudes), len(q.Amplitudes))
	//left matrix multiplcation: gate x state
	for col := 0; col < size; col++ {
		var temp complex128 = complex(0.0, 0.0)
		for row := 0; row < size; row++ {
			temp += gate.matrix[col*size+row] * q.Amplitudes[row] * gate.constant
		}
		tempArr[col] = temp
	}
	if DEBUG_STATE {
		fmt.Println("Obtained", tempArr)
	}

	for i := 0; i < size; i++ {
		q.Amplitudes[i] = tempArr[i]
	}

}
