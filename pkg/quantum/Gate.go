package quantum

import (
	"fmt"
	"math"
)

//DEBUG_GATE toggles debug messages for gate functions
var DEBUG_GATE bool = false

//Gate represents a quantum gate. it factors out the constant in front of a matrix to attempt to reduce calculations.
type Gate struct {
	constant complex128
	matrix   []complex128 //matrices are stored in 1d arrays ordered by rows.
}

// Hadamard returns a Hadamard Gate of size qbit_size
func Hadamard(qbitSize int) (Gate, bool) {
	size := float64(qbitSize)
	var hadamard Gate
	if qbitSize <= 0 {
		//#TODO: ERROR HANDLING
		return hadamard, true
	}
	var constant = complex(1/math.Sqrt(2), 0)
	arraySize := int(math.Pow(2, 2*size))
	hadamard.matrix = make([]complex128, arraySize, arraySize)
	basicGate := [4]complex128{complex(1.0, 0.0), complex(1.0, 0.0),
		complex(1.0, 0.0), complex(-1.0, 0.0)}
	if DEBUG_GATE {
		fmt.Println(len(hadamard.matrix))
	}

	c, gate, err := createGate(qbitSize, constant, basicGate[:])
	if err {
		//#TODO ERROR HANDLING
	}
	hadamard.constant = c
	for i := 0; i < len(gate); i++ {
		hadamard.matrix[i] = gate[i]
	}

	return hadamard, false
}

//PauliX returns a PauliX Gate of size qbit_size
func PauliX(qbitSize int) (Gate, bool) {
	size := float64(qbitSize)
	var paulix Gate
	if qbitSize <= 0 {
		//#TODO: ERROR HANDLING
		return paulix, true
	}
	var constant = complex(1, 0)
	arraySize := int(math.Pow(2, 2*size))
	paulix.matrix = make([]complex128, arraySize, arraySize)
	basicGate := [4]complex128{complex(0.0, 0.0), complex(1.0, 0.0),
		complex(1.0, 0.0), complex(0.0, 0.0)}
	if DEBUG_GATE {
		fmt.Println(len(paulix.matrix))
	}

	c, gate, err := createGate(qbitSize, constant, basicGate[:])
	if err {
		//#TODO ERROR HANDLING
	}
	paulix.constant = c
	for i := 0; i < len(gate); i++ {
		paulix.matrix[i] = gate[i]
	}

	return paulix, false
}

//PauliY returns a PauliY Gate of size gbit_size
func PauliY(qbitSize int) (Gate, bool) {
	size := float64(qbitSize)
	var pauliy Gate
	if qbitSize <= 0 {
		//#TODO: ERROR HANDLING
		return pauliy, true
	}
	var constant = complex(1.0, 0.0)
	arraySize := int(math.Pow(2, 2*size))
	pauliy.matrix = make([]complex128, arraySize, arraySize)
	basicGate := [4]complex128{complex(0.0, 0.0), complex(0.0, -1.0),
		complex(0.0, 1.0), complex(0.0, 0.0)}
	if DEBUG_GATE {
		fmt.Println(len(pauliy.matrix))
	}

	c, gate, err := createGate(qbitSize, constant, basicGate[:])
	if err {
		//#TODO ERROR HANDLING
	}
	pauliy.constant = c
	for i := 0; i < len(gate); i++ {
		pauliy.matrix[i] = gate[i]
	}

	return pauliy, false
}

// PauliZ returns a PauliZ Gate of size qbit_size
func PauliZ(qbitSize int) (Gate, bool) {
	size := float64(qbitSize)
	var pauliz Gate
	if qbitSize <= 0 {
		//#TODO: ERROR HANDLING
		return pauliz, true
	}
	var constant = complex(1.0, 0)
	arraySize := int(math.Pow(2, 2*size))
	pauliz.matrix = make([]complex128, arraySize, arraySize)
	basicGate := [4]complex128{complex(1.0, 0.0), complex(0.0, 0.0),
		complex(0.0, 0.0), complex(-1.0, 0.0)}
	if DEBUG_GATE {
		fmt.Println(len(pauliz.matrix))
	}

	c, gate, err := createGate(qbitSize, constant, basicGate[:])
	if err {
		//#TODO ERROR HANDLING
	}
	pauliz.constant = c
	for i := 0; i < len(gate); i++ {
		pauliz.matrix[i] = gate[i]
	}

	return pauliz, false
}

func createGate(qbitSize int, c complex128, basicGate []complex128) (complex128, []complex128, bool) {
	var newGate []complex128 = basicGate[:]
	var newC complex128 = c
	var err bool
	for i := 1; i < qbitSize; i++ {
		newC, newGate, err = tensorProduct(newGate, newC, basicGate[:], c)
	}
	return newC, newGate, err
}

func tensorProduct(A []complex128, c1 complex128,
	B []complex128, c2 complex128) (complex128, []complex128, bool) {
	res := make([]complex128, len(A)*len(B), len(A)*len(B))
	var c complex128
	if c1 == c2 && c1 == complex(1/math.Sqrt(2), 0) {
		c = complex(0.5, 0.0) //most common constants will be sqrt(2), so here we avoid
		// at least some rounding errors
	} else {
		c = c1 * c2
	}

	rank1 := int(math.Sqrt(float64(len(A))))
	rank2 := int(math.Sqrt(float64(len(B))))
	i := 0
	for col := 0; col < rank1; col++ {
		for colB := 0; colB < rank2; colB++ {
			if DEBUG_GATE {
				fmt.Println("========")
			}
			for row := 0; row < rank1; row++ {
				aIJ := A[col*rank1+row]
				if DEBUG_GATE {
					fmt.Println("-")
				}
				for rowB := 0; rowB < rank2; rowB++ {
					bJK := B[colB*rank2+rowB]
					res[i] = aIJ * bJK
					if DEBUG_GATE {
						fmt.Println(aIJ, bJK)
					}
					i++
				}
			}
		}
	}

	return c, res, false
}
