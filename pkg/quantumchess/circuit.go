package quantumchess

import (
	"fmt"
	"github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/quantum"
)

//ApplyCircuit converts an action string to a QuantumGate of size qbitSize to be used on input state
// returns the state cast to the real part of the QuantumState obtained
func ApplyCircuit(action string, qbitSize int, state []float64 ) []float64 {
	complexState := parseFloat64ArrayToComplex128(state)
	gate, err := parseCircuit(action, qbitSize)
	if err {
		fmt.Println(err)
	}
	cs := &quantum.QuantumState{}
	cs.SetState(complexState)
	cs.ApplyGate(gate)
	return parseComplex128ArrayToFloat64(cs.Amplitudes)

}

func parseCircuit(action string, qbitSize int) (quantum.Gate, bool){
	if action == "Hadamard"{
		return quantum.Hadamard(qbitSize)

	}else if action == "PauliX"{
		return quantum.PauliX(qbitSize)
	} else {
		return quantum.PauliZ(qbitSize)
	}

}

func parseFloat64ArrayToComplex128(state []float64) []complex128{
	var complexState []complex128
	for _, s := range state{
		complexState = append(complexState, complex(s, 0.0))
	}
	return complexState
}

func parseComplex128ArrayToFloat64(cState []complex128) []float64{
	var state []float64
	for _, cs := range cState{
		state = append(state, real(cs))
	}
	return state
}