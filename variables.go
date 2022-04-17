package main

import "fmt"

func main() {
    // Stupid comment
    /* Most
      stupid
         comment */
    
    var one int = 1
    var two = 2
    three := 3.1 //works only inside func and need a value
    fmt.Println(one,two,three)
    
    var a string //""
    var b int  //0
    var c float32 //0
    var d bool //False
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    
    var e, f, g int = 10,20,30
    fmt.Println(e,f,g)

    x, y, z := 1,2,"three"
    fmt.Println(x,y,z)
}
