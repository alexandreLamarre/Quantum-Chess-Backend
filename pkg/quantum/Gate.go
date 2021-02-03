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

func Hadamard(qbit_size int) Gate {
  size := float64(qbit_size)
  var hadamard Gate;
  var sqrt_constant = complex(math.Sqrt(2), 0)
  hadamard.constant = sqrt_constant
  array_size := int(math.Pow(2, 2*size))
  hadamard.matrix = make([]complex128, array_size, array_size);
  basic_gate := [4]complex128{complex(1.0, 0.0), complex(1.0, 0.0),
                                complex(1.0, 0.0), complex(-1.0, 0.0)}
  if(DEBUG_GATE){fmt.Println(len(hadamard.matrix))}
  if qbit_size == 1 {
    for i := 0; i < len(basic_gate); i++{
      hadamard.matrix[i] = basic_gate[i]
    }
  } else{
    var new_gate []complex128
    var new_c complex128
    for i:= 1; i < qbit_size; i++{
      if(i == 1){
        new_gate, new_c = tensorProduct(basic_gate[:], sqrt_constant, basic_gate[:], sqrt_constant)
      } else{
        new_gate, new_c = tensorProduct(new_gate, new_c, basic_gate[:], sqrt_constant)
      }
      if(DEBUG_GATE){ fmt.Println("tensor product",new_gate, new_c)}
    }
    hadamard.constant = new_c

    for i:= 0; i < len(new_gate); i++{
      hadamard.matrix[i] = new_gate[i]
    }
  }
  return hadamard
}

func tensorProduct(A []complex128, c1 complex128,
                    B []complex128, c2 complex128) (new_gate []complex128, new_c complex128){
  res:= make([]complex128, len(A)*len(B), len(A)*len(B))
  var c complex128
  if(c1 == c2 && c1 == complex(math.Sqrt(2), 0)){
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

  return res, c
}
