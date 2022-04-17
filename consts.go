package main
import ("fmt")

//read-only
const PI float32 = 3.14
const EULER float32 = 2.71

//ok
const (
  A int = 2
  B = 2.2
  C = "careca"
)

func main() {
    fmt.Println(PI)
    fmt.Println(EULER)
    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(C)
}
