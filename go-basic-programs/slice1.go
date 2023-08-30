package main

import "fmt"

func main() {
    // Creating a slice using make
    a := make([]int, 5) // Creates a slice of length 5

    // Creating a slice from an array
    array := [5]int{1, 2, 3, 4, 5}
    b := array[1:4] // Creates a slice from index 1 (inclusive) to 4 (exclusive)

    // Creating a slice from an existing slice
    c := b[1:3] // Creates a slice from index 1 (inclusive) to 3 (exclusive)

    fmt.Println(a) // [0 0 0 0 0]
    fmt.Println(b) // [2 3 4]
    fmt.Println(c) // [3 4]
}

