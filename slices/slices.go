package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    outer := make([][]uint8, dx)
    for i := range outer {
        inner := make([]uint8, dy)
        outer[i] = inner 
        for j := range outer[i] {
            inner[j] = uint8((i + j) / 2)
        }

    }
    return outer;        
}

func main() {
	pic.Show(Pic)
}
