package quantum

import(
  "fmt"
)

var DEBUG = false

func TestAllQuantum() bool {
  if(DEBUG){selectDebug()}

  fmt.Println("==== Gates ====\n")
  test:= testGates()
  fmt.Println()
  fmt.Println("Gates test sucessful?", test)
  fmt.Println()

  fmt.Println("==== States ==== \n")
  test2 := testStates()
  fmt.Println()
  fmt.Println("States test successful?", test2)

  fmt.Println("==== States n Gates ==== \n")
  test3 := testGatesAndStates()
  fmt.Println()
  fmt.Println("States n Gates test successful?", test3)

  fmt.Println("Passed All quantum tests?")
  return test && test2
}

/**
Turns on function println messages for debugging
depending on specific file we want to debug

DEBUG_GATE: Gate.go
**/
func selectDebug(){
  DEBUG_GATE = false
  DEBUG_STATE = true
}

func testGates() bool{
  h,err := Hadamard(1)
  if(err){ return false}
  h2, err := Hadamard(2)
  if(err){ return false}
  h3, err := Hadamard(3)
  if(err){ return false}
  fmt.Println(h)
  fmt.Println(h2)
  fmt.Println(h3)
  px,err := PauliX(2)
  if(err){return false}
  fmt.Println()
  fmt.Println(px)
  return true
}

func testStates() bool{
  s := MakeState(1)
  fmt.Println(s)
  var direk1 [2]complex128
  direk1[0] = complex(1.0, 0.0)
  direk1[1] = complex(0.0, 0.0)
  s.SetState(direk1[:])
  fmt.Println(s)
  return true
}

func testGatesAndStates() bool {
  s := MakeState(1)
  var direk1 [2]complex128
  direk1[0] = complex(1.0, 0.0)
  direk1[1] = complex(0.0, 0.0)
  s.SetState(direk1[:])
  fmt.Println("state", s)
  h, err:= Hadamard(1)
  if(err){return false}
  fmt.Println("gate", h)
  s.ApplyGate(h)
  fmt.Println("mixed state", s)
  return true
}
