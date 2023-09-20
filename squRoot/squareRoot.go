package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) float64 {
    fmt.Println("Finding sqrt for ", x)
    z := 1.0;
    for i := 0; i < 10; i++ {
        zsqu := z * z;
        fmt.Println("square ", zsqu)
        dif := zsqu - x
        fmt.Println("diff ", dif)
        change := dif / (2 * math.Abs(z))
        fmt.Println("change: ", change)
        if dif == 0  {
            fmt.Println("squrt is: ", z, " found in iterations: ", i)
            return z
        } else if dif > 0 {
            // z too big
            fmt.Println("too big")
            z = z - change;
        } else {
            // z too small
            fmt.Println("too small")
            z = z - change;
        }
        fmt.Println(z)
    }
    return z;
}

func main() {
   sqrt(4) 
   sqrt(9) 
   sqrt(16) 
   sqrt(25) 
   sqrt(10) 
}


