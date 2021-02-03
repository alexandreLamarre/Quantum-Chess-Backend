package quantum

import(
  "fmt"
  "math"
  "time"
)

var DEBUG = false

func TestAllQuantum() bool {
  if(DEBUG){selectDebug()}

  h := Hadamard(1)
  h2 := Hadamard(2)
  start := time.Now()
  h3 := Hadamard(8)
  duration := time.Since(start)
  fmt.Println("time to execute Hadamard", duration)
  fmt.Println(h)
  fmt.Println(h2)
  size:= int(math.Sqrt(float64(len(h3.matrix))))
  fmt.Println("")
  fmt.Println("")
  for i:= 0; i < size; i++{
    // fmt.Println(h3.matrix[i*size:(i+1)*size])
  }
  fmt.Println("Passed All quantum tests?")
  return true
}

/**
Turns on function println messages for debugging
depending on specific file we want to debug

DEBUG_GATE: Gate.go
**/
func selectDebug(){
  DEBUG_GATE = true
}
