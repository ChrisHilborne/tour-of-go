package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))	
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
    fmt.Println("Finding sqrt for ", x)
    z := 1.0;
    for i := 0; i < 10; i++ {
        zsqu := z * z;
        dif := zsqu - x
        change := dif / (2 * math.Abs(z))
        if dif == 0  {
            return z, nil
        } else if dif > 0 {
            // z too big
            z = z - change;
        } else {
            // z too small
            z = z - change;
        }
        fmt.Println(z)
    }
    return z, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
