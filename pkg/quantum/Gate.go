package quantum

import(
  "math"
  "fmt"
)

var DEBUG_GATE bool = false

/**
Gate struct defines Quantum Gates
Matrices in quantum gates are square and are stored in 1D arrays ordered by rows
**/
type Gate struct{
  constant complex128
  matrix []complex128 //matrices are stored in 1d arrays ordered by rows.
}

func Hadamard(qbit_size int) (Gate, bool) {
  size := float64(qbit_size)
  var hadamard Gate;
  if(qbit_size <= 0){
    //#TODO: ERROR HANDLING
    return hadamard, true
  }
  var constant = complex(1/math.Sqrt(2), 0)
  array_size := int(math.Pow(2, 2*size))
  hadamard.matrix = make([]complex128, array_size, array_size);
  basic_gate := [4]complex128{complex(1.0, 0.0), complex(1.0, 0.0),
                                complex(1.0, 0.0), complex(-1.0, 0.0)}
  if(DEBUG_GATE){fmt.Println(len(hadamard.matrix))}

  c,gate,err := createGate(qbit_size, constant, basic_gate[:])
  if err{
    //#TODO ERROR HANDLING
  }
  hadamard.constant = c
  for i:= 0; i < len(gate); i++{
    hadamard.matrix[i] = gate[i]
  }

  return hadamard, false
}

func PauliX(qbit_size int) (Gate,bool){
  size := float64(qbit_size)
  var paulix Gate;
  if(qbit_size <= 0){
    //#TODO: ERROR HANDLING
    return paulix, true
  }
  var constant = complex(1, 0)
  array_size := int(math.Pow(2, 2*size))
  paulix.matrix = make([]complex128, array_size, array_size);
  basic_gate := [4]complex128{complex(0.0, 0.0), complex(1.0, 0.0),
                                complex(1.0, 0.0), complex(0.0, 0.0)}
  if(DEBUG_GATE){fmt.Println(len(paulix.matrix))}

  c,gate,err := createGate(qbit_size, constant, basic_gate[:])
  if err{
    //#TODO ERROR HANDLING
  }
  paulix.constant = c
  for i:= 0; i < len(gate); i++{
    paulix.matrix[i] = gate[i]
  }

  return paulix, false
}

func PauliY(qbit_size int) (Gate, bool){
  size := float64(qbit_size)
  var pauliy Gate;
  if(qbit_size <= 0){
    //#TODO: ERROR HANDLING
    return pauliy, true
  }
  var constant = complex(1.0, 0.0)
  array_size := int(math.Pow(2, 2*size))
  pauliy.matrix = make([]complex128, array_size, array_size);
  basic_gate := [4]complex128{complex(0.0, 0.0), complex(0.0, -1.0),
                                complex(0.0, 1.0), complex(0.0, 0.0)}
  if(DEBUG_GATE){fmt.Println(len(pauliy.matrix))}

  c,gate,err := createGate(qbit_size, constant, basic_gate[:])
  if err{
    //#TODO ERROR HANDLING
  }
  pauliy.constant = c
  for i:= 0; i < len(gate); i++{
    pauliy.matrix[i] = gate[i]
  }

  return pauliy, false
}

func PauliZ(qbit_size int) (Gate, bool){
  size := float64(qbit_size)
  var pauliz Gate;
  if(qbit_size <= 0){
    //#TODO: ERROR HANDLING
    return pauliz, true
  }
  var constant = complex(1.0, 0)
  array_size := int(math.Pow(2, 2*size))
  pauliz.matrix = make([]complex128, array_size, array_size);
  basic_gate := [4]complex128{complex(1.0, 0.0), complex(0.0, 0.0),
                                complex(0.0, 0.0), complex(-1.0, 0.0)}
  if(DEBUG_GATE){fmt.Println(len(pauliz.matrix))}

  c,gate,err := createGate(qbit_size, constant, basic_gate[:])
  if err{
    //#TODO ERROR HANDLING
  }
  pauliz.constant = c
  for i:= 0; i < len(gate); i++{
    pauliz.matrix[i] = gate[i]
  }

  return pauliz, false
}

func createGate(qbit_size int, c complex128, basic_gate []complex128) (complex128, []complex128, bool){
    var new_gate []complex128 = basic_gate[:]
    var new_c complex128 = c
    var err bool
    for i:= 1; i < qbit_size; i++{
        new_c, new_gate, err = tensorProduct(new_gate, new_c, basic_gate[:], c)
    }
    return new_c, new_gate, err
}



func tensorProduct(A []complex128, c1 complex128,
                    B []complex128, c2 complex128) (complex128, []complex128, bool){
  res:= make([]complex128, len(A)*len(B), len(A)*len(B))
  var c complex128
  if(c1 == c2 && c1 == complex(1/math.Sqrt(2), 0)){
    c = complex(0.5, 0.0) //most common constants will be sqrt(2), so here we avoid
                            // at least some rounding errors
  } else{
    c = c1*c2
  }

  rank1 := int(math.Sqrt(float64(len(A))))
  rank2 := int(math.Sqrt(float64(len(B))))
  i:= 0
  for col := 0; col < rank1; col++{
    for col_b := 0; col_b < rank2; col_b++{
      if(DEBUG_GATE){fmt.Println("========")}
      for row:= 0; row < rank1; row++{
        a_ij := A[col*rank1 + row]
        if(DEBUG_GATE){fmt.Println("-")}
        for row_b:= 0; row_b < rank2; row_b++{

          b_jk := B[col_b*rank2 + row_b]
          res[i] = a_ij * b_jk
          if(DEBUG_GATE){fmt.Println(a_ij, b_jk)}
          i++
        }
      }
    }
  }

  return c, res, false
}
